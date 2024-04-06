package mailtrap

// type permissions struct {
// 	client *Mailtrap
// }

// func (p *permissions) GetResources(access_id string) string {
// 	path := Path{
// 		Url:    uACCOUNTS_URL + "{account_id}/account_access/{account_access_id}/permissions/resources",
// 		Method: "GET",
// 		Params: []Params{{Key: "account_id", Value: p.client.AccountID}, {Key: "account_access_id", Value: access_id}},
// 	}
// 	method, u := makePath(path)
// 	res, _ := p.client.MakeRequest(method, u)
// 	return res

// }

// // TODO: add body params
// func (p *permissions) Manage(access_id string) string {
// 	path := Path{
// 		Url:    uACCOUNTS_URL + "{account_id}/account_access/{account_access_id}/permissions/bulk",
// 		Method: "PUT",
// 		Params: []Params{{Key: "account_id", Value: p.client.AccountID}, {Key: "account_access_id", Value: access_id}},
// 	}
// 	method, u := makePath(path)
// 	res, _ := p.client.MakeRequest(method, u)
// 	return res
// }
