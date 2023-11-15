package ips

import (
	"fmt"
	"testing"
)

func TestGetRealAddressByIP(t *testing.T) {
	ip := "127.0.0.1"
	loc := GetRealAddressByIP(ip)
	fmt.Println(loc)
}
