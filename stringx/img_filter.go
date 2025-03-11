package utils

import "regexp"

func ContainsImage(question string) bool {
	// 匹配常见的图片 URL 或 HTML 图片标签
	// 更新正则表达式以排除无效的 URL
	imagePattern := `(?i)<img[^>]+src=["']https?:\/\/[^\s"']+?\.(jpg|jpeg|png|gif|webp|bmp|svg)["'][^>]*>|https?:\/\/[^\s]+\.(jpg|jpeg|png|gif|webp|bmp|svg)`
	re := regexp.MustCompile(imagePattern)
	return re.MatchString(question)
}
