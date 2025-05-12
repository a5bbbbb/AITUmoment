package models

import "time"

type User struct {
	Id                 int    `db:"id"`
	Name               string `db:"username"`
	EducationalProgram uint8  `db:"educational_program"`
	ProgramName        string `db:"program_name"`
	PublicName         string `db:"public_name"`
	Email              string `db:"email"`
	Passwd             string `db:"password"`
	Bio                string `db:"bio"`
	Group              uint8  `db:"group_id"`
	Verified           bool   `db:"verified"`
}

type EduProgram struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type Group struct {
	Id                 int    `db:"id"`
	EducationalProgram uint8  `db:"educational_program_id"`
	Year               int16  `db:"year"`
	Number             uint8  `db:"number"`
	EduProgName        string `db:"name"`
	GroupName          string `form:"name"`
}

type Thread struct {
	Id           int       `db:"thread_id"`
	CreatorID    int       `db:"creator_id"`
	CreatorName  string    `db:"creator_name"`
	Content      string    `db:"content"`
	CreateDate   time.Time `db:"create_date"`
	UpVotes      int       `db:"up_votes"`
	ParentThread *int      `db:"parent_thread_id"`
	UserUpvoted  bool      `db:"has_upvote"`
}

type EmailVerification struct {
	Email            string
	PublicName       string
	VerificationLink string
}
