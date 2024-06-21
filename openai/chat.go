package openai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/HiBugEnterprise/gotools/httpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func CreateChat(ctx context.Context, w http.ResponseWriter, reqURL string, data *ChatReq) (answer *LLMAnswer, err error) {
	if data.Stream {
		return StreamChat(ctx, w, reqURL, data)
	}
	return NoStreamChat(ctx, w, reqURL, data)
}

func NoStreamChat(ctx context.Context, w http.ResponseWriter, reqURL string, data *ChatReq) (answer *LLMAnswer, err error) {
	if reqURL == "" {
		return nil, errors.New("url is empty")
	}

	var resp *LLMAnswer
	if _, err = httpc.Post(reqURL,
		httpc.SetBody(data),
		httpc.SetResult(&resp)); err != nil {
		err = errors.Wrap(err, "request llm chat failed")
		return
	}

	if resp == nil {
		err = errors.New("llm return empty answer")
		return nil, err
	}

	if resp.Code != 200 {
		err = errors.Wrap(errors.New(""), "llm return error code")
		return nil, err
	}

	return resp, nil
}

func StreamChat(ctx context.Context, w http.ResponseWriter, reqURL string, data *ChatReq) (answer *LLMAnswer, err error) {
	if reqURL == "" {
		return nil, errors.New("url is empty")
	}

	if _, err = url.Parse(reqURL); err != nil {
		return
	}

	llmResp, err := SentHttpReqToModel(ctx, reqURL, data)
	if err != nil {
		return
	}
	defer llmResp.Body.Close()

	if w == nil {
		return nil, errors.New("http response writer is nil")
	}
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

func SentModelSSEResp(w http.ResponseWriter, sseResp *http.Response) (answer *LLMAnswer, err error) {
	reader := bufio.NewReader(sseResp.Body)
	headerData := []byte("data: ")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return nil, errors.New("not support http flusher")
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
	var answerDTO LLMAnswer
	if err = json.Unmarshal(lastAns, &answerDTO); err != nil {
		return
	}

	if answerDTO.Code != http.StatusOK {
		return
	}

	return &answerDTO, nil
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
