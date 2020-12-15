package utils

import (
    "encoding/base64"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "time"
)

const (
    DefaultTimeOut = 30
)

type HttpRequestHelper struct {
    BasicAuthUser BasicAuthUser
    Res           *http.Response
}

type Header map[string][]string
type Query map[string]string
type BasicAuthUser struct {
    Name string
    Pass string
}

func (h *HttpRequestHelper) Do(req *http.Request, netClient *http.Client) ([]byte, error) {
    res, err := netClient.Do(req)
    h.Res = res
    if err != nil {
        return nil, err
    }
    bodyData, err := ioutil.ReadAll(res.Body)
    defer res.Body.Close()
    if err != nil {
        return nil, err
    }
    return bodyData, nil
}

func (h *HttpRequestHelper) GetUrl(u string, query Query) *url.URL {
    if query == nil {
        reqUrl, _ := url.Parse(u)
        return reqUrl
    }
    if len(query) > 0 {
        q := make(url.Values)
        keys := make([]string, 0, len(query))
        for k := range query {
            keys = append(keys, k)
        }
        for _, k := range keys {
            vs := query[k]
            q.Add(k, vs)
        }
        reqUrl, _ := url.Parse(u + "?" + q.Encode())
        return reqUrl
    } else {
        reqUrl, _ := url.Parse(u)
        return reqUrl
    }
}

func (h *HttpRequestHelper) GetNetClient(timeout time.Duration) *http.Client {
    if timeout != 0 {
        return &http.Client{
            Timeout: timeout,
        }
    } else {
        return &http.Client{
            Timeout: time.Second * DefaultTimeOut,
        }
    }
}

/**
Make a complex request body
*/
func (h *HttpRequestHelper) GetRequestBodyApi(requestType string, requetUrl *url.URL, dataBody []byte, header map[string][]string) *http.Request {
    var he map[string][]string
    if header != nil {
        he = header
    } else {
        he = map[string][]string{
            "authorization": {"Basic " + h.GetBasicAuthString()},
            "Content-type":  {"application/json"},
        }
    }
    return &http.Request{
        Method: requestType,
        URL:    requetUrl,
        Body:   ioutil.NopCloser(strings.NewReader(string(dataBody))),
        Header: he,
    }
}

func (h *HttpRequestHelper) GetBasicAuthString() string {
    return base64.StdEncoding.EncodeToString([]byte(h.BasicAuthUser.Name + ":" + h.BasicAuthUser.Pass))
}
