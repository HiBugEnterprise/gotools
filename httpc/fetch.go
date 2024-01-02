package httpc

import (
	"github.com/HiBugEnterprise/gotools/errorx"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/jsonx"

	"github.com/avast/retry-go/v4"
)

func Stringify(v interface{}) string {
	b, err := jsonx.Marshal(v)

	if err != nil {
		return "{}"
	}
	return string(b)
}

type FetchConf struct {
	Method      string
	URL         string
	Data        any
	Headers     map[string]string
	AccessToken string
}

const maxAttempt = 3

func Fetch(c *FetchConf) (contents []byte, err error) {
	err = retry.Do(func() error {
		body := ""
		if c.Method == http.MethodPost {
			body = Stringify(c.Data)
		}

		request, err := http.NewRequest(c.Method, c.URL, strings.NewReader(body))
		if err != nil {
			return err
		}
		if c.Method == http.MethodPost {
			request.Header.Add("Content-Type", "application/json;charset=utf-8")
		}

		for k, v := range c.Headers {
			request.Header.Add(k, v)
		}
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			httpCode := http.StatusInternalServerError
			// 检查 resp 是否为 nil，确保不会在 defer 函数中发生 panic
			if resp != nil {
				httpCode = resp.StatusCode
			}
			return errorx.New("http-Client", httpCode, "发起HTTP请求异常").
				WithMetadata(errorx.Metadata{"req": c}).
				WithError(err)
		}

		defer func() {
			if err = resp.Body.Close(); err != nil {
				return
			}
		}()

		if contents, err = io.ReadAll(resp.Body); err != nil {
			return err
		}

		return nil
	},
		retry.Attempts(maxAttempt), // 重试3次
		retry.OnRetry(func(n uint, err error) {
			log.Printf("http请求发送失败，当前重试次数%d,error:%v\n", n, err.Error())
		}),
	)

	return
}

// FormFetch 表单提交
func FormFetch(c *FetchConf) (contents []byte, err error) {
	err = retry.Do(func() error {
		strData, ok := c.Data.(string)
		if !ok {
			return err
		}
		request, err := http.NewRequest(c.Method, c.URL, strings.NewReader(strData))
		if err != nil {
			return err
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		for k, v := range c.Headers {
			request.Header.Add(k, v)
		}

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			httpCode := http.StatusInternalServerError
			// 检查 resp 是否为 nil，确保不会在 defer 函数中发生 panic
			if resp != nil {
				httpCode = resp.StatusCode
			}
			return errorx.New("http-Client", httpCode, "发起HTTP请求异常").
				WithMetadata(errorx.Metadata{"req": c}).
				WithError(err)
		}

		defer func() {
			if err = resp.Body.Close(); err != nil {
				log.Printf("close HTTP响应异常，入参:%v", c)
				return

			}
		}()

		if contents, err = io.ReadAll(resp.Body); err != nil {
			return err
		}

		return nil
	},
		retry.Attempts(maxAttempt), // 重试3次
		retry.OnRetry(func(n uint, err error) {
			log.Printf("HTTP请求失败, 当前重试次数: %d，入参：%v\n", n, c)
		}),
	)

	return
}
