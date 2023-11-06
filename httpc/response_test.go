package httpc

import (
	"github.com/HiBugEnterprise/gotools/errorx"
	"testing"
)

func TestReponse(t *testing.T) {
	err := errorx.New("test", int(errorx.CodeInternalErr), "test")

	RespError(nil, nil, err)
}
