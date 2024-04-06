package mailtrap

type Mailtrap struct {
	key       string
	accountID string
}

func New(key string) *Mailtrap {
	return &Mailtrap{
		key: key,
	}
}

func (m *Mailtrap) SetAccountID(id string) *Mailtrap {
	return &Mailtrap{
		key:       m.key,
		accountID: id,
	}
}

func (m *Mailtrap) Accounts() *accounts {
	return NewAccounts(m.key)
}

func (m *Mailtrap) AccountAccess() *accountAccess {
	return NewAccountAccessService(m.key)
}

func (m *Mailtrap) Billing() *billing {
	return NewBillingService(m.key, m.accountID)
}

func (m *Mailtrap) Email() *email {
	return NewEmailService(m.key)
}

func (m *Mailtrap) TestEmail() *test {
	return NewTestService(m.key, m.accountID)
}

func (m *Mailtrap) SetBody() *sendEmailBody {
	return &sendEmailBody{}
}
