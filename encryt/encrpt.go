package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/kakuilan/kgo"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

// PasswordHash 密码加密
func PasswordHash(password string) string {
	bytes, _ := kgo.KEncr.PasswordHash([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// PasswordVerify 密码验证
func PasswordVerify(password, hash string) bool {
	return kgo.KEncr.PasswordVerify([]byte(password), []byte(hash))
}
func HideStar(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			res2 := Substr2(str, 0, 3)
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^1([358][0-9]|4[579]|66|7[0135678]|9[89])[0-9]{8}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			result = Substr2(str, 0, 3) + "****" + Substr2(str, 7, 11)
		} else {
			nameRune := []rune(str)
			lens := len(nameRune)
			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
			} else if lens == 4 {
				result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
			} else if lens > 4 {
				result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
			}
		}
		return
	}
}
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}
func HashTimestamp(timestamp string) string {
	data := []byte(timestamp)
	md5Hash := md5.New()
	md5Hash.Write(data)
	md5Value := md5Hash.Sum(nil)
	return hex.EncodeToString(md5Value)
}
