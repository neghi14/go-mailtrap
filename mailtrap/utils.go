package mailtrap

import (
	"fmt"
	"strings"
)

func makePath(path Path) (string, string) {
	var str string = path.Url
	var newUrl string
	if len(path.Params) > 0 {
		newUrl = str
		for _, val := range path.Params {
			str = strings.Replace(newUrl, "{"+val.Key+"}", val.Value, 1)
			newUrl = str
		}
	} else {
		newUrl = path.Url
	}

	return path.Method, newUrl
}

func makeQuery(query []Query) string {

	var str string
	for _, val := range query {
		str += fmt.Sprintf("%v=%v%v", val.Key, val.Value, "&")
	}
	str = "?" + str
	return str
}
