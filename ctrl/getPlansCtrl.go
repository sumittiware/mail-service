package ctrl

import data "mail-service/models"

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
	// TODO : add the get all plans logic

	// if err != nil {
	// 	return GetPlansResponseBody{
	// 		Error: err,
	// 	}
	// }

	// return GetPlansResponseBody{
	// 	Plans: plans,
	// }

	return GetPlansResponseBody{
		Error: nil,
	}
}
