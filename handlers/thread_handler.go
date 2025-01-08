package handlers

import (
	"aitu-moment/db/repository"
	"aitu-moment/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger = logrus.New()

func initThreadHandler() {

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
}

type ThreadHandler struct {
	repo *repository.ThreadRepository
}

func NewThreadHandler(repo *repository.ThreadRepository) *ThreadHandler {
	initThreadHandler()
	return &ThreadHandler{repo: repo}
}

func (h *ThreadHandler) GetThreads(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{
		"module":   "handlers",
		"function": "GetThreads",
		"action":   "handling request",
	})
	log.Info("Received request")
	threads, err := h.repo.GetAllThreads()
	if err != nil {
		log.Error("Error getting all threads: ", err, "error response")
		getErrorResponse(c, err.Error())
		return
	}
	log.Info("success, ", len(threads), " threads returned")
	c.HTML(http.StatusOK, "threads.html", gin.H{
		"threads": threads,
	})
}

func (h *ThreadHandler) CreateThread(c *gin.Context) {
	log := logger.WithFields(logrus.Fields{
		"module":   "handlers",
		"function": "createThread",
		"action":   "handling request",
	})
	log.Info("Received request")
	creatorId, err := strconv.Atoi(c.PostForm("creator_id"))
	parentThreadId, err := strconv.Atoi(c.PostForm("parent_thread_id"))
	thread := models.Thread{
		ThreadId:       int(creatorId),
		Content:        c.PostForm("content"),
		ParentThreadId: int(parentThreadId),
	}
	log.Warn(thread)
	threaedId, err := h.repo.CreateThread(thread)
	if err != nil {
		log.Error("Error getting all threads: ", err, "error response")
		getErrorResponse(c, err.Error())
		return
	}
	log.Info("success, ", threaedId, " created and returned")
	c.HTML(http.StatusOK, "thread.html", gin.H{
		"threads": thread,
	})
}
