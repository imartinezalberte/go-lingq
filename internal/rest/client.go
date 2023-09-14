package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/imartinezalberte/go-lingq/internal/utils"
)

type (
	UnimplementedMethodErr error

	Querier interface {
		ToQuery() (url.Values, error)
	}

	PathParameterer interface {
		ToPathParameter() (map[string]string, error)
	}

	Headerers interface {
		ToHeader() (http.Header, error)
	}

	BasicAuthenticator interface {
		ToBasic() (string, string, error)
	}

	Bodier interface {
		ToBody() (any, error)
	}

	Methoder interface {
		ToMethod() string
	}

	Requester interface {
		Methoder
		Querier
		Headerers
		PathParameterer
		BasicAuthenticator
		Bodier
	}

	// DummyRequester is used in structs that does not implement some of the Requester methods
	// to fast our development phase and reduce boilerplate code.
	DummyRequester struct{}

	GetDummyRequester    struct{ DummyRequester }
	PostDummyRequester   struct{ DummyRequester }
	DeleteDummyRequester struct{ DummyRequester }
	PutDummyRequester    struct{ DummyRequester }

	ReqResInfo struct {
		Method     string
		URL        string
		StatusCode int
		BodyRes    []byte
	}

	APIResponseErroer interface {
		APIResponseErr(error, ReqResInfo) error
	}

	DummyAPIResponseErr struct{}
)

var ErrUnimplementedMethod UnimplementedMethodErr = errors.New("Unimplemented method")

func (DummyRequester) ToQuery() (url.Values, error) {
	return url.Values{}, ErrUnimplementedMethod
}

func (DummyRequester) ToHeader() (http.Header, error) {
	return http.Header{}, ErrUnimplementedMethod
}

func (DummyRequester) ToPathParameter() (map[string]string, error) {
	return map[string]string{}, ErrUnimplementedMethod
}

func (DummyRequester) ToBasic() (string, string, error) {
	return "", "", ErrUnimplementedMethod
}

func (DummyRequester) ToBody() (any, error) {
	return nil, ErrUnimplementedMethod
}

func (GetDummyRequester) ToMethod() string {
	return http.MethodGet
}

func (PostDummyRequester) ToMethod() string {
	return http.MethodPost
}

func (DeleteDummyRequester) ToMethod() string {
	return http.MethodDelete
}

func (PutDummyRequester) ToMethod() string {
	return http.MethodPut
}

func (DummyAPIResponseErr) APIResponseErr(err error, reqRes ReqResInfo) error {
	if err != nil {
		return err
	}
	return errors.New(
		strings.Join(
			[]string{reqRes.Method, reqRes.URL, strconv.Itoa(reqRes.StatusCode)},
			utils.Space,
		),
	)
}

func MapReqRes(req *resty.Request, res *resty.Response) ReqResInfo {
	return ReqResInfo{
		Method:     req.Method,
		URL:        req.URL,
		StatusCode: res.StatusCode(),
		BodyRes:    res.Body(),
	}
}

func ReqGen[T Requester](cl *resty.Client, ctx context.Context, request T) (*resty.Request, error) {
	req := cl.R().SetContext(ctx)

	if username, pass, err := request.ToBasic(); err != nil &&
		!errors.Is(err, ErrUnimplementedMethod) {
		return req, err
	} else if err == nil {
		req = req.SetBasicAuth(username, pass)
	}

	if queryParams, err := request.ToQuery(); err != nil &&
		!errors.Is(err, ErrUnimplementedMethod) {
		return req, err
	} else if err == nil {
		req = req.SetQueryParamsFromValues(queryParams)
	}

	if pathParams, err := request.ToPathParameter(); err != nil &&
		!errors.Is(err, ErrUnimplementedMethod) {
		return req, err
	} else if err == nil {
		req = req.SetPathParams(pathParams)
	}

	if headers, err := request.ToHeader(); err != nil && !errors.Is(err, ErrUnimplementedMethod) {
		return req, err
	} else if err == nil {
		req = req.SetHeaderMultiValues(headers)
	}

	if body, err := request.ToBody(); err != nil && !errors.Is(err, ErrUnimplementedMethod) {
		return req, err
	} else if err == nil && request.ToMethod() != http.MethodGet {
		req = req.SetBody(body)
	}

	return req, nil
}

func ExecReq[E APIResponseErroer, K any](
	req *resty.Request,
	requester interface{ ToMethod() string },
	endpoint string,
) (result K, err error) {
	var (
		resErr E
		res    *resty.Response
	)
	res, err = req.SetError(&resErr).SetResult(&result).Execute(requester.ToMethod(), endpoint)
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		if e := json.Unmarshal(res.Body(), &resErr); e == nil {
			return result, resErr.APIResponseErr(err, MapReqRes(req, res))
		}
	}

	if err != nil || res.IsError() {
		return result, resErr.APIResponseErr(err, MapReqRes(req, res))
	}

	if v, ok := any(result).(APIResponseErroer); ok {
		return result, v.APIResponseErr(err, MapReqRes(req, res))
	}

	return
}

func Exec[T Requester, E APIResponseErroer, K any](
	cl *resty.Client,
	ctx context.Context,
	request T,
	endpoint string,
) (result K, err error) {
	req, err := ReqGen(cl, ctx, request)
	if err != nil {
		return result, err
	}

	return ExecReq[E, K](req, request, endpoint)
}
