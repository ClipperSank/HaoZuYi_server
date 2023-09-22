package services

import (
	"fyno/server/internal/models"

	"github.com/google/uuid"
)

type contractService struct {
	contractRepository models.ContractRepository
}

func NewContractService(pr models.ContractRepository) models.ContractService {
	return &contractService{
		contractRepository: pr,
	}
}

func (cs *contractService) GetAllContracts() ([]*models.Contract, error) {
	return cs.contractRepository.GetAll()
}

func (cs *contractService) GetContract(id uuid.UUID) (*models.Contract, error) {
	return cs.contractRepository.Get(id)
}

func (cs *contractService) CreateContract(c *models.Contract) (uuid.UUID, error) {
	return cs.contractRepository.Create(c)
}

func (cs *contractService) DeleteAllContracts() error {
	return cs.contractRepository.DeleteAll()
}
