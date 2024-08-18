package client

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type Censys struct {
	Key      string
	Client   *http.Client
	NoRepeat bool
}

type Hit struct {
	IP   string
	Port int
}

func (r *Censys) Get(page int, query string) (*[]*Hit, error) {
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
				Port: int(s.Get("port").Int()),
			})
		} else {
			for _, s := range v.Get("services.#(service_name==\"HTTP\")#").Array() {
				hits = append(hits, &Hit{
					IP:   ip,
					Port: int(s.Get("port").Int()),
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
