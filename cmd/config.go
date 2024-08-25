package cmd

import (
	"log"
	"os"
	"strings"
	//"github.com/ShadowObj/yescloudflare/log"
)

const (
	defaultKeyLen    = 92
	defaultPageStart = 1
	defaultPageEnd   = 10
	defaultOutput    = "ip.txt"
)

type Config struct {
	Key      string
	Port     PortList
	ASN      ASNList
	Region   RegionList
	query    string
	Output   string
	Config   string
	Auto     bool
	Norepeat bool
	Page     PageRange

	logger *log.Logger
}

// Check checks whether the config is vaild and generate the query
func (c *Config) Check() {
	c.logger = log.Default()
	if data, err := os.ReadFile(c.Config); err == nil {
		tc := &tomlConfig{}
		if err := tc.Parse(string(data)); err != nil {
			c.logger.Fatalf("配置文件 %s 格式不正确.", c.Config)
		}
		if err = tc.Pairing(c); err != nil {
			c.logger.Fatalf("配置文件 %s 格式不正确: %s", c.Config, err)
		}
		c.logger.Printf("使用配置文件 %s 中的配置.", c.Config)
	}
	if c.Key == "" || len(c.Key) != defaultKeyLen {
		c.logger.Fatalf("Correct APIKEY is Required! (-key APIKEY)")
	}
	c.logger.Printf("APIKEY: %s****%s\n", c.Key[:2], c.Key[len(c.Key)-2:])
	c.query = "NOT autonomous_system.asn={13335,209242} and services.software.vendor='CloudFlare'"
	if len(c.Port.intS) > 0 {
		c.query += " and services.port={" + strings.Join(c.Port.strS, ",") + "}"
	} else {
		c.logger.Printf("未指定端口，不会有端口被过滤.\n")
	}
	if len(c.ASN) > 0 {
		c.query += " and autonomous_system.asn={" + strings.Join([]string(c.ASN), ",") + "}"
	}
	if len(c.Region) > 0 {
		c.query += " and location.country_code={" + strings.Join([]string(c.Region), ",") + "}"
	}
	if c.Page.Start == 0 || c.Page.End == 0 {
		c.Page.Start, c.Page.End = defaultPageStart, defaultPageEnd
	}
	if c.Output == "" {
		c.Output = defaultOutput
	}
}

func (c *Config) GetQuery() string {
	return c.query
}

func (c *Config) GetLogger() *log.Logger {
	return c.logger
}
