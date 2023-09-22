package services

import (
	"fyno/server/internal/models"
	"fyno/server/internal/repositories"
)

type Services struct {
	Users      		models.UserService
	SearchRecords 	models.SearchRecordService
	Houses			models.HouseService
	Contracts		models.ContractService
	// Reviews 		   models.ReviewService

	S3         		models.S3Service
}

func NewServices(repositories *repositories.Repositories) *Services {
	return &Services{
		Users:     			NewUserService(repositories.Users),
		SearchRecords: 		NewSearchRecordService(repositories.SearchRecords),
		Houses:				NewHouseService(repositories.Houses),
		Contracts:			NewContractService(repositories.Contracts),
		// Reviews: 		      NewReviewService(repositories.ReviewServic),

		S3:					NewS3Service(),
	}
}

