package ips

import (
	"errors"
	"fmt"
	"github.com/PandaXGO/PandaKit/httpclient"

	"net"
	"net/http"
	"strings"
)

const ipurl = "https://ip.cn/api/index"
const UNKNOWN = "XX XX"

// GetIP returns request client ip.
func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("Remote-Host")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forwarded-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil

		}

	}

	ip = r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}

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
