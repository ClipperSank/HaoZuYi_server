package handlers

import (
	"fyno/server/internal/models"
	"fyno/server/internal/services"
)

type Handlers struct {
	Users      			models.UserHandlers
	SearchRecords 		models.SearchRecordHandlers
	Houses				models.HouseHandlers
	Contracts			models.ContractHandlers
	// Reviews 			   models.ReviewHandlers

	Posts			      models.PostHandlers
	Messages   			models.MessageHandlers
	WebSockets 			models.WebSocketHandlers
	Locations  			models.LocationHandlers
	Categories 			models.CategoryHandlers
	S3         			models.S3Handlers
}

func NewHandlers(serv *services.Services) *Handlers {
	return &Handlers{
		Users:      		NewUserHandlers(serv.Users),
		SearchRecords: 		NewSearchRecordHandlers(serv.SearchRecords),
		Houses:				NewHouseHandlers(serv.Houses),
		Contracts:			NewContractHandlers(serv.Contracts),
		// Reviews: 			  	NewReviewHandlers(serv.Reviews),
		
		S3:         		NewS3Handlers(serv.S3),
	}
}
