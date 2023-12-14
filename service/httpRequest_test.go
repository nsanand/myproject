package service

import (
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestGetSuccess(t *testing.T) {
	defer gock.Off()

	gock.New("http://foo.com").
		Reply(200).
		BodyString("foo foo")

	res, err := HttpAPIClient.Get("http://foo.com")

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)
	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, string(body), "foo foo")

	// Verify that we don't have pending mocks
	assert.Equal(t, gock.IsDone(), true)
}

func TestGetError(t *testing.T) {
	defer gock.Off()

	gock.New("http://foo.com").
		Reply(500)

	res, _ := HttpAPIClient.Get("http://foo.com")

	assert.Equal(t, res.StatusCode, 500)
	assert.Equal(t, gock.IsDone(), true)
}

func TestGetError1(t *testing.T) {
	_, err := HttpAPIClient.Get("123")

	assert.NotNil(t, err)
}
