package handlers

import (
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/http/server/handlers/utils"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/services"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThreadsHandler struct {
	userService   *services.UserService
	eduService    *services.EduService
	groupService  *services.GroupService
	threadService *services.ThreadService
}

func NewThreadsHandler(
	userService *services.UserService,
	eduService *services.EduService,
	groupService *services.GroupService,
	threadService *services.ThreadService,
) *ThreadsHandler {
	return &ThreadsHandler{
		userService:   userService,
		eduService:    eduService,
		groupService:  groupService,
		threadService: threadService,
	}
}

func (h *ThreadsHandler) FeedPage(c *gin.Context) {
	userID, err := utils.GetUserFromClaims(c)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting claims in main page %v", err.Error())
		c.HTML(http.StatusInternalServerError, "index.html", nil)
		return
	}

	user, _, _, err := h.userService.GetFullUserInfo(*userID)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting userInfo in main page %v", err.Error())
		c.HTML(http.StatusInternalServerError, "index.html", nil)
		return
	}

	threads, err := h.threadService.GetParentThreads(userID)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting threads in main page %v", err.Error())
		c.HTML(http.StatusInternalServerError, "home.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"name":    user.Name,
		"threads": threads,
	})

}

func (h *ThreadsHandler) GetSubThreads(c *gin.Context) {

	userID, err := utils.GetUserFromClaims(c)

	if err != nil {
		logger.GetLogger().Error("Error during getting claims in getting sub threads")
		c.HTML(http.StatusBadRequest, "createThreadPage.html", gin.H{
			"error": "Error during getting claims in getting sub threads",
		})
		return
	}

	threadIdStr := c.Query("threadID")

	threadID, err := strconv.Atoi(threadIdStr)

	if err != nil {
		logger.GetLogger().Errorf("Error during parsing threadID in getting sub threads %v", err.Error())
		c.HTML(http.StatusBadRequest, "home.html", gin.H{
			"error": "Invalid threadID",
		})
		return
	}

	subThreads, parentThread, err := h.threadService.GetSubThreads(threadID, *userID)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting sub threads %v", err.Error())
		c.HTML(http.StatusInternalServerError, "home.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "threadPage.html", gin.H{
		"parentThread": parentThread,
		"subThreads":   subThreads,
	})

}

func (h *ThreadsHandler) CreateThreadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "createThreadPage.html", nil)
}

func (h *ThreadsHandler) SaveThread(c *gin.Context) {
	content, exists := c.GetPostForm("content")
	if !exists {
		logger.GetLogger().Error("No content during thread creation")
		c.HTML(http.StatusBadRequest, "createThreadPage.html", gin.H{
			"error": "No content during thread creation",
		})
		return
	}

	var parentThreadID *int
	var err error

	parentThreadStr, exists := c.GetPostForm("parentThreadID")
	if exists && parentThreadStr != "" {
		val, err := strconv.Atoi(parentThreadStr)
		if err != nil {
			logger.GetLogger().Error("Error getting parentThread in creating thread")
			c.HTML(http.StatusBadRequest, "createThreadPage.html", gin.H{
				"error": "Invalid parentThreadID",
			})
			return
		}
		parentThreadID = &val
	}

	userID, err := utils.GetUserFromClaims(c)

	if err != nil {
		logger.GetLogger().Error("Error during getting claims in creating thread")
		c.HTML(http.StatusBadRequest, "createThreadPage.html", gin.H{
			"error": "Error during getting claims in creating thread",
		})
		return
	}

	newThreadID, err := h.threadService.SaveThread(content, *userID, parentThreadID)

	if err != nil {
		logger.GetLogger().Errorf("Error during creating thread, %v", err.Error())
		c.HTML(http.StatusBadRequest, "createThreadPage.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	var threadID int
	if parentThreadID == nil {
		threadID = *newThreadID
	} else {
		threadID = *parentThreadID
	}

	subThreads, parentThread, err := h.threadService.GetSubThreads(threadID, *userID)
	if err != nil {
		logger.GetLogger().Errorf("Error during getting subthreads and parent thread, %v", err.Error())
		c.HTML(http.StatusBadRequest, "createThreadPage.html", gin.H{
			"error": err.Error(),
		})
	}

	c.HTML(http.StatusOK, "threadPage.html", gin.H{
		"parentThread": parentThread,
		"subThreads":   subThreads,
	})
}

func (h *ThreadsHandler) SaveUpvote(c *gin.Context) {

	userID, err := utils.GetUserFromClaims(c)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting claims in main page %v", err.Error())
		c.HTML(http.StatusInternalServerError, "index.html", nil)
		return
	}

	threadIdStr := c.Query("threadID")
	threadID, err := strconv.Atoi(threadIdStr)

	if err != nil {
		logger.GetLogger().Errorf("Error during parsing threadID in saving upvote %v", err.Error())
		c.HTML(http.StatusBadRequest, "home.html", gin.H{
			"error": "Invalid threadID",
		})
		return
	}

	upvoteStr := c.Query("upvote")
	upvote, err := strconv.ParseBool(upvoteStr)

	if err != nil {
		logger.GetLogger().Errorf("Error during parsing upvote in saving upvote %v", err.Error())
		c.HTML(http.StatusBadRequest, "home.html", gin.H{
			"error": "Invalid upvote",
		})
		return
	}

	err = h.threadService.SaveUpvote(&threadID, userID, &upvote)

	if err != nil {
		logger.GetLogger().Errorf("Error during saving upvote %v", err.Error())
		c.HTML(http.StatusBadRequest, "home.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	thread, err := h.threadService.GetThread(threadID, *userID)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting thread in saving upvote %v", err.Error())
		c.HTML(http.StatusBadRequest, "home.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "upvote.html", gin.H{
		"Id":          thread.Id,
		"UserUpvoted": thread.UserUpvoted,
		"UpVotes":     thread.UpVotes,
	})

}
