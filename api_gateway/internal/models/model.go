package models

import "time"

type User struct {
	Id                 int    `json:"id"                  form:"id"`
	Name               string `json:"name"                form:"username"`
	EducationalProgram uint8  `json:"educational_program" form:"educational_program"`
	ProgramName        string `json:"program_name"        form:"program_name"`
	PublicName         string `json:"public_name"         form:"public_name"`
	Email              string `json:"email"               form:"email"`
	Passwd             string `json:"passwd"              form:"passwd"`
	Bio                string `json:"bio"                 form:"bio"`
	Group              uint8  `json:"group"               form:"group"`
	Verified           bool   `json:"verified"`
}

type EduProgram struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	Id                 int    `json:"id"                  form:"id"`
	EducationalProgram uint8  `json:"educational_program" form:"educational_program"`
	Year               int16  `json:"year"                form:"year"`
	Number             uint8  `json:"number"              form:"number"`
	EduProgName        string `json:"name"                form:"name"`
	GroupName          string `form:"name"`
}

type Thread struct {
	Id           int       `json:"id"           	form:"id"`
	CreatorID    int       `json:"creator"      	form:"creator"`
	CreatorName  string    `json:"creator_name"    	form:"creator_name"`
	Content      string    `json:"content"      	form:"content"`
	CreateDate   time.Time `json:"created_data" 	form:"created_data"`
	UpVotes      int       `json:"up_votes"    		form:"up_votes"`
	ParentThread *int      `json:"parent_thread" 	form:"parent_thread"`
	UserUpvoted  bool      `json:"hasUpvote" 		form:"hasUpvote"`
}
