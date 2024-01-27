package chat

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStreamChat_WithEmptyURL(t *testing.T) {
	server := httptest.NewRecorder()

	_, err := CreateStreamChat(context.Background(), server, "", nil)
	assert.Equal(t, errors.New("url is empty"), err)
}

func TestCreateStreamChat_WithInvalidURL(t *testing.T) {
	server := httptest.NewRecorder()

	_, err := CreateStreamChat(context.Background(), server, "ht2sd:123", nil)
	assert.NotNil(t, err)
}

func TestCreateStreamChat_Normal(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := "http://127.0.0.1:8080/api/chat"
		answer, err := CreateStreamChat(context.Background(), w, url, nil)

		assert.NotEmpty(t, answer)
		assert.Nil(t, err)

		w.Write([]byte("data: answer: " + answer.Response + "\n"))
		t.Log("answer: ", answer)
	}))
	defer mockServer.Close()

}

func TestCreateStreamChat_WithError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("data: server error\n"))
	}))
	defer mockServer.Close()

	server := httptest.NewRecorder()
	answer, err := CreateStreamChat(context.Background(), server, mockServer.URL, nil)
	assert.Empty(t, answer)
	assert.NotNil(t, err)
}
