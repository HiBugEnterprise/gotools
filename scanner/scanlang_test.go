package scanner

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanLanguage(t *testing.T) {
	language, err := ScanLanguage("E:\\GoWork\\src\\casbin-demo")
	assert.Nil(t, err)

	fmt.Printf("language: %+v\n", language)
}
