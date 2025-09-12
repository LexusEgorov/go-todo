package models

//TODO: add fields
//TODO: move to right places
type AuthDTO struct{}
type RegisterDTO struct{}
type TokensDTO struct{}

type Auth struct{}

func (a Auth) ToDTO() AuthDTO {
	return AuthDTO{}
}

type Register struct{}

func (r Register) ToDTO() RegisterDTO {
	return RegisterDTO{}
}

type Tokens struct{}

func (t Tokens) ToDTO() TokensDTO {
	return TokensDTO{}
}
