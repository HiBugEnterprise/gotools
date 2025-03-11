package utils

import (
	"regexp"
	"strings"
)

// EscapeJSONBackslashes 处理 JSON 字符串中的反斜杠转义问题
// 特别针对包含 LaTeX 公式的 JSON 数据
func EscapeJSONBackslashes(content string) string {
	// 处理常见的 LaTeX 命令和符号
	latexCommands := []string{
		// 括号和分隔符
		`\(`, `\)`, `\{`, `\}`, `\[`, `\]`, `\|`,

		// 数学运算符
		`\times`, `\div`, `\pm`, `\mp`, `\cdot`,

		// 关系运算符
		`\neq`, `\approx`, `\equiv`, `\leq`, `\geq`, `\ll`, `\gg`,

		// 集合符号
		`\cup`, `\cap`, `\subset`, `\supset`, `\in`, `\notin`,

		// 上下标和根号
		`\sqrt`, `\frac`, `\sum`, `\prod`, `\int`,

		// 希腊字母
		`\alpha`, `\beta`, `\gamma`, `\delta`, `\epsilon`, `\zeta`, `\eta`,
		`\theta`, `\iota`, `\kappa`, `\lambda`, `\mu`, `\nu`, `\xi`,
		`\pi`, `\rho`, `\sigma`, `\tau`, `\upsilon`, `\phi`, `\chi`, `\psi`, `\omega`,
		`\Gamma`, `\Delta`, `\Theta`, `\Lambda`, `\Xi`, `\Pi`, `\Sigma`, `\Phi`, `\Psi`, `\Omega`,

		// 其他常用符号
		`\infty`, `\partial`, `\nabla`, `\forall`, `\exists`, `\rightarrow`, `\leftarrow`,
		`\Rightarrow`, `\Leftarrow`, `\leftrightarrow`, `\Leftrightarrow`,
	}

	// 转义所有 LaTeX 命令
	for _, cmd := range latexCommands {
		content = strings.ReplaceAll(content, cmd, strings.ReplaceAll(cmd, `\`, `\\`))
	}

	// 处理其他可能的转义序列
	// 1. 处理未被上面列表覆盖的 LaTeX 命令（以反斜杠开头的单词）
	re := regexp.MustCompile(`\\([a-zA-Z]+)`)
	content = re.ReplaceAllString(content, `\\$1`)

	// 2. 处理数学上标、下标等
	content = strings.ReplaceAll(content, `\^`, `\\^`)
	content = strings.ReplaceAll(content, `\_`, `\\_`)

	// 3. 处理 JSON 中已有的合法转义序列，避免重复转义
	// 先将合法的转义序列临时替换为特殊标记
	tempMarkers := map[string]string{
		`\\`: `__DOUBLE_BACKSLASH__`,
		`\"`: `__ESCAPED_QUOTE__`,
		`\/`: `__ESCAPED_SLASH__`,
		`\b`: `__ESCAPED_BACKSPACE__`,
		`\f`: `__ESCAPED_FORMFEED__`,
		`\n`: `__ESCAPED_NEWLINE__`,
		`\r`: `__ESCAPED_RETURN__`,
		`\t`: `__ESCAPED_TAB__`,
		`\u`: `__ESCAPED_UNICODE__`,
	}

	for escape, marker := range tempMarkers {
		content = strings.ReplaceAll(content, escape, marker)
	}

	// 4. 处理剩余的单个反斜杠（可能是未处理的特殊字符）
	content = strings.ReplaceAll(content, `\`, `\\`)

	// 5. 恢复之前标记的合法转义序列
	for escape, marker := range tempMarkers {
		content = strings.ReplaceAll(content, marker, escape)
	}

	// 6. 修复 Unicode 转义序列
	// 将 \\uXXXX 改回 \uXXXX
	unicodeRe := regexp.MustCompile(`\\\\u([0-9a-fA-F]{4})`)
	content = unicodeRe.ReplaceAllString(content, `\u$1`)

	return content
}
