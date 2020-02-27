package util

import "github.com/nazilin/goutils/util/rest"
type RestHelper interface {
	Post(jsonStr string, headers map[string]string, params map[string]string) (string, map[string][]string, error)
	Get(headers map[string]string, params map[string]string) (string, map[string][]string, error)
}

func NewRestHelper(url string) RestHelper {
	return rest.NewRestSession(url)
}
