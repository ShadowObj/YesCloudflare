package cmd

import (
	"strconv"
	"strings"
)

type PortList []string

func (l *PortList) String() string {
	return strings.Join(*l, ",")
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

func (l *ASNList) String() string {
	return strings.Join(*l, ",")
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

func (l *RegionList) String() string {
	return strings.Join(*l, ",")
}

func (l *RegionList) Set(c string) error {
	for _, v := range strings.Split(c, ",") {
		if len(v) == 2 {
			*l = append(*l, v)
		}
	}
	return nil
}

type PageRange struct {
	Start int
	End   int
}

func (r *PageRange) String() string {
	return strconv.Itoa(r.Start) + "-" + strconv.Itoa(r.End)
}

func (r *PageRange) Set(c string) error {
	strL := strings.Split(c, "-")
	if len(strL) > 1 {
		var err error
		if r.Start, err = strconv.Atoi(strL[0]); err != nil {
			r.Start, r.End = 0, 0
		}
		if r.End, err = strconv.Atoi(strL[1]); err != nil {
			r.Start, r.End = 0, 0
		}
		if r.Start > r.End || r.Start < 1 || r.End > 100 {
			r.Start, r.End = 0, 0
		}
	}
	return nil
}
