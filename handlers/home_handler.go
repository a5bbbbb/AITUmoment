package handlers

import (
	"aitu-moment/db/repository"
	"aitu-moment/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetHome(c *gin.Context) {
	users, err := h.repo.GetAllUsers()
	if err != nil {
		getErrorResponse(c, err.Error())
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name":  "Awesome",
		"users": users,
	})

}

func (h *UserHandler) SaveUser(c *gin.Context) {
	edu_prog_int := c.PostForm("educational_program")
	educational_program_int, _ := strconv.Atoi(edu_prog_int)

	user := models.User{
		Name:               c.PostForm("username"),
		EducationalProgram: uint8(educational_program_int),
		Program_name:       "",
	}
	id, err := h.repo.CreateUser(user)
	if err != nil {
		getErrorResponse(c, err.Error())
		return
	}

	savedUser, err := h.repo.GetUser(id)
	if err != nil {
		getErrorResponse(c, err.Error())
		return
	}

	c.HTML(http.StatusOK, "user.html", gin.H{
		"Name":         savedUser.Name,
		"Program_name": savedUser.Program_name,
	})

}

func getErrorResponse(c *gin.Context, errMessage string) {
	fmt.Println("ERRORRRR")
	c.HTML(
		http.StatusInternalServerError,
		"error.html", gin.H{
			"error": "There is some error happened:" + errMessage,
		})

}
