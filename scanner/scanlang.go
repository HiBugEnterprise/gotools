package scanner

import (
	"encoding/json"
	"github.com/HiBugEnterprise/gotools/errorx"
	"github.com/HiBugEnterprise/gotools/jsonc"
	"os/exec"
)

type LangStatHeader struct {
	ClocUrl        string  `json:"cloc_url"`
	ClocVersion    string  `json:"cloc_version"`
	ElapsedSeconds float64 `json:"elapsed_seconds"`
	NFiles         uint32  `json:"n_files"`
	NLines         uint32  `json:"n_lines"`
	FilesPerSecond float64 `json:"files_per_second"`
	LinesPerSecond float64 `json:"lines_per_second"`
}

type LanguageStat struct {
	NFiles  uint32 `json:"nFiles"`
	Blank   uint32 `json:"blank"`
	Comment uint32 `json:"comment"`
	Code    uint32 `json:"code"`
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
		err = errorx.Internal(err, "扫描项目出错")
		return nil, err
	}

	resp := &ProjectLangStat{}
	if err = ParseProjectLangStat(output, resp); err != nil {
		err = errorx.Internal(err, "反序列化编程语言统计数据出错")
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
