package utils

import (
	"testing"
)

func TestImagePattern(t *testing.T) {
	// 测试用例
	tests := []struct {
		input    string
		expected bool
	}{
		// 正向测试：HTML img 标签
		{"<img src=\"https://example.com/image.jpg\">", true},
		{"<img src='http://example.com/image.png'/>", true},
		{"<img class=\"test\" src=\"https://example.com/image.gif\" alt=\"test\">", true},

		// 正向测试：直接图片 URL
		{"https://example.com/image.jpg", true},
		{"http://example.com/path/to/image.png", true},
		{"https://example.com/test.gif", true},
		{"https://example.com/test.webp", true},
		{"https://example.com/test.bmp", true},
		{"https://example.com/test.svg", true},

		// 负向测试：不包含图片
		{"This is a test string without any image.", false},
		{"<a href=\"https://example.com\">Link</a>", false},
		{"https://example.com/document.pdf", false},
		{"https://example.com/video.mp4", false},

		// 边界测试：部分匹配但不符合规则
		{"<img src=\"not_a_valid_url\">", false},
		{"https://example.com/image.txt", false},
	}

	// 执行测试
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := ContainsImage(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
