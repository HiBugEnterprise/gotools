package chat

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	jsonx "github.com/HiBugEnterprise/gotools/jsonc"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

type Answer struct {
	Code     int    `json:"code"`
	Response string `json:"response"`
	Tokens   int64  `json:"tokens"`
}

func CreateStreamChat(ctx context.Context, w http.ResponseWriter, reqURL string, data any) (answer Answer, err error) {
	if reqURL == "" {
		return Answer{}, errors.New("url is empty")
	}

	if _, err = url.Parse(reqURL); err != nil {
		return
	}

	llmResp, err := SentHttpReqToModel(ctx, reqURL, data)
	if err != nil {
		return
	}
	defer llmResp.Body.Close()
	SetSSEHeader(w)

	return SentModelSSEResp(w, llmResp)
}

func SentHttpReqToModel(ctx context.Context, reqURL string, requestBody any) (resp *http.Response, err error) {
	reqBodyBytes, err := jsonx.Marshal(requestBody)
	if err != nil {
		return
	}
	reqBodyBuff := bytes.NewBuffer(reqBodyBytes)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, reqBodyBuff)
	if err != nil {
		return
	}

	httpReq.Header.Set("Content-Type", "application/json")

	cli := http.Client{}
	if resp, err = cli.Do(httpReq); err != nil {
		return
	}

	return
}

func SentModelSSEResp(w http.ResponseWriter, sseResp *http.Response) (answer Answer, err error) {
	reader := bufio.NewReader(sseResp.Body)
	headerData := []byte("data: ")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return Answer{}, errors.New("not support http flusher")
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
			return
		}

		flusher.Flush()
	}

	var answerDTO Answer
	if err = jsonx.Unmarshal(lastAns, &answerDTO); err != nil {
		return
	}

	if answerDTO.Code != http.StatusOK {
		return
	}

	return answerDTO, nil
}

func SetSSEHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 如果服务器和客户端之间有任何代理, 那将建议设置响应头 X-Accel-Buffering 为 no。
	w.Header().Set("X-Accel-Buffering", "no")
	// 在第一次渲染调用之前必须先行设置状态代码和响应头文件
	w.WriteHeader(http.StatusOK)
}
