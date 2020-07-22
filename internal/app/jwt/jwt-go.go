package jwt

type AccessToken struct {
	SecretKey byte
}

type RefreshToken struct {
}

func (t *RefreshToken) Refresh() {

}

func (t *RefreshToken) Generate() {

}

func (t *AccessToken) Decode() {

}

func (t *AccessToken) Encoded() {

}

func (t *AccessToken) Update() {

}
