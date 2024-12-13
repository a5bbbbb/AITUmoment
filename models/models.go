package models

type User struct {
	Id                 int    `db:"id"                  json:"id"`
	Name               string `db:"username"            json:"name"`
	EducationalProgram uint8  `db:"educational_program" json:"educational_program"`
	Program_name       string `db:"program_name"        json:"program_name"`
}
