package model

type User struct{
	ID int `json:"id_usuario"`
	Name string `json:"nome_usuario"`
	Email string `json:"email_usuario"`
	Senha string `json:"senha_usuario"`
}