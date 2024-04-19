package ctrl

import (
	"mail-service/config"
	data "mail-service/models"
)

type SendMailBody struct {
	To      string `json:"to" validate:"required,email"`
	Subject string `json:"subject" validate:"required"`
	Body    string `json:"body" validate:"required"`
}

type SendMailResponse struct {
	Status string `json:"status"`
	Error  error  `json:"error"`
}

type SendMailFlow struct {
	Body SendMailBody
}

func (f *SendMailFlow) Run() SendMailResponse {
	if err := f.Validate(); err != nil {
		return SendMailResponse{Status: "failed", Error: err}
	}
	return f.do()
}

func (f *SendMailFlow) Validate() error {
	// TODO : Handle the send mail logic!
	return nil
}

func (f *SendMailFlow) do() SendMailResponse {
	// TODO : Handle the send mail logic!
	message := data.Message{
		To:      f.Body.To,
		Subject: f.Body.Subject,
		Data:    f.Body.Body,
	}
	config.ApplicationConfig.Wait.Add(1)
	config.ApplicationConfig.Mailer.MailerChan <- message

	return SendMailResponse{Status: "success", Error: nil}
}
