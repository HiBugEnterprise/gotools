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
	_, err := CreateChat(context.Background(), server, "", true, &data)
	assert.Equal(t, errors.New("url is empty").Error(), err.Error())

	data1 := ChatReq{Stream: false}
	_, err = CreateChat(context.Background(), server, "", false, &data1)
	assert.Equal(t, errors.New("url is empty").Error(), err.Error())
}

func TestCreateChat_HttpWriteNil(t *testing.T) {
	data := ChatReq{Stream: true}
	_, err := CreateChat(context.Background(), nil, "asdasd", true, &data)
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
	_, err := CreateChat(context.Background(), server, url, true, &data)
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
	_, err := CreateChat(context.Background(), server, url, false, &data)
	assert.NoError(t, err)
}

func TestCreateChat_teachingChatApi(t *testing.T) {
	url := "http://example.com/your-api-address"
	server := httptest.NewRecorder()
	msg := []*ImgChatMsg{
		{
			Role: "user",
			Content: []*Content{
				{
					Type: "text",
					Text: "",
				},
				{
					Type: "image_url",
					ImageURL: &ImageURL{
						URL: "https://picture.gptkong.com/20240601/21541ce7a68a654b579b7c5c9e450984d9.png",
					},
				},
			},
		},
	}
	data := TeachingChatCompletion{
		ID:       "12345",
		Messages: msg,
		Name:     "shizhe",
		Model:    "gpt4o",
		Stream:   true,
	}
	_, err := CreateChat(context.Background(), server, url, true, data)
	assert.NoError(t, err)
}
