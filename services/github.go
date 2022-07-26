package services

import (
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetGithubRawFile(user string, repo string, version string, path string) ([]byte, string) {
	for _, endpoint := range conf.Config.Proxy.Github.Endpoint {
		url := utils.Replace(map[string]string{
			"{{user}}":    user,
			"{{repo}}":    repo,
			"{{version}}": version,
			"{{path}}":    strings.Trim(path, "/"),
		}, endpoint)

		resp, err := client.Get(url)
		if err != nil {
			utils.Log().Debug("向 %s 发送 GET 请求时错误，%s", url, err)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				continue
			}

			return body, resp.Header.Get("Content-Type")
		} else {
			utils.Log().Debug("向 %s 发送 GET 请求时错误，返回状态码 %d", url, resp.StatusCode)
		}
	}

	return nil, ""
}
