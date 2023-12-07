package jsonc

import (
	"errors"

	"github.com/bytedance/sonic"
)

func Marshal(val any) ([]byte, error) {
	return sonic.Marshal(val)
}

func Unmarshal(data []byte, v any) error {
	if len(data) == 0 {
		return errors.New("jsonx: empty data to unmarshal")
	}

	return sonic.Unmarshal(data, v)
}
