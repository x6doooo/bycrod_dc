package yahoo

import (
    "net/http"
    "time"
    "io/ioutil"
    "encoding/json"
    "io"
    "compress/gzip"
    "errors"
    //"bycrod_center/model"
    "strings"
    "bycrod_dc/service/util"
    "github.com/corpix/uarand"
)

var (
    hosts = []string{
        "query1.finance.yahoo.com",
        "query2.finance.yahoo.com",
    }
    client *http.Client

    //uaIdx = 0
/*    uas = []string{
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:50.0) Gecko/20100101 Firefox/50.0",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.98 " +
            "Safari/537.36",
    }*/
)

func init() {
    client = &http.Client{
        Timeout: time.Second * 5,
    }
}

func newRequest(host, code, interval, rangeStr string) *http.Request {
    wholeUrl := "http://" + host + "/v8/finance/chart/" + code + "?formatted=true&interval=" + interval + "&range=" +
        rangeStr

    util.Logger.Info(wholeUrl)
    req, _ := http.NewRequest("GET", wholeUrl, nil)

    req.Header.Add("User-Agent", uarand.GetRandom())
    req.Header.Add("Host", host)
    req.Header.Add("Connection", "keep-alive")
    req.Header.Add("Pragma", "no-cache")
    req.Header.Add("Cache-Control", "no-cache")
    req.Header.Add("Upgrade-Insecure-Requests", "1")
    req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
    req.Header.Add("Accept-Encoding", "gzip, deflate, sdch")
    req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
    return req
}



func Get(code, interval, rangeStr string) (respData RespResult, err error) {
    if strings.Contains(code, ".") {
        code = strings.Replace(code, ".", "-", -1)
    }

    for _, host := range hosts {
        req := newRequest(host, code, interval, rangeStr)
        var resp *http.Response
        resp, err = client.Do(req)
        if err != nil {
            continue
        }

        var reader io.ReadCloser
        switch resp.Header.Get("Content-Encoding") {
        case "gzip":
            reader, err = gzip.NewReader(resp.Body)
            if err != nil {
                return
            }
        default:
            reader = resp.Body
        }
        var body []byte
        body, err = ioutil.ReadAll(reader)
        if err != nil {
            continue
        }
        reader.Close()
        resp.Body.Close()
        err = json.Unmarshal(body, &respData)
        if err != nil {
            continue
        }
        if respData.Chart.Error == nil {
            break
        } else {
            err = errors.New(*respData.Chart.Error)
        }
    }
    return
}

