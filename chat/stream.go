package chat

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/HiBugEnterprise/gotools/errorx"
	"github.com/HiBugEnterprise/gotools/httpc"
	jsonx "github.com/HiBugEnterprise/gotools/jsonc"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

type AnswerDTO struct {
	Code     int    `json:"code"`
	Response string `json:"response"`
}

func CreateStreamChat(w http.ResponseWriter, reqURL string, data any) (answer string, err error) {
	if reqURL == "" {
		return "", errors.New("url is empty")
	}
	if _, err = url.Parse(reqURL); err != nil {
		return "", errorx.Internal(err, "%s is an invalid URL", reqURL)
	}

	res, err := httpc.Post(reqURL, httpc.SetBody(data))
	if err != nil {
		return
	}
	setSSEHeader(w)
	reader := bufio.NewReader(res.RawBody())

	headerData := []byte("data: ")
	flusher, ok := w.(http.Flusher)
	if !ok {

		return "", errors.New("not support http flusher")
	}
	var lastAns []byte

	for {
		var line []byte
		if line, err = reader.ReadBytes('\n'); err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
				break
			}
			return
		}

		line = bytes.TrimSpace(line)
		line = bytes.TrimPrefix(line, headerData)

		strData := string(line)
		if strData == "" || strData == "\n" || strData == "\r" || strData == "data:" {
			continue
		}
		lastAns = line

		regexStr := "^: ping - \\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\.\\d{6}$"
		regex := regexp.MustCompile(regexStr)
		if isMatch := regex.MatchString(strData); isMatch {
			continue
		}
		_, err = fmt.Fprintf(w, "data: %v\n\n", strData)
		if err != nil {

			return "", errorx.Internal(err, "model response sent to http failed")
		}

		flusher.Flush()
	}

	var answerDTO AnswerDTO
	if err = jsonx.Unmarshal(lastAns, &answerDTO); err != nil {
		return
	}

	if answerDTO.Code != http.StatusOK {
		return "", errorx.Internal(err, "model response failed")
	}

	answer = answerDTO.Response

	return
}

func setSSEHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 如果服务器和客户端之间有任何代理, 那将建议设置响应头 X-Accel-Buffering 为 no。
	w.Header().Set("X-Accel-Buffering", "no")
	// 在第一次渲染调用之前必须先行设置状态代码和响应头文件
	w.WriteHeader(http.StatusOK)
}
