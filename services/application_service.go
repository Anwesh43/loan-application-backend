package services

import (
	"com.loan.demo/daos"
	"com.loan.demo/models"
	"gorm.io/gorm"
)

type IApplicationService interface {
	CreateApplication(daos.InitiateApplicationRequest) int
	FetchApplications() []models.Application
}

type ApplicationService struct {
	db *gorm.DB
}

func (applicationService ApplicationService) CreateApplication(request daos.InitiateApplicationRequest) int {

	application := &models.Application{
		Name:            request.Name,
		YearEstablished: request.YearEstablished,
		Status:          "Initiated",
	}
	applicationService.db.Create(application)
	return application.Id
}

func (applicationService ApplicationService) FetchApplications() []models.Application {
	applications := make([]models.Application, 0)
	applicationService.db.Find(&applications)
	return applications
}

func NewApplicationService(db *gorm.DB) IApplicationService {
	return ApplicationService{
		db: db,
	}
}
