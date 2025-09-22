package models

//TODO: add fields
//TODO: move to right places
type AuthDTO struct{}
type RegisterDTO struct{}
type TokensDTO struct{}

type BlockDTO struct {
	UID   int    `json:"uid"`
	Token string `json:"token"`
}

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

type Block struct {
	UID   int
	Token string
}

func (b Block) ToDTO() BlockDTO {
	return BlockDTO(b)
}
