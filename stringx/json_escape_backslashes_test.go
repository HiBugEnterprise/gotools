package utils

import (
	_ "embed"
	"testing"
)

//go:embed json_escape_backslashes_test.json
var jsonConvertTest []byte

func TestEscapeJSONBackslashes(t *testing.T) {
	res := EscapeJSONBackslashes(string(jsonConvertTest))
	t.Log(res)
}
