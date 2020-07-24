package apiserver

import (
	"errors"
	"fmt"
	"github.com/Gentostage/golang-auth/internal/app/jwt"
	"github.com/Gentostage/golang-auth/internal/app/model"
	"github.com/Gentostage/golang-auth/internal/app/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type server struct {
	router  *gin.Engine
	logger  *logrus.Logger
	store   store.Store
	access  *jwt.AccessToken
	refresh *jwt.RefreshToken
}

func newServer(config Config, store store.Store) *server {
	logger := logrus.New()
	loggerLevel, _ := logrus.ParseLevel(config.LogLevel)
	logger.SetLevel(loggerLevel)
	access := &jwt.AccessToken{
		SecretKey:  "vdgjfesbf tc,jug,jutkufr,jf,juf,f,f,uj f",
		TimeToLive: 2,
	}
	refresh := &jwt.RefreshToken{
		TimeToLive: 60,
	}
	s := &server{
		router:  gin.New(),
		logger:  logger,
		store:   store,
		access:  access,
		refresh: refresh,
	}
	s.router.Use(s.Middle())

	s.configureRoute()
	return s
}
func (s *server) Middle() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, _ := context.Cookie("Access_Token")
		tokenHeader := context.Request.Header.Get("X-Auth-Token")

		if tokenHeader != "" {
			token = tokenHeader
		}
		if token != "" {
			err := s.store.InvalidToken().Find(token)
			if err != nil {
				return
			}
			tokenStruct, err := s.access.Decode(token)
			if err != nil {
				context.Set("token_struct", tokenStruct)
				return
			}
			context.Set("user_id", tokenStruct.Payload.User)
			context.Set("token_struct", tokenStruct)
			fmt.Println(tokenStruct)
		}
	}

}

func (s *server) configureRoute() {
	s.router.GET("/user/", func(context *gin.Context) {
		u := &model.User{}
		userId, exist := context.Get("user_id")

		if !exist {
			context.String(http.StatusUnauthorized, "Ошибка доступа")
			return
		}

		u.ID = userId.(primitive.ObjectID)
		user, err := s.store.User().Get(u)
		if err != nil {
			s.logger.Info(err)
			context.String(http.StatusNotFound, err.Error())
		} else {
			context.JSON(http.StatusOK, user)
		}
	})

	s.router.POST("/user/create", func(context *gin.Context) {
		user := &model.User{}
		if err := context.BindJSON(&user); err != nil {
			context.String(http.StatusBadRequest, err.Error())
		}
		userTemp := &model.User{
			Email: user.Email,
		}
		u, _ := s.store.User().Get(userTemp)
		if u == nil {
			err := user.Validate()
			if err == nil {
				err = user.GeneratePassword()
				if err == nil {
					err = s.store.User().Create(user)
					if err == nil {
						context.JSON(http.StatusOK, user)
						return
					}
				}
			}
			s.logger.Error(err)
			context.String(http.StatusBadRequest, err.Error())

		} else {
			context.String(http.StatusBadRequest, "Пользователь с таким email уже существует")
		}

	})

	s.router.POST("/user/login", func(context *gin.Context) {
		user := &model.User{}
		if err := context.BindJSON(&user); err != nil {
			context.String(http.StatusBadRequest, err.Error())
		}
		tempUser := &model.User{
			Email: user.Email,
		}
		u, err := s.store.User().Get(tempUser)
		if u != nil {
			if u.ComparePassword(user.Password) {
				token, err := s.access.Encode(u)
				if err != nil {
					context.String(http.StatusInternalServerError, err.Error())
					return
				}

				refreshToken, createTime := s.refresh.Generate(u)
				modelToken := &model.Token{
					UserId:       u.ID,
					RefreshToken: refreshToken,
					RegisterTime: createTime,
					Alive:        true,
				}
				err = modelToken.GenerateHashToken(token)
				if err != nil {
					context.String(http.StatusInternalServerError, err.Error())
					return
				}
				err = s.store.Token().Create(modelToken)
				if err != nil {
					context.String(http.StatusInternalServerError, err.Error())
				}
				context.SetCookie("Access_Token", token, 3600, "/", "127.0.0.1", false, true)
				context.JSON(http.StatusOK, struct {
					AccessToken  string `json:"access_token"`
					RefreshToken string `json:"refresh_token"`
				}{
					token,
					refreshToken,
				})
				return
			}
		}
		s.logger.Error(err)
		context.String(http.StatusBadRequest, "Неверно имя пользователя или пароль")
	})

	s.router.POST("/user/logout", func(context *gin.Context) {
		token, _ := context.Cookie("Access_Token")
		tokenHeader := context.Request.Header.Get("X-Auth-Token")

		if tokenHeader != "" {
			token = tokenHeader
		}
		if token != "" {
			err := s.store.InvalidToken().Insert(token)
			if err != nil {
				context.String(http.StatusBadRequest, err.Error())
			}
		}
		context.Status(http.StatusOK)
	})

	s.router.POST("/token/refresh", func(context *gin.Context) {
		token := &struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}{}
		if err := context.BindJSON(&token); err != nil {
			context.String(http.StatusBadRequest, err.Error())
			return
		}
		tokenStruct, _ := context.Get("token_struct")
		tokenData := tokenStruct.(jwt.TokenStructData)
		tokenModel := model.Token{
			UserId: tokenData.Payload.User,
			Alive:  true,
		}

		tokenResult, err := s.store.Token().GetAllAliveTokensByUser(&tokenModel)
		if err != nil {
			context.String(http.StatusUnauthorized, err.Error())
			return
		}
		var validToken *model.Token
		for _, r := range tokenResult {
			err = r.CompareTokens(token.RefreshToken, token.AccessToken)
			if err == nil {
				validToken = r
				s.logger.Info(validToken)
				break
			}
		}
		if validToken == nil {
			context.String(http.StatusForbidden, errors.New("Доступ запрещен ").Error())
			return
		}
		err = s.store.Token().Close(validToken)
		user, err := s.store.User().Get(&model.User{ID: validToken.UserId})
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		token.AccessToken, err = s.access.Encode(user)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		var timeCreate time.Time
		token.RefreshToken, timeCreate = s.refresh.Generate(user)
		tokenModel = model.Token{
			ID:           primitive.ObjectID{},
			RefreshToken: token.RefreshToken,
			RegisterTime: timeCreate,
			Alive:        true,
			UserId:       user.ID,
		}
		err = tokenModel.GenerateHashToken(token.AccessToken)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		err = s.store.Token().Create(&tokenModel)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		context.JSON(http.StatusOK, token)
	})

}
