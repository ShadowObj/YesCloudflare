package cmd

import (
	"log"
	"os"
	"strings"
	//"github.com/ShadowObj/yescloudflare/log"
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

	logger *log.Logger
}

func (conf *Config) Check() {
	conf.logger = log.Default()
	if data, err := os.ReadFile(conf.Config); err == nil {
		tc := &tomlConfig{}
		if err := tc.Parse(string(data)); err != nil {
			conf.logger.Fatalf("配置文件 %s 格式不正确.", conf.Config)
		}
		tc.Pairing(conf)
		conf.logger.Printf("使用配置文件 %s 中的配置.", conf.Config)
	}
	if conf.Key == "" || len(conf.Key) != 92 {
		conf.logger.Fatalf("Correct APIKEY is Required! (-key APIKEY)")
	}
	conf.logger.Printf("APIKEY: %s****%s\n", conf.Key[:2], conf.Key[len(conf.Key)-2:])
	conf.query = "(NOT autonomous_system.asn=13335) and (services.software.vendor='CloudFlare')"
	if len(conf.Port) > 0 {
		conf.query += " and (services.port=" + strings.Join([]string(conf.Port), " or services.port=") + ")"
	} else {
		conf.logger.Printf("未指定端口，不会有端口被过滤.\n")
	}
	if len(conf.ASN) > 0 {
		conf.query += " and (autonomous_system.asn=" + strings.Join([]string(conf.ASN), " or autonomous_system.asn=") + ")"
	}
	if len(conf.Region) > 0 {
		conf.query += " and (location.country_code=" + strings.Join([]string(conf.ASN), " or location.country_code=") + ")"
	}
}

func (conf *Config) GetQuery() string {
	return conf.query
}

func (conf *Config) GetLogger() *log.Logger {
	return conf.logger
}
