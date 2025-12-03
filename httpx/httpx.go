package httpx

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/bang-go/opt"
)

const (
	ContentRaw     = "Raw"            //原始请求
	ContentForm    = "Form"           //Form请求
	ContentJson    = "Json"           //Json请求
	DefaultTimeout = 30 * time.Second //默认请求时间
)
const (
	MethodGet     = http.MethodGet
	MethodHead    = http.MethodHead
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodPatch   = http.MethodPatch // RFC 5789
	MethodDelete  = http.MethodDelete
	MethodConnect = http.MethodConnect
	MethodOptions = http.MethodOptions
	MethodTrace   = http.MethodTrace
)

type Config struct {
	Timeout time.Duration
	Trace   bool
}

type Client interface {
	Send(ctx context.Context, req *Request, opts ...opt.Option[requestOptions]) (resp *Response, err error)
}

type clientEntity struct {
	config *Config
}

func New(conf *Config) Client {
	return &clientEntity{
		config: conf,
	}
}

func (c clientEntity) Send(ctx context.Context, req *Request, opts ...opt.Option[requestOptions]) (resp *Response, err error) {
	options := &requestOptions{}
	opt.Each(options, opts...)
	httpUrl, err := req.getUrl()
	if err != nil {
		return
	}
	method, err := req.getMethod()
	if err != nil {
		return
	}
	reqBody := req.getBody()
	var httpReq *http.Request
	var httpRes *http.Response
	if httpReq, err = http.NewRequestWithContext(ctx, method, httpUrl, reqBody); err != nil { //新建http请求
		return
	}
	req.setHeaders(httpReq) //init headers
	//basic auth
	if options.baseAuth != nil {
		httpReq.SetBasicAuth(options.baseAuth.Username, options.baseAuth.Password)
	}
	req.setCookie(httpReq) ////init cookie

	httpClient := http.Client{}
	if c.config.Timeout > 0 {
		httpClient.Timeout = c.config.Timeout
	}
	startTime := time.Now()
	if httpRes, err = httpClient.Do(httpReq); err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(httpRes.Body)
	endTime := time.Now()
	elapsed := endTime.Sub(startTime).Seconds()
	resp = req.packResponse(httpRes, elapsed)
	return
}
