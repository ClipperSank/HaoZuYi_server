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

func (hs *houseService) GetAllHouses(user_id uuid.UUID) ([]*models.House, error) {
	return hs.houseRepository.GetAll(user_id)
}

func (hs *houseService) GetHouse(house_id uuid.UUID) (*models.House, error) {
	return hs.houseRepository.Get(house_id)
}

func (hs *houseService) CreateHouse(h *models.House) (uuid.UUID, error) {
	return hs.houseRepository.Create(h)
}

func (hs *houseService) CreateHouseImage(h []models.HouseImage, houseID uuid.UUID) error {
	for _, v := range h {
		id := uuid.New()
		err := hs.houseRepository.CreateHouseImage(id, v, houseID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (hs *houseService) DeleteAllHouses() error {
	return hs.houseRepository.DeleteAll()
}
