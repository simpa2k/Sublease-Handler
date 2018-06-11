package utils

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"subLease/src/server"
	"subLease/src/server/database"
	"subLease/test/utils/mockDatabase"
	"testing"
)

func AssertRequestResponseMatchesOracle(t *testing.T, requestMethod string, endpoint string, body io.Reader, oracleGetter func(db database.Database) ([]byte, error)) database.Database {
	res, db := requestToServerWithMockDatabase(requestMethod, endpoint, body)
	jsonBytes, err := oracleGetter(db)
	assertNoErrorAndCorrectResponse(err, jsonBytes, res, t)

	return db
}

func AssertResponseMatchesOracle(t *testing.T, res *httptest.ResponseRecorder, oracleGetter func() ([]byte, error)) {
	jsonBytes, err := oracleGetter()
	assertNoErrorAndCorrectResponse(err, jsonBytes, res, t)
}

func assertNoErrorAndCorrectResponse(err error, jsonBytes []byte, res *httptest.ResponseRecorder, t *testing.T) {
	if err != nil {
		panic(err)
	}
	if expected, actual, equal := EqualJSON(string(jsonBytes), res.Body.String()); !equal {
		t.Error("Expected:\n\t"+expected+"\nbut got:\n\t", actual)
	}
}

func requestToServerWithMockDatabase(method string, endpoint string, body io.Reader) (*httptest.ResponseRecorder, database.Database) {
	r, db := SetupServerWithMockDatabase()
	return RequestToServer(r, method, endpoint, body), db
}

func RequestToServer(r http.Handler, method string, endpoint string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		panic(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	return res
}

func SetupServerWithMockDatabase() (http.Handler, database.Database) {
	db := mockDatabase.Create()
	s := server.Create(db)
	r := s.SetupRouter()

	return r, db
}

func EqualJSON(json1 string, json2 string) (string, string, bool) {
	replacedAndTrimmed1 := strings.Replace(strings.Trim(json1, ""), "\n", "", -1)
	replacedAndTrimmed2 := strings.Replace(strings.Trim(json2, ""), "\n", "", -1)

	return replacedAndTrimmed1, replacedAndTrimmed2, replacedAndTrimmed1 == replacedAndTrimmed2
}

func BuildQuery(baseUrl string, queries []struct {
	Key   string
	Value string
}) string {
	u, _ := url.Parse(baseUrl)
	q := u.Query()

	for _, query := range queries {
		q.Set(query.Key, query.Value)
	}

	u.RawQuery = q.Encode()
	completeUrl := u.String()

	return completeUrl
}
