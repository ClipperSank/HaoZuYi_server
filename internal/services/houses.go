package services

import (
	"fyno/server/internal/models"

	"github.com/google/uuid"
)

type houseService struct {
	houseRepository models.HouseRepository
}

func NewHouseService(hr models.HouseRepository) models.HouseService {
	return &houseService{
		houseRepository: hr,
	}
}

func (ps *houseService) GetAllHouses() ([]*models.House, error) {
	return ps.houseRepository.GetAll()
}

func (ps *houseService) GetHouse(id uuid.UUID) (*models.House, error) {
	return ps.houseRepository.Get(id)
}

func (ps *houseService) CreateHouse(p *models.House) (uuid.UUID, error) {
	return ps.houseRepository.Create(p)
}

func (ps *houseService) CreateHouseImage(p []models.HouseImage, houseID uuid.UUID) error {
	for _, v := range p {
		id := uuid.New()
		err := ps.houseRepository.CreateHouseImage(id, v, houseID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ps *houseService) DeleteAllHouses() error {
	return ps.houseRepository.DeleteAll()
}
