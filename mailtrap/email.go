package mailtrap

import (
	"github.com/neghi14/go-mailtrap/internal/request"
)

type email struct {
	key string
}

func NewEmailService(key string) *email {
	return &email{
		key: key,
	}
}

// TODO: Add bodyparams
func (e *email) Send(body *sendEmailBody) interface{} {
	path := Path{
		Url:    uSEND_URL,
		Method: "POST",
	}
	method, u := makePath(path)
	res, _ := request.MakeRequestWithBody(method, u, e.key, body)
	return res
}

// TODO: Add bodyparams
func (e *email) SendBulk(body *[]sendEmailBody) interface{} {
	path := Path{
		Url:    uBULK_URL,
		Method: "POST",
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, e.key)
	return res
}
