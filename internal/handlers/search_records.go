package handlers

import (
	"fmt"
	"fyno/server/internal/models"
	"fyno/server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type searchrecordHandlers struct {
	searchrecordService models.SearchRecordService
}

func NewSearchRecordHandlers(srs models.SearchRecordService) models.SearchRecordHandlers {
	return &searchrecordHandlers{
		searchrecordService: srs,
	}
}

func (srh *searchrecordHandlers) GetAllSearchRecords(c *gin.Context) {
	searchrecords, err := srh.searchrecordService.GetAllSearchRecords()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"searchrecords": searchrecords})
}

func (srh *searchrecordHandlers) GetSearchRecord(c *gin.Context) {
	fmt.Println("GetSearchRecord")
	searchrecordID := c.Param("id")
	fmt.Println("searchrecordID", searchrecordID)
	searchrecord, err := srh.searchrecordService.GetSearchRecord(utils.StringToUUID(searchrecordID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"searchrecord": searchrecord})
}

func (srh *searchrecordHandlers) CreateSearchRecord(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	var input *models.SearchRecord
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = utils.StringToUUID(userID)
	fmt.Println("input", input)
	searchrecordID, err := srh.searchrecordService.CreateSearchRecord(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"searchrecordID": searchrecordID})
}
