package scanner

import (
	"encoding/json"
	"github.com/HiBugEnterprise/gotools/jsonc"
	"os/exec"
)

type LangStatHeader struct {
	ClocUrl        string  `json:"cloc_url"`         // cloc的GitHub地址
	ClocVersion    string  `json:"cloc_version"`     //版本
	ElapsedSeconds float64 `json:"elapsed_seconds"`  //扫描所用的时间（秒）
	NFiles         uint32  `json:"n_files"`          //扫描的文件总数
	NLines         uint32  `json:"n_lines"`          //扫描的行数总数
	FilesPerSecond float64 `json:"files_per_second"` //每秒扫描文件的数目
	LinesPerSecond float64 `json:"lines_per_second"` //每秒扫描行的数目
}

type LanguageStat struct {
	NFiles  uint32 `json:"nFiles"`  // 使用该编程语言编写的文件数
	Blank   uint32 `json:"blank"`   // 该编程语言的空行数
	Comment uint32 `json:"comment"` // 该编程语言的注释行数
	Code    uint32 `json:"code"`    // 该编程语言的代码行数
}

type ProjectLangStat struct {
	Header LangStatHeader          `json:"header"`
	Langs  map[string]LanguageStat `json:"langs"`
}

func ScanLanguage(filePath string) (*ProjectLangStat, error) {
	cmd := exec.Command("cloc", "--exclude-dir=tmp", "--exclude-ext=.tpl,.md", "--json", ".")
	cmd.Dir = filePath
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	resp := &ProjectLangStat{}
	if err = ParseProjectLangStat(output, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ParseProjectLangStat 反序列化json到ProjectLangStat
func ParseProjectLangStat(jsonBytes []byte, projectLangStat *ProjectLangStat) error {
	//反序列化json
	var raw map[string]json.RawMessage
	err := jsonc.Unmarshal(jsonBytes, &raw)
	if err != nil {
		return err
	}
	//反序列化json
	for key, langStat := range raw {
		if key == "header" {
			err = json.Unmarshal(langStat, &projectLangStat.Header)
			if err != nil {
				return err
			}
		} else {
			if projectLangStat.Langs == nil {
				projectLangStat.Langs = make(map[string]LanguageStat)
			}
			var lang LanguageStat
			if err = json.Unmarshal(langStat, &lang); err != nil {
				return err
			}
			projectLangStat.Langs[key] = lang
		}
	}
	return nil
}
