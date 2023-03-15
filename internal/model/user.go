package model

type UserCreateInput struct {
	Passport string
	Password string
	Nickname string
}

type UserSignInInput struct {
	Passport string
	Password string
}

type UserLoginInput struct {
	Mobile string
	Code   string
}
