package httptool

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/yanlong-l/go-mall/common/errcode"
	"github.com/yanlong-l/go-mall/common/logger"
	"github.com/yanlong-l/go-mall/common/util"
	"io/ioutil"
	"net/http"
	"time"
)

type requestOption struct {
	ctx     context.Context
	timeout time.Duration
	data    []byte
	headers map[string]string
}

func defaultRequestOptions() *requestOption {
	return &requestOption{
		ctx:     context.Background(),
		timeout: time.Second * 5,
		data:    nil,
		headers: make(map[string]string),
	}
}

type Option interface {
	apply(opt *requestOption) error
}

type optionFunc func(opt *requestOption) error

func (f optionFunc) apply(opt *requestOption) error {
	return f(opt)
}

func WithContext(ctx context.Context) Option {
	return optionFunc(func(opts *requestOption) (err error) {
		opts.ctx = ctx
		return
	})
}

func WithTimeout(timeout time.Duration) Option {
	return optionFunc(func(opts *requestOption) (err error) {
		opts.timeout, err = timeout, nil
		return
	})
}

func WithHeaders(headers map[string]string) Option {
	return optionFunc(func(opts *requestOption) (err error) {
		for k, v := range headers {
			opts.headers[k] = v
		}
		return
	})
}

func WithData(data []byte) Option {
	return optionFunc(func(opts *requestOption) (err error) {
		opts.data, err = data, nil
		return
	})
}

func Request(method string, url string, options ...Option) (httpStatusCode int, respBody []byte, err error) {
	start := time.Now()
	reqOpts := defaultRequestOptions() // 默认的请求选项
	ctx := reqOpts.ctx
	for _, opt := range options { // 在reqOpts上应用通过options设置的选项
		err = opt.apply(reqOpts)
		if err != nil {
			return
		}
	}
	defer func() {
		if err != nil {
			logger.Error(ctx, "HTTP_REQUEST_ERROR_LOG", "method", method, "url", url, "body", reqOpts.data, "reply", respBody, "err", err)
		}
	}()
	// 创建请求对象
	req, err := http.NewRequest(method, url, bytes.NewReader(reqOpts.data))
	if err != nil {
		return
	}
	req = req.WithContext(reqOpts.ctx)
	defer req.Body.Close()

	// 在Header中添加追踪信息 把内部服务串起来
	traceId, spanId, _ := util.GetTraceInfoFromCtx(reqOpts.ctx)
	reqOpts.headers["traceid"] = traceId
	reqOpts.headers["spanid"] = spanId
	if len(reqOpts.headers) != 0 { // 设置请求头
		for key, value := range reqOpts.headers {
			req.Header.Add(key, value)
		}
	}
	// 发起请求
	client := &http.Client{Timeout: reqOpts.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// 记录请求日志
	dur := time.Since(start).Milliseconds()
	if dur >= 3000 { // 超过 3秒 返回, 记一条 Warn 日志
		logger.Warn(ctx, "HTTP_REQUEST_SLOW_LOG", "method", method, "url", url, "body", reqOpts.data, "reply", respBody, "err", err, "dur/ms", dur)
	} else {
		logger.Debug(ctx, "HTTP_REQUEST_DEBUG_LOG", "method", method, "url", url, "body", reqOpts.data, "reply", respBody, "err", err, "dur/ms", dur)
	}

	httpStatusCode = resp.StatusCode
	if httpStatusCode != http.StatusOK {
		// 返回非 200 时Go的 http 库不回返回error, 这里处理成error 调用方好判断
		err = errcode.Wrap("request api error", errors.New(fmt.Sprintf("non 200 response, response code: %d", httpStatusCode)))
		return
	}

	respBody, _ = ioutil.ReadAll(resp.Body)
	return
}

// Get 发起GET请求
func Get(ctx context.Context, url string, options ...Option) (httpStatusCode int, respBody []byte, err error) {
	options = append(options, WithContext(ctx))
	return Request("GET", url, options...)
}

// Post 发起POST请求
func Post(ctx context.Context, url string, data []byte, options ...Option) (httpStatusCode int, respBody []byte, err error) {
	// 默认自带Header Content-Type: application/json 可通过 传递 WithHeaders 增加或者覆盖Header信息
	defaultHeader := map[string]string{"Content-Type": "application/json"}
	var newOptions []Option
	newOptions = append(newOptions, WithHeaders(defaultHeader), WithData(data), WithContext(ctx))
	newOptions = append(newOptions, options...)

	httpStatusCode, respBody, err = Request("POST", url, newOptions...)
	return
}
