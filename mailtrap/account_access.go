package mailtrap

import (
	"fmt"

	"github.com/neghi14/go-mailtrap/internal/request"
)

type accountAccess struct {
	key       string
	accountID string
}

func NewAccountAccessService(key string) *accountAccess {
	return &accountAccess{
		key: key,
	}
}

func (a *accountAccess) SetAccountID(id string) *accountAccess {
	return &accountAccess{
		accountID: id,
		key:       a.key,
	}
}

// ListAndInvite implements the https://mailtrap.io/api/accounts/{account_id}/account_accesses endpoint
// it gets the list of accounts for which the specifier_type is User,
// you need admin/owner access for it to work
// id implies to Account ID
func (a *accountAccess) ListAndInviteAccess(query []Query) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/account_accesses",
		Method: "GET",
		Params: []Params{{Key: "account_id", Value: a.accountID}},
	}
	method, u := makePath(path)
	fmt.Println(path)
	q := makeQuery(query)
	res, _ := request.MakeRequest(method, u+q, a.key)
	return res
}

// Remove Account Access implements https://mailtrap.io/api/accounts/{account_id}/account_accesses/{account_access_id}
// endpoint.
// If specifier type is User, it removes user permissions.
// If specifier type is Invite or ApiToken, it removes specifier along with permissions.
// You have to be an account admin/owner for this endpoint to work.
// id implies to Account ID
func (a *accountAccess) RemoveAccountAccess(access_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/account_access/{account_access_id}",
		Method: "DELETE",
		Params: []Params{{Key: "account_id", Value: a.accountID}, {Key: "account_access_id", Value: access_id}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, a.key)
	return res
}
