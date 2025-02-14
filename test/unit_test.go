package tests

import (
	"aitu-moment/models"
	"aitu-moment/services"
	"encoding/json"
	"fmt"
	"testing"

	"aitu-moment/db"
)

func TestEduProgServiceUnitTest(t *testing.T) {

	defer db.Close()
	userService := services.NewUserService()
	created_user, err := userService.CreateUser(&models.User{
		Name:               "unit_test_user",
		EducationalProgram: 1,
		Email:              "unit_test_user@mail.kz",
		Passwd:             "unit_test_passwd",
		PublicName:         "public_unit_test_user_name",
		Group:              1,
	})
	js, _ := json.Marshal(created_user)
	fmt.Println(string(js))
	if err != nil {
		t.Errorf("Error creating user: %s", err)
	}
}
