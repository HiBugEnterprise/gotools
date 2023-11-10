package ips

import (
	"fmt"
	"github.com/PandaXGO/PandaKit/httpclient"
)

const ipurl = "https://ip.cn/api/index"
const UNKNOWN = "XX XX"

func GetRealAddressByIP(ip string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	url := fmt.Sprintf("%s?ip=%s&type=1", ipurl, ip)
	res := httpclient.NewRequest(url).Get()
	if res == nil || res.StatusCode != 200 {
		return UNKNOWN
	}
	toMap, err := res.BodyToMap()
	if err != nil {
		return UNKNOWN
	}
	return toMap["address"].(string)
}
