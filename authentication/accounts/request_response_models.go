package accounts

type requestBodyCreateAccount struct {
	Name     string
	Email    string
	Password string
}

type requestBodyLoginAccount struct {
	Name     string
	Password string
}
