package services

import (
	"fyno/server/internal/models"

	"github.com/google/uuid"
)

type searchrecordService struct {
	searchrecordRepository models.SearchRecordRepository
}

func NewSearchRecordService(srr models.SearchRecordRepository) models.SearchRecordService {
	return &searchrecordService{
		searchrecordRepository: srr,
	}
}

func (srs *searchrecordService) GetAllSearchRecords() ([]*models.SearchRecord, error) {
	return srs.searchrecordRepository.GetAll()
}

func (srs *searchrecordService) GetSearchRecord(id uuid.UUID) (*models.SearchRecord, error) {
	return srs.searchrecordRepository.Get(id)
}

func (srs *searchrecordService) CreateSearchRecord(sr *models.SearchRecord) (uuid.UUID, error) {
	return srs.searchrecordRepository.Create(sr)
}

func (srs *searchrecordService) DeleteAllSearchRecords() error {
	return srs.searchrecordRepository.DeleteAll()
}
