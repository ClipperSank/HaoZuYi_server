package handlers

import (
	"fmt"
	"fyno/server/internal/models"
	"fyno/server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contractHandlers struct {
	contractService models.ContractService
}

func NewContractHandlers(cs models.ContractService) models.ContractHandlers {
	return &contractHandlers{
		contractService: cs,
	}
}

func (ch *contractHandlers) GetAllContracts(c *gin.Context) {
	contracts, err := ch.contractService.GetAllContracts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"contracts": contracts})
}

func (ch *contractHandlers) GetContract(c *gin.Context) {
	fmt.Println("GetContract")
	contractID := c.Param("id")
	fmt.Println("contractID", contractID)
	contract, err := ch.contractService.GetContract(utils.StringToUUID(contractID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"contract": contract})
}

func (ch *contractHandlers) CreateContract(c *gin.Context) {
	renterID := c.MustGet("renterID").(string)

	var input *models.Contract
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.RenterID = utils.StringToUUID(renterID)
	fmt.Println("input", input)
	contractID, err := ch.contractService.CreateContract(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"contractID": contractID})
}
