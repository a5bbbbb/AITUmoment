package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"aitu-moment/db"
	"aitu-moment/db/repository"
	"aitu-moment/handlers"
	"aitu-moment/models"
)

func main() {
	databaseConnection, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer databaseConnection.Close()

	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	userRepo := repository.NewUserRepository(databaseConnection.GetDB())
	userHandler := handlers.NewUserHandler(userRepo)
	e.GET("/", userHandler.GetHome)

	// e.POST("/users", func(c *gin.Context) {
	// 	edu_prog_int := c.PostForm("educational_program")
	// 	educational_program_int, _ := strconv.Atoi(edu_prog_int)
	//
	// 	user := models.User{
	// 		Name:               c.PostForm("username"),
	// 		EducationalProgram: uint8(educational_program_int),
	// 		Program_name:       "",
	// 	}
	// 	id, err := trySaveUser(*databaseConnection, user)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	user = *tryGetOne(*databaseConnection, id)
	// 	c.HTML(http.StatusOK, "user.html", gin.H{
	// 		"Name":         user.Name,
	// 		"Program_name": user.Program_name,
	// 	})
	// })

	e.POST("/users", userHandler.SaveUser)

	e.Run(":8080")
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

func tryGetOne(db db.Database, id int64) *models.User {
	var user models.User
	query := ` 
        SELECT u.*, ep.name as program_name 
        FROM users u 
        JOIN educational_programs ep ON u.educational_program = ep.id 
        WHERE u.id = $1`

	err := db.GetDB().Get(&user, query, id)
	if err != nil {
		log.Fatalln(err)
	}
	return &user
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
