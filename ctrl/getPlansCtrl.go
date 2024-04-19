package ctrl

import (
	"log"
	"mail-service/config"
	data "mail-service/models"
)

type GetPlansRequestBody struct {
}

type GetPlansResponseBody struct {
	Plans []*data.Plan `json:"plans"`
	Error error        `json:"error"`
}

type GetPlansFlow struct {
	RequestBody GetPlansRequestBody
}

func (f *GetPlansFlow) Run() GetPlansResponseBody {
	// validate the request body
	if err := f.validate(); err != nil {
		return GetPlansResponseBody{
			Error: err,
		}
	}

	return f.do()
}

func (f *GetPlansFlow) validate() error {
	// TODO : validate the request body
	return nil
}

func (f *GetPlansFlow) do() GetPlansResponseBody {
	plans, err := config.ApplicationConfig.Models.Plan.GetAll()

	if err != nil {
		log.Println("Error getting plans: ", err)
		return GetPlansResponseBody{
			Error: err,
		}
	}
	return GetPlansResponseBody{
		Plans: plans,
		Error: nil,
	}
}
