package main

import (
	"fmt"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	connections := readJSONFiles("/home/echils/.config/qv2ray/connections")
	for a, b := range connections {
		domain := extractDomain(b)
		if domain != "" {
			ipv4 := resolveDomain(domain)
			if domain != ipv4 {
				b = strings.ReplaceAll(b, domain, ipv4)
				os.WriteFile(a, []byte(b), 0644)
			}
		}
	}
}

// 读取指定目录下的所有json文件
func readJSONFiles(direction string) map[string]string {

	jsonMap := make(map[string]string)

	filepath.Walk(direction, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("json files read failed%v\n", err)
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("json files read failed%v\n", err)
				return err
			}
			jsonMap[path] = string(data)
		}
		return nil
	})
	return jsonMap
}

// 替换域名为IPv4地址
func resolveDomain(domain string) string {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return domain
	}
	for _, ip := range ips {
		if ip.To4() != nil {
			return ip.String()
		}
	}
	return domain
}

// 提取JSON中的域名
func extractDomain(jsonStr string) string {
	re := regexp.MustCompile(`"address":\s*"([^"]+)"`)
	match := re.FindStringSubmatch(jsonStr)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}
