package models

type User struct {
    Id                  int    `db:"id"                  json:"id"                  form:"id"`
    Name                string `db:"username"            json:"name"                form:"username"`
    EducationalProgram  uint8  `db:"educational_program" json:"educational_program" form:"educational_program"`
    ProgramName         string `db:"program_name"        json:"program_name"        form:"program_name"`
    PublicName          string `db:"public_name"         json:"public_name"         form:"public_name"`
    Email               string `db:"email"               json:"email"               form:"email"`
    Passwd              string `db:"password"            json:"passwd"              form:"passwd"`
    Bio                 string `db:"bio"                 json:"bio"                 form:"bio"`
    Group               uint8  `db:"group_id"            json:"group"               form:"group"`
}


type EduProgram struct{
    Id          int             `db:"id" json:"id"`
    Name        string          `db:"name" json:"name"`
}


type Group struct{
    Id                  int    `db:"id"                  json:"id"                  form:"id"`
    EducationalProgram  uint8  `db:"educational_program_id" json:"educational_program" form:"educational_program"`
    Year                int16  `db:"year"                json:"year"                form:"year"`
    Number              uint8  `db:"number"              json:"number"              form:"number"`
    EduProgName         string `db:"name"                json:"name"                form:"name"`
    GroupName           string `form:"name"`
}
