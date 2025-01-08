package models

import "time"

type User struct {
	Id                 int    `db:"id"                  json:"id"`
	Name               string `db:"username"            json:"name"`
	EducationalProgram uint8  `db:"educational_program" json:"educational_program"`
	Program_name       string `db:"program_name"        json:"program_name"`
}

type Thread struct {
	ThreadId       int       `db:"thread_id"		json:"thread_id"`
	CreatorId      int       `db:"creator_id" 		json:"creator_id"`
	Content        string    `db:"content" 			json:"content"`
	CreateDate     time.Time `db:"create_date" 		json:"create_date"`
	UpVotes        int       `db:"up_votes" 		json:"up_votes"`
	ParentThreadId int       `db:"parent_thread_id" json:"parent_thread_id"`
}
