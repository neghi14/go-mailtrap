package mailtrap

import (
	"github.com/neghi14/go-mailtrap/internal/request"
)

type test struct {
	key       string
	accountID string
	inboxID   string
	projectID string
}

func NewTestService(key, id string) *test {
	return &test{
		key:       key,
		accountID: id,
	}
}

func (t *test) SetAccountID(id string) *test {
	return &test{
		accountID: id,
		key:       t.key,
		inboxID:   t.inboxID,
		projectID: t.projectID,
	}
}

func (t *test) SetInboxID(id string) *test {
	return &test{
		accountID: t.accountID,
		key:       t.key,
		inboxID:   id,
		projectID: t.projectID,
	}
}

func (t *test) SetProjectID(id string) *test {
	return &test{
		accountID: t.accountID,
		key:       t.key,
		inboxID:   t.inboxID,
		projectID: id,
	}
}

func (t *test) GetInbox() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) CreateInbox() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/projects/{project_id}/inboxes",
		Method: "POST",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "project_id", Value: t.projectID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) GetInboxAttributes() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) DeleteInbox() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}",
		Method: "DELETE",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) UpdateInbox() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}",
		Method: "PATCH",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) CleanInbox() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/clean",
		Method: "PATCH",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}
func (t *test) MarkInboxAsRead() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/all_read",
		Method: "PATCH",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) ResetInboxCredentials() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/reset_credentials",
		Method: "PATCH",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) EnableInboxEmail() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/toggle_email_username",
		Method: "PATCH",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}
func (t *test) TestResetEmail() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/reset_email_username",
		Method: "PATCH",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

// TODO: Add Queryparams
func (t *test) TestGetAttachments(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/attachments",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

// TODO: Add Queryparams
func (t *test) TestGetSingleAttachment(message_id, attachment_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/attachments/{attachment_id}",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id},
			{Key: "attachment_id", Value: attachment_id}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetSingleEmailMessage(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestUpdateEmailMessage(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}",
		Method: "PATCH",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestDeleteEmailMessage(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}",
		Method: "DELETE",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}
func (t *test) TestGetAllEmailMessage() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages",
		Method: "PATCH",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestForwardEmailMessage(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/forward",
		Method: "POST",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetEmailSpamScore(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/spam_report",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetEmailHTMLAnalysis(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/analyze",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}
	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetEmailText(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/body.txt",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetEmailRaw(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/body.raw",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetEmailSource(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/body.htmlsource",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetEmailHTML(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/body.html",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetEmailEML(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/body.eml",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetEmailHeaders(message_id string) interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/inboxes/{inbox_id}/messages/{message_id}/mail_headers",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "inbox_id", Value: t.inboxID},
			{Key: "message_id", Value: message_id}},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

// TODO: Add body Params
func (t *test) TestCreateProject() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/projects",
		Method: "POST",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
		},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetAllProject() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/projects",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
		},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestGetSingleProject() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/projects/{project_id}",
		Method: "GET",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "project_id", Value: t.projectID},
		},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

// TODO: add body params
func (t *test) TestUpdateProject() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/projects/{project_id}",
		Method: "PATCH",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "project_id", Value: t.projectID},
		},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

func (t *test) TestDeleteProject() interface{} {
	path := Path{
		Url:    uACCOUNTS_URL + "{account_id}/projects/{project_id}",
		Method: "DELETE",
		Params: []Params{
			{Key: "account_id", Value: t.accountID},
			{Key: "project_id", Value: t.projectID},
		},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequest(method, u, t.key)
	return res
}

// TODO: add body params
func (t *test) TestSendEmail(body *sendEmailBody) string {
	path := Path{
		Url:    uTEST_URL + "/{inbox_id}",
		Method: "POST",
		Params: []Params{
			{Key: "inbox_id", Value: t.inboxID},
		},
	}

	method, u := makePath(path)
	res, _ := request.MakeRequestWithBody(method, u, t.key, body)
	return res
}
