package openai

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestCreateChat_WithEmptyURL(t *testing.T) {
	server := httptest.NewRecorder()

	data := ChatReq{Stream: true}
	_, err := CreateChat(context.Background(), server, "", &data)
	assert.Equal(t, errors.New("url is empty").Error(), err.Error())

	data1 := ChatReq{Stream: false}
	_, err = CreateChat(context.Background(), server, "", &data1)
	assert.Equal(t, errors.New("url is empty").Error(), err.Error())
}

func TestCreateChat_HttpWriteNil(t *testing.T) {
	data := ChatReq{Stream: true}
	_, err := CreateChat(context.Background(), nil, "asdasd", &data)
	assert.Equal(t, errors.New("http response writer is nil").Error(), err.Error())
}

func TestCreateChatStreamNormal(t *testing.T) {
	url := "localhost:8080"
	server := httptest.NewRecorder()
	data := ChatReq{
		Id:    "abcd",
		Model: "asd",
		Messages: []*Message{
			{
				Role:    "user",
				Content: "*******",
			},
		},
		Stream: true,
	}
	_, err := CreateChat(context.Background(), server, url, &data)
	assert.NoError(t, err)
}

func TestCreateChatNoStreamNormal(t *testing.T) {
	url := "localhost:8080"
	server := httptest.NewRecorder()
	data := ChatReq{
		Id:    "abcd",
		Model: "asd",
		Messages: []*Message{
			{
				Role:    "user",
				Content: "*****",
			},
		},
		Stream: false,
	}
	_, err := CreateChat(context.Background(), server, url, &data)
	assert.NoError(t, err)
}
