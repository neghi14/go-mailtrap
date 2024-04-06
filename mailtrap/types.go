package mailtrap

const (
	uBASE_URL     = "https://mailtrap.io/api/"
	uSEND_URL     = "https://send.api.mailtrap.io/api/send"
	uBULK_URL     = "https://bulk.api.mailtrap.io/api/send"
	uTEST_URL     = "https://sandbox.api.mailtrap.io/api/send"
	uACCOUNTS_URL = uBASE_URL + "accounts/"
)

type Params struct {
	Key   string
	Value string
}

type Path struct {
	Url    string
	Method string
	Params []Params
}

type Query Params
