package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

type PortList []string

func (l *PortList) String() string {
	return fmt.Sprintf("%v", *l)
}

func (l *PortList) Set(c string) error {
	for _, v := range strings.Split(c, ",") {
		p, err := strconv.Atoi(v)
		if err == nil && p > 0 && p < 65535 {
			*l = append(*l, v)
		}
	}
	return nil
}

type ASNList []string

func (p *ASNList) String() string {
	return fmt.Sprintf("%v", *p)
}

func (l *ASNList) Set(c string) error {
	for _, v := range strings.Split(c, ",") {
		p, err := strconv.Atoi(v)
		if err == nil && p > 0 {
			*l = append(*l, v)
		}
	}
	return nil
}

type RegionList []string

func (p *RegionList) String() string {
	return fmt.Sprintf("%v", *p)
}

func (l *RegionList) Set(c string) error {
	for _, v := range strings.Split(c, ",") {
		if len(v) == 2 {
			*l = append(*l, v)
		}
	}
	return nil
}
