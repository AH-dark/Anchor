package services

import (
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetGithubRawFile 从 GitHub Raw 获取源文件信息
func GetGithubRawFile(user string, repo string, version string, path string) []byte {
	for _, endpoint := range conf.Config.Proxy.Github.Endpoint {
		url := utils.Replace(map[string]string{
			"{{user}}":    user,
			"{{repo}}":    repo,
			"{{version}}": version,
			"{{path}}":    strings.Trim(path, "/"),
		}, endpoint)

		resp, err := client.Get(url)
		if err != nil {
			utils.Log().Error("向 %s 发送 GET 请求时错误，%s", url, err)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				utils.Log().Error("解析 Body 失败，%s", err)
				continue
			}

			return body
		} else {
			utils.Log().Error("向 %s 发送 GET 请求时错误，返回状态码 %d", url, resp.StatusCode)
		}
	}

	return nil
}

// CheckGithubWhiteList 检验仓库信息是否在白名单内
func CheckGithubWhiteList(user string, repo string) bool {
	if conf.Config.Proxy.Github.WhiteList == nil || len(conf.Config.Proxy.Github.WhiteList) == 0 {
		return true
	}

	for _, v := range conf.Config.Proxy.Github.WhiteList {
		t := strings.Split(v, "/")
		if len(t) != 2 {
			continue
		}

		if t[0] == "*" {
			return true
		}

		if t[0] == user {
			if t[1] == repo {
				return true
			} else if t[1] == "*" {
				return true
			}
		}
	}

	return false
}
