package mailtrap

import "github.com/neghi14/go-mailtrap/internal/request"

type accounts struct {
	key string
}

func NewAccounts(key string) *accounts {
	return &accounts{
		key: key,
	}
}

//Get All implements https://mailtrap.io/api/accounts
//Get a list of your Mailtrap accounts.
func (a *accounts) GetAll() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL,
		Method: "GET",
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, a.key)
	return res
}
