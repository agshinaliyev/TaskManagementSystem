package service

import "github.com/agshinaliyev/ms-app/db"

type RateService interface {
	GetRates()
}

type rateService struct {
	repo db.Repository
}

func NewRateService(repo db.Repository) RateService {
	return &rateService{repo: repo}
}
func (r rateService) GetRates() {

}
