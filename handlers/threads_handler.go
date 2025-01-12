package handlers

import (
	"aitu-moment/db/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ThreadHandler struct {
	repo *repository.UserRepository
}

func NewThreadHandler(repo *repository.UserRepository) *ThreadHandler {
	return &ThreadHandler{repo: repo}
}

func (h *ThreadHandler) GetThreads(c *gin.Context) {

	logThis := log.WithFields(logrus.Fields{
		"package":  "home_handler",
		"function": "GetThreads",
	})

	log.Info("Starting to fetch threads")

	filter := parseThreadFilter(c)
	logThis.WithFields(logrus.Fields{
		"page":       filter.Page,
		"page_size":  filter.PageSize,
		"search":     filter.Search,
		"creator_id": filter.CreatorID,
	}).Debug("Parsed thread filter")

	threads, totalCount, err := h.repo.FetchThreads(filter)
	if err != nil {
		logThis.WithError(err).Error("Failed to fetch threads")
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Failed to fetch threads",
		})
		return
	}

	logThis.WithFields(logrus.Fields{
		"thread_count": len(threads),
		"total_count":  totalCount,
	}).Info("Successfully fetched threads")

	totalPages := (totalCount + filter.PageSize - 1) / filter.PageSize
	hasNextPage := filter.Page < totalPages
	hasPrevPage := filter.Page > 1
	logThis.Info("Rendering full page response")

	responseMap := gin.H{
		"threads":     threads,
		"currentPage": filter.Page,
		"totalPages":  totalPages,
		"hasNextPage": hasNextPage,
		"prevPage":    max(filter.Page-1, 0),
		"nextPage":    min(filter.Page+1, totalPages),
		"hasPrevPage": hasPrevPage,
	}
	logThis.Info("Rendering filter page")
	c.HTML(http.StatusOK, "threads.html", responseMap)

	logThis.Info("Rendering HTMX partial response")
	c.HTML(http.StatusOK, "thread-list.html", responseMap)
}

func parseThreadFilter(c *gin.Context) repository.ThreadFilter {
	filter := repository.ThreadFilter{
		Page:     1,
		PageSize: 2,
	}

	if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil && page > 0 {
		filter.Page = page
	}
	if pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10")); err == nil && pageSize > 0 {
		filter.PageSize = pageSize
	}

	if creatorID, err := strconv.ParseInt(c.Query("creator_id"), 10, 64); err == nil {
		filter.CreatorID = &creatorID
	}

	if parentID, err := strconv.ParseInt(c.Query("parent_id"), 10, 64); err == nil {
		filter.ParentThreadID = &parentID
	}

	if startDate, err := time.Parse("2006-01-02", c.Query("start_date")); err == nil {
		filter.StartDate = &startDate
	}
	if endDate, err := time.Parse("2006-01-02", c.Query("end_date")); err == nil {
		filter.EndDate = &endDate
	}

	if minUpVotes, err := strconv.Atoi(c.Query("min_upvotes")); err == nil {
		filter.MinUpVotes = &minUpVotes
	}

	filter.Search = c.Query("search")
	order := c.Query("order")
	filter.Order = &order
	orderby := c.Query("order_by")
	filter.OrderBy = &orderby

	return filter
}
