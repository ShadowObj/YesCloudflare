package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ShadowObj/yescloudflare/client"
)

const (
	show_welcome = `# ShadowObj/yescloudflare
查询Cloudflare反代节点小工具
程序原型由 Joey Huang 开发
Telegram反馈群: https://t.me/+ft-zI76oovgwNmRh
`
	version = "1.0.0"
)

var (
	key    string
	port   int
	asn    int
	region string
	query  string
	file   string
	auto   bool
)

var (
	norepeat    bool
	invaildPort bool
)

func checkArguments() {
	flag.StringVar(&key, "key", "", "指定API密钥")
	flag.IntVar(&port, "port", 0, "指定端口 (默认全部)")
	flag.IntVar(&asn, "asn", 0, "指定ASN")
	flag.StringVar(&region, "region", "", "指定地区")
	flag.StringVar(&file, "file", "ip.txt", "指定输出文件 (默认ip.txt)")
	flag.BoolVar(&auto, "A", false, "自动获取下一页内容 (默认需要确认)")
	flag.BoolVar(&auto, "auto", false, "自动获取下一页内容 (默认需要确认)")
	flag.BoolVar(&norepeat, "norepeat", false, "自动去除重复IP (默认不去除)")
	flag.Parse()

	if key == "" || len(key) != 92 {
		log.Fatalf("Correct APIKEY is Required! (-key APIKEY)")
	}
	fmt.Printf("APIKEY: %s****%s\n", key[:2], key[len(key)-2:])
	if port <= 0 || port >= 65535 {
		fmt.Printf("端口 %d 无效, 不会有端口被过滤.\n", port)
		invaildPort = true
	}
	query = "NOT autonomous_system.asn=13335 and services.software.vendor='CloudFlare'"
	if asn != 0 {
		query += " and autonomous_system.asn=" + strconv.Itoa(asn)
	}
	if region != "" {
		query += " and location.country='" + region + "'"
	}
}

func execute() {
	var (
		f    *os.File
		c    *client.Censys
		err  error
		hits *[]*client.Hit
	)
	if f, err = os.Create(file); err != nil {
		log.Fatalf("Open %s failed: %v", file, err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	c = &client.Censys{Key: key, Client: &http.Client{}, NoRepeat: norepeat}
	for i := 0; i < 100; i++ {
		page := i + 1
		if !auto {
			inputT := ""
			fmt.Printf("继续获取第 %d 页内容? (Y/N, Default Y): ", page)
			fmt.Scanln(&inputT)
			if inputT == "N" {
				break
			}
		}
		fmt.Printf("正在获取第 %d 页内容...\n", page)
		if hits, err = c.Get(page, query); err != nil {
			log.Fatalf("Get failed: %v\n(page: %d, query: %s)", err, page, query)
		}
		fmt.Printf("在第 %d 页中发现了 %d 个节点.\n", page, len(*hits))
		for _, v := range *hits {
			if !invaildPort {
				if v.Port != port {
					continue
				}
			}
			if _, err = writer.WriteString(fmt.Sprintf("%s:%d\n", v.IP, v.Port)); err != nil {
				log.Fatalf("Unable to write into buffer: %v", err)
			}
		}
	}
}

func main() {
	fmt.Print(show_welcome)
	fmt.Print("版本 VERSION: ", version, "\n")
	checkArguments()
	execute()
}
