package api_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yubing24/golang-spa/api"
)

func TestHelloWorldHandler(t *testing.T) {
	server := api.NewApiServer(api.Config{})

	req, _ := http.NewRequest(http.MethodGet, "/helloworld", nil)
	res := httptest.NewRecorder()

	server.ServeHTTP(res, req)
	assert.Equal(t, res.Result().StatusCode, http.StatusOK)

	output, _ := ioutil.ReadAll(res.Body)
	assert.Contains(t, strings.ToLower(string(output)), "hello, world", nil)
}
