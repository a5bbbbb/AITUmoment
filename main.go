package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"aitu-moment/db"
	"aitu-moment/models"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func receiveMessages(c *gin.Context) {

	jsonData, err := c.GetRawData()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Invalid request body.",
		})
		fmt.Println("Couldn't read request body:", err.Error())
		return
	}

	req := make(map[string]string)

	if err = json.Unmarshal(jsonData, &req); err != nil {

		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": `Invalid JSON: JSON must have only "message" field and its value of type string. Check missing quotes and braces.`,
		})

		fmt.Println("Couldn't map data from json: ")
		fmt.Println(err.Error())

		return
	}

	if _, ok := req["message"]; !ok {

		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": `Invalid JSON: absent "message" field.`,
		})

		fmt.Println("No message field: ")
		fmt.Println(req)
		return
	}

	if len(req) != 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": `Invalid JSON: redundant fields, only "message" field is allowed.`,
		})
		fmt.Println("Processed request and raw request don't match: ")
		fmt.Println(req)
		return
	}

	fmt.Println("Received message: ", req["message"])

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data successfully received",
	})
}

func main() {
	databaseConnection, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer databaseConnection.Close()

	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	e.GET("/", func(c *gin.Context) {
		users := tryGet(*databaseConnection)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name":  "Awesome",
			"users": users,
		})
	})

	e.POST("/users", func(c *gin.Context) {
		edu_prog_int := c.PostForm("educational_program")
		educational_program_int, _ := strconv.Atoi(edu_prog_int)

		user := models.User{
			Name:               c.PostForm("username"),
			EducationalProgram: uint8(educational_program_int),
			Program_name:       "",
		}
		id, err := trySaveUser(*databaseConnection, user)
		if err != nil {
			fmt.Println(err.Error())
		}

		var userSaved *models.User

		userSaved, err = tryGetUser(*databaseConnection, uint64(id))

		c.HTML(http.StatusOK, "user.html", gin.H{
			"Id":           userSaved.Id,
			"Name":         userSaved.Name,
			"Program_name": userSaved.Program_name,
		})
	})

	e.GET("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.Header("HX-Reswap", "none")
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": `Invalid id: only numbers are allowed`,
			})
			log.Panic(err)
			return
		}

		user, err := tryGetUser(*databaseConnection, uint64(id))

		if err != nil {
			c.Header("HX-Reswap", "none")
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": `Invalid id: user with this id does not exist`,
			})
			log.Panic(err)
			return
		}

		c.HTML(http.StatusOK, "user.html", gin.H{
			"Name":         user.Name,
			"Program_name": user.Program_name,
			"Id":           user.Id,
		})
	})

	e.PUT("users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.Header("HX-Reswap", "none")
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": `Invalid id: only numbers are allowed`,
			})
			log.Panic(err)
			return
		}

		user, err := tryGetUser(*databaseConnection, uint64(id))

		if err != nil {
			c.Header("HX-Reswap", "none")
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": `Invalid id: user with this id does not exist`,
			})
			log.Panic(err)
			return
		}

		ed_prog, err := strconv.Atoi(c.PostForm("educational_program"))

		if c.PostForm("educational_program") == "" {
			ed_prog = int(user.EducationalProgram)
			err = nil
		}

		if err != nil {
			c.Header("HX-Reswap", "none")
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": `Invalid educational program: it must be within [1, 14] decimal range`,
			})
			log.Panic(err)
			return
		}

		newUser := models.User{
			Id:                 id,
			Name:               c.PostForm("username"),
			EducationalProgram: uint8(ed_prog),
			Program_name:       "",
		}

		if newUser.Name == "" {
			newUser.Name = user.Name
		}

		err = tryUpdateUser(*databaseConnection, newUser)

		if err != nil {
			c.Header("HX-Reswap", "none")
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"status":  "fail",
				"message": `Failed to update user: internal error`,
			})
			log.Panic(err)
			return
		}

		user, _ = tryGetUser(*databaseConnection, uint64(id))

		c.HTML(http.StatusOK, "user.html", gin.H{
			"Name":         user.Name,
			"Program_name": user.Program_name,
			"Id":           user.Id,
		})

	})

	e.DELETE("users/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": `Invalid id: only numbers are allowed`,
			})
			log.Panic(err)
			return
		}

		_, err = tryGetUser(*databaseConnection, uint64(id))

		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": `Invalid id: user with this id does not exist`,
			})
			log.Panic(err)
			return
		}

		err = tryDeleteUser(*databaseConnection, uint64(id))

		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": `Internal error: failed to delete user`,
			})
			log.Println(err)
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "User successfully deleted",
		})
	})

	e.GET("/a1", receiveMessages)

	e.POST("/a1", receiveMessages)

	e.Run(":8080")
}

func tryDeleteUser(db db.Database, id uint64) error {

	err := db.GetDB().QueryRow(
		"DELETE FROM users WHERE id=$1",
		id,
	)

	if err != nil {
		return err.Err()
	}

	return nil
}

func tryUpdateUser(db db.Database, user models.User) error {

	err := db.GetDB().QueryRow(
		"UPDATE users SET username=$1, educational_program=$2 WHERE id=$3",
		user.Name, user.EducationalProgram, user.Id,
	)

	if err != nil {
		return err.Err()
	}

	return nil
}

func trySaveUser(db db.Database, user models.User) (int64, error) {
	var insertedID int
	err := db.GetDB().QueryRow(
		"INSERT INTO users (username, educational_program) VALUES ($1, $2) RETURNING id",
		user.Name, user.EducationalProgram,
	).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return int64(insertedID), nil
}

func tryGetUser(db db.Database, id uint64) (*models.User, error) {
	var user models.User
	query := ` 
        SELECT u.*, ep.name as program_name 
        FROM users u 
        JOIN educational_programs ep ON u.educational_program = ep.id 
        WHERE u.id = $1`

	err := db.GetDB().Get(&user, query, id)
	return &user, err
}

func tryGet(db db.Database) []models.User {
	rows, _ := db.GetDB().
		Queryx(" SELECT u.*, ep.name as program_name FROM users u JOIN educational_programs ep ON u.educational_program = ep.id")
	defer rows.Close()

	users := make([]models.User, 0)

	for rows.Next() {
		var user models.User
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatalf("Got some error during reading through resultset, %s", err.Error())
		}

		users = append(users, user)
	}
	return users
}
