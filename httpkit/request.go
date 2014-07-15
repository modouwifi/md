package httpkit

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	//"mime/multipart"
	//"net"
	"net/http"
	//"net/http/httputil"
	"net/url"
	"reflect"
)

type Request struct {
	Method    string
	URL       string
	UserAgent string
	Header    *http.Header

	req *http.Request
	res *http.Response
}

func New(method, url, userAgent string) *Request {
	var request Request
	var req http.Request
	request.req = &req
	req.Header = http.Header{}
	request.Method = method
	request.URL = url
	request.UserAgent = userAgent
	return &request
}

func (r *Request) GetReq() *http.Request {
	return r.req
}

func (r *Request) GetRes() *http.Response {
	return r.res
}

func (r *Request) SetCookie(c *http.Cookie) error {
	r.req.Header.Add("Cookie", c.String())
	return nil
}

func (r *Request) SetTransport() {}

func (r *Request) Param() {}

func (r *Request) SetProtocolVersion(vers string) {
	if len(vers) == 0 {
		vers = "HTTP/1.1"
	}
	major, minor, ok := http.ParseHTTPVersion(vers)
	if ok {
		r.req.Proto = vers
		r.req.ProtoMajor = major
		r.req.ProtoMinor = minor
	}
}

func (r *Request) Body(data interface{}) {
	switch t := data.(type) {
	case nil:
	case string:
		bf := bytes.NewBufferString(t)
		r.req.Body = ioutil.NopCloser(bf)
		r.req.ContentLength = int64(bf.Len())
	case []byte:
		bf := bytes.NewBuffer(t)
		r.req.Body = ioutil.NopCloser(bf)
		r.req.ContentLength = int64(bf.Len())
	case io.ReadCloser:
		r.req.Body = t
	default:
		v := reflect.ValueOf(data)
		if !v.IsValid() {
			break
		}
		if v.Type().Kind() == reflect.Ptr {
			v = reflect.Indirect(v)
			if !v.IsValid() {
				break
			}
		}

		j, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		bf := bytes.NewBuffer(j)
		r.req.Body = ioutil.NopCloser(bf)
		r.req.ContentLength = int64(bf.Len())
		r.req.Header.Set("Content-Type", "application/json")
	}
}

func (r *Request) Do() (*http.Response, error) {
	url, err := url.Parse(r.URL)
	r.req.Method = r.Method
	r.req.URL = url

	if r.UserAgent != "" {
		r.req.Header.Set("User-Agent", r.UserAgent)
	}

	/*
		dump, err := httputil.DumpRequest(r.req, true)
		if err != nil {
			println(err.Error())
		}
		println(string(dump))
	*/

	var jar http.CookieJar
	client := &http.Client{
		//Transport: trans,
		Jar: jar,
	}

	resp, err := client.Do(r.req)
	if err != nil {
		return nil, err
	}
	r.res = resp
	return resp, nil
}

func (r *Request) Bytes() ([]byte, error) {
	resp, err := r.Do()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Request) String() (string, error) {
	data, err := r.Bytes()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (r *Request) ToJSON(v interface{}) error {
	data, err := r.Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

func (r *Request) Response() (*http.Response, error) {
	return r.Do()
}

func (r *Request) Pipe() {}
