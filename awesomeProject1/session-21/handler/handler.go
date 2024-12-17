package handler

import (
	"github.com/agshinaliyev/ms-app/db"
	"github.com/agshinaliyev/ms-app/service"
	"net/http"
)

type ApplicationHandler struct {
	AppService service.RateService
}

func NewApplicationHandler() *ApplicationHandler {
	repo := db.NewRepository(db.GetDb())
	return &ApplicationHandler{AppService: service.NewRateService(repo)}
}

func (app ApplicationHandler) GetRates(w http.ResponseWriter, r *http.Request) {

}
func (app ApplicationHandler) SetRate(w http.ResponseWriter, r *http.Request) {

}
