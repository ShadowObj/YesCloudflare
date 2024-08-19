package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ShadowObj/yescloudflare/client"
	"github.com/ShadowObj/yescloudflare/cmd"
)

const (
	show_welcome = `# ShadowObj/yescloudflare
查询Cloudflare反代节点小工具
程序原型由 Joey Huang 开发
Telegram反馈群: https://t.me/+ft-zI76oovgwNmRh
`
	version = "1.1.0"
	help    = `
Usage of YesCloudflare:
-c/-config string
	  指定配置文件 (默认config.toml)
	  注意: 指定配置文件不会覆盖命令行参数。
-o/-output ip.txt
	  指定输出文件 (默认ip.txt)
-A/-auto
	  自动获取下一页内容 (默认需要确认)
-key apikey
	  指定APIKEY
-norepeat
	  自动去除重复IP (默认不去除)
-port port
	  指定端口 (默认全部, 可用英文逗号分隔)
-asn asn1,asn2
	  指定ASN (默认全部, 可用英文逗号分隔)
-region CN,HK,JP,KR,TW
	  指定地区ISO3166二字码
	  (默认全部, 可用英文逗号分隔)
`
)

func parseConf(conf *cmd.Config) {
	flag.StringVar(&conf.Config, "c", "ip.txt", "")
	flag.StringVar(&conf.Config, "config", "config.toml", "")
	flag.StringVar(&conf.Output, "o", "ip.txt", "")
	flag.StringVar(&conf.Output, "output", "ip.txt", "")
	flag.BoolVar(&conf.Auto, "A", false, "")
	flag.BoolVar(&conf.Auto, "auto", false, "")
	flag.StringVar(&conf.Key, "key", "", "")
	flag.BoolVar(&conf.Norepeat, "norepeat", false, "")
	flag.Var(&conf.Port, "port", "")
	flag.Var(&conf.ASN, "asn", "")
	flag.Var(&conf.Region, "region", "")
	flag.Usage = func() { fmt.Print(help) }
	flag.Parse()
}

func main() {
	fmt.Print(show_welcome)
	fmt.Print("版本 VERSION: ", version, "\n")
	conf := &cmd.Config{}
	parseConf(conf)
	conf.Check()
	client.Exec(conf, &client.Censys{
		Key: conf.Key, Client: &http.Client{}, NoRepeat: conf.Norepeat,
	})
}
