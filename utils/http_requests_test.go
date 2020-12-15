package utils

import (
    "strconv"
    "testing"
    "time"
)

func TestHttpRequestHelper_GetUrl(t *testing.T) {

    h := HttpRequestHelper{}

    strUrl := "http://google.com.br"
    query := Query{
        "teste": "teste",
    }

    reqUrl := h.GetUrl(strUrl, query)
    if reqUrl.Host != "google.com.br" && reqUrl.Scheme == "http" && reqUrl.RawQuery == "teste=teste" {
        t.Error("url need to be http://google.com.br?teste=teste")
    }

    strUrl = "http://google.com.br"
    reqUrl = h.GetUrl(strUrl, nil)
    if reqUrl.Host != "google.com.br" && reqUrl.Scheme == "http" {
        t.Error("url need to be http://google.com.br")
    }

}

func TestHttpRequestHelper_GetNetClient(t *testing.T) {

    h := HttpRequestHelper{}

    net := h.GetNetClient(time.Second * 30)
    if net.Timeout != (time.Second * 30) {
        t.Error("timeout need to be time.Second * 30")
    }

    net = h.GetNetClient(0)
    if net.Timeout != (time.Second * DefaultTimeOut) {
        t.Error("timeout need to be time.Second * " + strconv.FormatInt(DefaultTimeOut, 10))
    }
}
