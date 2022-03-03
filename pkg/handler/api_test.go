package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testTable struct {
	name                 string
	inputBody            string
	expectedStatusCode   int
	expectedResponseBody string
	testErrorMessage     string
}

func TestHandler_getTitlesByUrls(t *testing.T) {
	router := new(Handler).InitRoutes()

	tt := []testTable{
		{
			name:                 "BadRequest",
			inputBody:            ``,
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"urls (array) parameter required"}`,
			testErrorMessage:     "Bad Request test failed",
		},
		{
			name:                 "BadUrls",
			inputBody:            `{"urls":"url"}`,
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"urls (array) parameter required"}`,
			testErrorMessage:     "Bad Urls test failed",
		},
		{
			name:                 "OK",
			inputBody:            `{"urls":["https://dota2.com","https://google.com","https://youtube.com","https://jetbrains.com","https://nesushestvuyushiysayt.com"]}`,
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"result":[{"status":200,"title":"Dota 2"},{"status":200,"title":"Google"},{"status":200,"title":"YouTube"},{"status":200,"title":"JetBrains: Essential tools for software developers and teams"},{"status":404,"title":""}]}`,
			testErrorMessage:     "OK status test failed",
		},
	}

	for _, testCase := range tt {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/titles", strings.NewReader(testCase.inputBody))
		router.ServeHTTP(w, req)

		if w.Code != testCase.expectedStatusCode {
			t.Errorf(testCase.testErrorMessage)
		}

		if testCase.name != "OK" {
			if w.Body.String() != testCase.expectedResponseBody {
				t.Errorf(testCase.testErrorMessage)
			}
			continue
		}

		rsp := Response{}
		expected := Response{}
		json.Unmarshal([]byte(testCase.expectedResponseBody), &expected)
		if err := json.Unmarshal([]byte(w.Body.String()), &rsp); err != nil {
			t.Errorf(testCase.testErrorMessage)
		}

		for i, exp := range expected.Result {
			if rsp.Result[i].Status != exp.Status || rsp.Result[i].Title != exp.Title {
				t.Errorf(testCase.testErrorMessage)
			}
		}
	}
}
