package ctrl

type SubscribePlanRequestBody struct {
}

type SubscribePlanResponseBody struct {
	Status string `json:"status"`
	Error  error  `json:"error"`
}

type SubscribePlanFlow struct {
	RequestBody SubscribePlanRequestBody
	UserId      string
	PlanId      string
}

func (f *SubscribePlanFlow) Run() SubscribePlanResponseBody {
	if err := f.validate(); err != nil {
		return SubscribePlanResponseBody{
			Error: err,
		}
	}

	return f.do()
}

func (f *SubscribePlanFlow) validate() error {
	return nil
}

func (f *SubscribePlanFlow) do() SubscribePlanResponseBody {

	return SubscribePlanResponseBody{
		Status: "success",
		Error:  nil,
	}
}
