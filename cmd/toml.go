package cmd

import (
	"sync"

	"github.com/BurntSushi/toml"
)

var once sync.Once

type tomlConfig struct {
	Config
	Port   string
	ASN    string
	Region string
	Page   string
}

func (tc *tomlConfig) Parse(data string) error {
	var err error
	once.Do(func() { _, err = toml.Decode(data, &tc) })
	if err != nil {
		return err
	}
	return nil
}

func (tc *tomlConfig) Pairing(conf *Config) error {
	conf.Key = tc.Key
	conf.Output = tc.Output
	conf.Auto = tc.Auto
	conf.Norepeat = tc.Norepeat
	conf.Port.Set(tc.Port)
	conf.ASN.Set(tc.ASN)
	conf.Region.Set(tc.Region)
	conf.Page.Set(tc.Page)
	return nil
}
