 package handlers
//
// import (
// 	"aitu-moment/db/repository"
// 	"aitu-moment/models"
// 	"net/http"
// 	"net/smtp"
// 	"os"
// 	"strconv"
// 	"time"
//
// 	"github.com/gin-gonic/gin"
// 	"github.com/sirupsen/logrus"
// )
//
// var log = logrus.New()
//
// func init() {
//
// 	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err == nil {
// 		log.Out = file
// 	} else {
// 		log.Info("Failed to log to file, using default stderr")
// 	}
// }
//
// type Thread struct {
// 	ID             int64     `db:"thread_id" json:"id"`
// 	CreatorID      int64     `db:"creator_id" json:"creator_id"`
// 	Content        string    `db:"content" json:"content"`
// 	CreateDate     time.Time `db:"create_date" json:"create_date"`
// 	UpVotes        int       `db:"up_votes" json:"up_votes"`
// 	ParentThreadID *int64    `db:"parent_thread_id" json:"parent_thread_id,omitempty"`
// }
//
// type UserHandler struct {
// 	repo *repository.UserRepository
// }
//
// func NewUserHandler(repo *repository.UserRepository) *UserHandler {
// 	return &UserHandler{repo: repo}
// }
//
// func (h *UserHandler) GetHome(c *gin.Context) {
// 	users, err := h.repo.GetAllUsers()
// 	if err != nil {
// 		log.WithFields(logrus.Fields{
// 			"scope":   "handler",
// 			"message": err.Error(),
// 		}).Error("ERROR!")
// 		getErrorResponse(c, err.Error())
// 		return
// 	}
// 	log.WithFields(logrus.Fields{
// 		"scope":   "handler",
// 		"message": "Successfully got home screen",
// 	}).Error("ERROR!")
// 	c.HTML(http.StatusOK, "index.html", gin.H{
// 		"name":  "Awesome",
// 		"users": users,
// 	})
//
// }
//
// func (h *UserHandler) GetThreads(c *gin.Context) {
//
// 	log.Info("Starting to fetch threads")
//
// 	filter := parseThreadFilter(c)
// 	log.WithFields(logrus.Fields{
// 		"page":       filter.Page,
// 		"page_size":  filter.PageSize,
// 		"search":     filter.Search,
// 		"creator_id": filter.CreatorID,
// 	}).Debug("Parsed thread filter")
//
// 	threads, totalCount, err := h.repo.FetchThreads(filter)
// 	if err != nil {
// 		log.WithError(err).Error("Failed to fetch threads")
// 		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
// 			"error": "Failed to fetch threads",
// 		})
// 		return
// 	}
//
// 	log.WithFields(logrus.Fields{
// 		"thread_count": len(threads),
// 		"total_count":  totalCount,
// 	}).Info("Successfully fetched threads")
//
// 	totalPages := (totalCount + filter.PageSize - 1) / filter.PageSize
// 	hasNextPage := filter.Page < totalPages
// 	hasPrevPage := filter.Page > 1
// 	log.Info("Rendering full page response")
//
// 	c.HTML(http.StatusOK, "threads.html", gin.H{
// 		"threads":     threads,
// 		"currentPage": filter.Page,
// 		"totalPages":  totalPages,
// 		"hasNextPage": hasNextPage,
// 		"hasPrevPage": hasPrevPage,
// 		"prevPage":    max(filter.Page-1, 0),
// 		"nextPage":    min(filter.Page+1, totalPages),
// 		"filter":      filter,
// 	})
//
// 	log.Info("Rendering HTMX partial response")
// 	c.HTML(http.StatusOK, "thread-list.html", gin.H{
// 		"threads":     threads,
// 		"currentPage": filter.Page,
// 		"totalPages":  totalPages,
// 		"hasNextPage": hasNextPage,
// 		"prevPage":    max(filter.Page-1, 0),
// 		"nextPage":    min(filter.Page+1, totalPages),
// 		"hasPrevPage": hasPrevPage,
// 	})
//
// }
//
// func parseThreadFilter(c *gin.Context) repository.ThreadFilter {
// 	filter := repository.ThreadFilter{
// 		Page:     1,
// 		PageSize: 10,
// 	}
//
// 	if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil && page > 0 {
// 		filter.Page = page
// 	}
// 	if pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10")); err == nil && pageSize > 0 {
// 		filter.PageSize = pageSize
// 	}
//
// 	if creatorID, err := strconv.ParseInt(c.Query("creator_id"), 10, 64); err == nil {
// 		filter.CreatorID = &creatorID
// 	}
//
// 	if parentID, err := strconv.ParseInt(c.Query("parent_id"), 10, 64); err == nil {
// 		filter.ParentThreadID = &parentID
// 	}
//
// 	if startDate, err := time.Parse("2006-01-02", c.Query("start_date")); err == nil {
// 		filter.StartDate = &startDate
// 	}
// 	if endDate, err := time.Parse("2006-01-02", c.Query("end_date")); err == nil {
// 		filter.EndDate = &endDate
// 	}
//
// 	if minUpVotes, err := strconv.Atoi(c.Query("min_upvotes")); err == nil {
// 		filter.MinUpVotes = &minUpVotes
// 	}
//
// 	filter.Search = c.Query("search")
//
// 	return filter
// }
//
// func (h *UserHandler) SaveUser(c *gin.Context) {
// 	edu_prog_int := c.PostForm("educational_program")
// 	educational_program_int, _ := strconv.Atoi(edu_prog_int)
//
// 	user := models.User{
// 		Name:               c.PostForm("username"),
// 		EducationalProgram: uint8(educational_program_int),
// 		Program_name:       "",
// 	}
// 	id, err := h.repo.CreateUser(user)
// 	if err != nil {
// 		log.WithFields(logrus.Fields{
// 			"scope":   "handler",
// 			"message": err.Error(),
// 		}).Error("ERROR!")
// 		getErrorResponse(c, err.Error())
// 		return
// 	}
//
// 	savedUser, err := h.repo.GetUser(id)
// 	if err != nil {
// 		log.WithFields(logrus.Fields{
// 			"scope":   "handler",
// 			"message": err.Error(),
// 		}).Error("ERROR!")
// 		getErrorResponse(c, err.Error())
// 		return
// 	}
// 	log.WithFields(logrus.Fields{
// 		"scope":   "handler",
// 		"message": "Successfully saved user"})
//
// 	c.HTML(http.StatusOK, "user.html", gin.H{
// 		"Name":         savedUser.Name,
// 		"Program_name": savedUser.Program_name,
// 	})
//
// }
//
// func (h *UserHandler) SendMail(c *gin.Context) {
//
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"
//
// 	from := "suprunoviktor@gmail.com"
// 	pass := "nfrf gnuf ndtg mvoh"
//
// 	email := "seraf5@bk.ru"
// 	log.Info(email)
// 	msg := c.PostForm("message")
//
// 	auth := smtp.PlainAuth("", from, pass, smtpHost)
//
// 	logrus.WithFields(logrus.Fields{
// 		"from":    from,
// 		"to":      email,
// 		"subject": msg,
// 	}).Info("Sending email started")
//
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(msg))
// 	if err != nil {
// 		logrus.WithFields(logrus.Fields{
// 			"from":    from,
// 			"to":      email,
// 			"subject": msg,
// 			"error":   err.Error(),
// 		}).Error("Failed to send email")
// 		getErrorResponse(c, err.Error())
// 		return
// 	}
//
// 	log.WithFields(logrus.Fields{
// 		"from":    from,
// 		"to":      email,
// 		"subject": msg,
// 	}).Info("Email sent successfully")
//
// 	c.HTML(http.StatusOK, "email.html", gin.H{})
// }
//
// func (h *UserHandler) GetMail(c *gin.Context) {
// 	c.HTML(http.StatusOK, "email.html", gin.H{})
//
// }
//
// func getErrorResponse(c *gin.Context, errMessage string) {
// 	log.WithFields(logrus.Fields{
// 		"scope":   "Database",
// 		"message": errMessage,
// 	}).Error("ERROR")
//
// 	c.HTML(
// 		http.StatusInternalServerError,
// 		"error.html", gin.H{
// 			"error": "There is some error happened:" + errMessage,
// 		})
//
// }
