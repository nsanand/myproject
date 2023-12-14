package assignment

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"myproject/service"
	"net/http"
	"strings"
	"testing"
)

var httpClient service.MockHttpClient = service.MockHttpClient{}

func TestMain(m *testing.M) {
	m.Run()
}

func TestFetchDataSuccess(t *testing.T) {
	respData := `{"userId":1,"id":1,"title":"Title text","completed": false}`

	httpClient.On("Get", "https://jsonplaceholder.typicode.com/todos/1").Return(
		&http.Response{Body: io.NopCloser(strings.NewReader(string(respData)))},
		nil)

	err := FetchData(&httpClient, 1)
	assert.Nil(t, err)
}

func TestFetchDataError(t *testing.T) {
	err := errors.New("error: data not found")
	httpClient.On("Get", "https://jsonplaceholder.typicode.com/todos/2").Return(
		&http.Response{Body: nil, StatusCode: 404},
		err)

	e := FetchData(&httpClient, 2)

	assert.NotNil(t, e)
}
