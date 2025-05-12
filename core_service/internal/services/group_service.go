package services

import (
	"fmt"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"
)

type GroupService struct {
	groupRepo GroupRepository
}

func NewGroupService(groupRepo GroupRepository) *GroupService {
	logger.GetLogger().Trace("Creating group service")
	return &GroupService{groupRepo: groupRepo}
}

func (s *GroupService) GetGroups(eduProg uint8) (*[]models.Group, error) {
	groups, err := s.groupRepo.GetGroups(eduProg)

	if err != nil {
		return nil, err
	}

	for i := range *groups {
		(*groups)[i].GroupName = generateGroupName(
			eduProg,
			(*groups)[i].Year,
			(*groups)[i].Number,
		)
	}

	return groups, err
}

func (s *GroupService) GetGroup(groupID int) (*models.Group, error) {
	group, err := s.groupRepo.GetGroupByID(groupID)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting group in service: %v", err.Error())
		return nil, err
	}

	group.GroupName = generateGroupName(group.EducationalProgram, group.Year, group.Number)

	return group, nil
}

func generateGroupName(eduProg uint8, year int16, number uint8) string {
	programAbbreviations := map[uint8]string{
		1:  "SE",
		2:  "CS",
		3:  "BDA",
		4:  "MT",
		5:  "MCS",
		6:  "BDH",
		7:  "CS",
		8:  "ST",
		9:  "IoT",
		10: "EE",
		11: "ITM",
		12: "ITE",
		13: "AIB",
		14: "DJ",
	}

	abbr, exists := programAbbreviations[eduProg]
	if !exists {
		return "UNKNOWN"
	}

	yearShort := year % 100

	// Format the group name: ABBR-YYNN
	// where ABBR is the program abbreviation
	// YY is the last two digits of the year
	// NN is the group number
	return fmt.Sprintf("%s-%02d%02d", abbr, yearShort, number)
}
