package api

import (
	"encoding/base64"
	"log"
	"net/http"
	"testing"
)

func TestCheckAuth(t *testing.T) {
	type test struct {
		username  string
		password  string
		response  string
		addHeader bool
	}

	testSuite := []test{
		test{
			username:  "admin",
			password:  "admin",
			response:  "",
			addHeader: true,
		},
		test{
			username:  "admin",
			password:  "",
			response:  "invalid password",
			addHeader: true,
		},
		test{
			username:  "",
			password:  "admin",
			response:  "user not found",
			addHeader: true,
		},
		test{
			response:  "unauthorized",
			addHeader: false,
		},
	}

	for _, mytest := range testSuite {
		r, err := http.NewRequest("", "", nil)
		if err != nil {
			log.Fatal(err)
		}

		if mytest.addHeader {
			header := base64.StdEncoding.EncodeToString([]byte(mytest.username + ":" + mytest.password))
			r.Header.Add("Authorization", "Basic "+header)
		}

		response := ""
		if err = CheckAuth(r); err != nil {
			response = err.Error()
		}

		if response != mytest.response {
			t.Error(
				"\nexpected response", mytest.response,
				"\nfound response", response,
				"\nfor test", mytest,
			)
		}
	}
}
