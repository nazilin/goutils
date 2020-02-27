package rest

import (
	resty "github.com/go-resty/resty/v2"
)
type restsession struct {
	url    string
	status bool
	client *resty.Client
}

func NewRestSession(url string) *restsession {
	return &restsession{
		url:    url,
		status: false,
		client: resty.New(),
	}
}
func (s *restsession) Post(jsonStr string, headers map[string]string, params map[string]string) (string, map[string][]string, error) {

	resp, err := s.populateInputHeaders(headers, s.client.R()).SetBody(jsonStr).Post(s.url)
	if err != nil || resp == nil {
		return "", nil, err
	}
	return resp.String(), s.extractResponseHeader(resp), nil

}
func (s *restsession) Get(headers map[string]string, params map[string]string) (string, map[string][]string, error) {
	return "", nil, nil
}
func (s *restsession) populateInputHeaders(headers map[string]string, request *resty.Request) *resty.Request {
	if headers != nil {
		for key, value := range headers {
			request = request.SetHeader(key, value)
		}
	}
	return request
}

func (s *restsession) extractResponseHeader(resp interface{}) map[string][]string {
	if resp != nil {
		return nil
	}
	return nil
}
