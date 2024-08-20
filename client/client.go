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
	Key    string
	Client *http.Client
	Conf   *cmd.Config
}

type Hit struct {
	IP   string
	Port string
}

// Exec performs a task with censys client
func Exec(c *Censys) {
	var (
		f    *os.File
		err  error
		hits *[]*Hit
	)
	conf := c.Conf
	logger := conf.GetLogger()
	if f, err = os.Create(conf.Output); err != nil {
		logger.Fatalf("Open %s failed: %v", conf.Output, err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	for i := c.Conf.Page.Start; i < (c.Conf.Page.End + 1); i++ {
		if !conf.Auto {
			inputT := ""
			logger.Printf("继续获取第 %d 页内容? (Y/N, Default Y): ", i)
			fmt.Scanln(&inputT)
			if inputT == "N" {
				break
			}
		}
		logger.Printf("正在获取第 %d 页内容...\n", i)
		if hits, err = c.get(i, conf.GetQuery()); err != nil {
			logger.Fatalf("Get failed: %v\n(page: %d, query: %s)", err, i, conf.GetQuery())
		}
		logger.Printf("在第 %d 页中发现了 %d 个节点.\n", i, len(*hits))
		for _, v := range *hits {
			ip := v.IP
			if strings.Contains(v.IP, ":") {
				ip = "[" + v.IP + "]"
			}
			if _, err = writer.WriteString(ip + ":" + v.Port + "\n"); err != nil {
				logger.Fatalf("Unable to write into buffer: %v", err)
			}
		}
	}
}

func (c *Censys) get(page int, query string) (*[]*Hit, error) {
	var (
		err     error
		hits    []*Hit
		req     *http.Request
		rawResp *http.Response
	)

	if req, err = c.newRequest(page, query); err != nil {
		return nil, err
	}
	if rawResp, err = c.Client.Do(req); err != nil {
		return nil, err
	}
	respData, _ := io.ReadAll(rawResp.Body)
	rawResp.Body.Close()
	respStr := string(respData)
	for _, v := range gjson.Get(respStr, "result.hits").Array() {
		ip := v.Get("ip").String()
		if c.Conf.Norepeat {
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

func (c *Censys) newRequest(page int, query string) (*http.Request, error) {
	req, err := http.NewRequest("GET", "https://search.censys.io/api/v2/hosts/search?per_page=100&virtual_hosts=EXCLUDE&sort=RELEVANCE", strings.NewReader(""))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+c.Key)
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = q.Encode()
	return req, err
}
