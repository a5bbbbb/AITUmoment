package repository

import (
	"aitu-moment/db"
	"aitu-moment/models"

	"github.com/jmoiron/sqlx"
)



type GroupRepo struct{
    db *sqlx.DB
}

func NewGroupRepo() *GroupRepo{
    return &GroupRepo{db: db.GetDB()}
}

func (r *GroupRepo) GetGroups(eduProg uint8) ([]models.Group, error){
    query := `  SELECT g.id, g.year, g.number, ep.name
                FROM groups g
                JOIN educational_programs ep ON g.educational_program_id = ep.id
                WHERE g.educational_program_id = $1
             `
	var groups []models.Group
	err := r.db.Select(&groups, query, eduProg)

	return groups, err
}


func (r *GroupRepo) GetGroupByID(groupID int) (*models.Group, error){
    query := 
    `
    SELECT g.id, g.year, g.number, ep.name, g.educational_program_id
    FROM groups g
    JOIN educational_programs ep ON g.educational_program_id = ep.id
    WHERE g.id = $1
    `

    var group models.Group
    err := r.db.Get(&group, query, groupID)

    return &group, err
}

