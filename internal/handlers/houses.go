package handlers

import (
	"fmt"
	"fyno/server/internal/models"
	"fyno/server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type houseHandlers struct {
	houseService models.HouseService
}

func NewHouseHandlers(hs models.HouseService) models.HouseHandlers {
	return &houseHandlers{
		houseService: hs,
	}
}

func (hh *houseHandlers) GetAllHouses(c *gin.Context) {
	houses, err := hh.houseService.GetAllHouses()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"houses": houses})
}

func (hh *houseHandlers) GetHouse(c *gin.Context) {
	fmt.Println("GetHouse")
	houseID := c.Param("id")
	fmt.Println("houseID", houseID)
	house, err := hh.houseService.GetHouse(utils.StringToUUID(houseID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"house": house})
}

func (hh *houseHandlers) CreateHouse(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	var input *models.House
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = utils.StringToUUID(userID)
	fmt.Println("input", input)
	houseID, err := hh.houseService.CreateHouse(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = hh.houseService.CreateHouseImage(input.HouseImage, input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"houseID": houseID})
}
