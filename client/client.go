package client

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ShadowObj/yescloudflare/cmd"
	"github.com/tidwall/gjson"
)

type Censys struct {
	Key      string
	Client   *http.Client
	NoRepeat bool
}

type Hit struct {
	IP   string
	Port string
}

func Exec(conf *cmd.Config, c *Censys) {
	var (
		f    *os.File
		err  error
		hits *[]*Hit
	)
	if f, err = os.Create(conf.Output); err != nil {
		conf.GetLogger().Fatalf("Open %s failed: %v", conf.Output, err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	for i := 0; i < 100; i++ {
		page := i + 1
		if !conf.Auto {
			inputT := ""
			conf.GetLogger().Printf("继续获取第 %d 页内容? (Y/N, Default Y): ", page)
			fmt.Scanln(&inputT)
			if inputT == "N" {
				break
			}
		}
		conf.GetLogger().Printf("正在获取第 %d 页内容...\n", page)
		if hits, err = c.get(page, conf.GetQuery()); err != nil {
			conf.GetLogger().Fatalf("Get failed: %v\n(page: %d, query: %s)", err, page, conf.GetQuery())
		}
		conf.GetLogger().Printf("在第 %d 页中发现了 %d 个节点.\n", page, len(*hits))
		for _, v := range *hits {
			if _, err = writer.WriteString(v.IP + ":" + v.Port + "\n"); err != nil {
				conf.GetLogger().Fatalf("Unable to write into buffer: %v", err)
			}
		}
	}
}

func (r *Censys) get(page int, query string) (*[]*Hit, error) {
	var (
		err     error
		hits    []*Hit
		req     *http.Request
		rawResp *http.Response
	)

	if req, err = r.newRequest(page, query); err != nil {
		return nil, err
	}
	if rawResp, err = r.Client.Do(req); err != nil {
		return nil, err
	}
	respData, _ := io.ReadAll(rawResp.Body)
	rawResp.Body.Close()
	respStr := string(respData)
	for _, v := range gjson.Get(respStr, "result.hits").Array() {
		ip := v.Get("ip").String()
		if r.NoRepeat {
			s := v.Get("services.#(service_name==\"HTTP\")")
			hits = append(hits, &Hit{
				IP:   ip,
				Port: s.Get("port").String(),
			})
		} else {
			for _, s := range v.Get("services.#(service_name==\"HTTP\")#").Array() {
				hits = append(hits, &Hit{
					IP:   ip,
					Port: s.Get("port").String(),
				})
			}
		}
	}
	return &hits, nil
}

func (r *Censys) newRequest(page int, query string) (*http.Request, error) {
	req, err := http.NewRequest("GET", "https://search.censys.io/api/v2/hosts/search?per_page=100&virtual_hosts=EXCLUDE&sort=RELEVANCE", strings.NewReader(""))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+r.Key)
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = q.Encode()
	return req, err
}
