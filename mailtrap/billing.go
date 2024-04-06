package mailtrap

import "github.com/neghi14/go-mailtrap/internal/request"

type billing struct {
	key       string
	accountID string
}

func NewBillingService(key, id string) *billing {
	return &billing{
		key:       key,
		accountID: id,
	}
}

func (b *billing) SetAccountID(id string) *billing {

	return &billing{
		key:       b.key,
		accountID: id,
	}
}

func (b *billing) Get() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/billing/usage",
		Method: "GET",
		Params: []Params{{Key: "account_id", Value: b.accountID}},
	}

	method, u := makePath(path)

	res, _ := request.MakeRequest(method, u, b.key)

	return res
}
