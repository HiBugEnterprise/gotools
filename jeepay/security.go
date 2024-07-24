package jeepay

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func checkSign(data any, key string, securityType string, respSign string) (bool, error) {
	if len(respSign) == 0 {
		return true, nil
	}
	dataMap := Struct2MapName(data)

	sign, err := encrypt(dataMap, key, securityType)
	if err != nil {
		return false, err
	}
	if respSign != sign {
		return false, nil
	}

	return true, nil
}

func encrypt(params map[string]interface{}, secret string, securityType string) (string, error) {
	// key 按照 ASCII码从小到大排序
	keys := make([]string, 0)
	for key, _ := range params {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var buffer bytes.Buffer
	for _, item := range keys {
		value := params[item]

		valRef := reflect.ValueOf(value)
		if !valRef.IsValid() || valRef.IsZero() {
			continue
		}
		buffer.WriteString(item)
		buffer.WriteString("=")
		buffer.WriteString(Strval(value))
		buffer.WriteString("&")
	}

	buffer.WriteString("key=")
	buffer.WriteString(secret)
	if securityType == "MD5" {
		return md5encrypt(buffer.Bytes()), nil
	} else if securityType == "RSA" {
		return "", nil
	} else {
		return md5encrypt(buffer.Bytes()), nil
	}
}

func md5encrypt(buffer []byte) string {
	sum := md5.Sum(buffer)
	sec := fmt.Sprintf("%x", sum)
	return strings.ToUpper(sec)
}
