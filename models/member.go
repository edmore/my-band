package models

type Member struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Speciality string `json:"speciality"`
}
