package model

type User struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Active    bool   `json:"active"`
}

var Users = []User{
	{1, "Renan", "Roberto", "renanroberto1@gmail.com", "123456789", true},
	{2, "Jamilly", "Roberta", "jamillyroberta1@gmail.com", "123456789", true},
}
