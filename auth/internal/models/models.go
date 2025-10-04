package models

//TODO: add fields
//TODO: move to right places
type AuthDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterDTO struct {
	TgID     int    `json:"tg_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokensDTO struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type BlockDTO struct {
	UID   int    `json:"uid"`
	Token string `json:"token"`
}

type Auth struct {
	Login    string
	Password string
}

type UserPassword struct {
	UID  int
	Hash string
}

func (a Auth) ToDTO() AuthDTO {
	return AuthDTO{}
}

type Register struct {
	TgID     int
	Login    string
	Password string
}

func (r Register) ToDTO() RegisterDTO {
	return RegisterDTO{}
}

type Tokens struct {
	Access  string
	Refresh string
}

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
