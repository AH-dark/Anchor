package services

import (
	"fmt"
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetNpmRawFile 从 Npm 获取源文件信息
func GetNpmRawFile(user string, pkg string, version string, path string) []byte {
	for _, endpoint := range conf.Config.Proxy.Npm.Endpoint {
		if user != "" {
			pkg = fmt.Sprintf("@%s/%s", user, pkg)
		}

		if version == "" {
			version = "latest"
		}

		url := utils.Replace(map[string]string{
			"{{package}}": pkg,
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

			return body
		} else {
			utils.Log().Debug("向 %s 发送 GET 请求时错误，返回状态码 %d", url, resp.StatusCode)
		}
	}

	return nil
}

// CheckNpmWhiteList 检验仓库信息是否在白名单内
func CheckNpmWhiteList(user string, pkg string) bool {
	if len(conf.Config.Proxy.Npm.WhiteList) == 0 {
		return true
	}

	for _, v := range conf.Config.Proxy.Npm.WhiteList {
		t := strings.Split(v, "/")
		if len(t) != 2 {
			continue
		}

		u := strings.TrimPrefix(t[0], "@")
		p := t[1]

		if u == "*" {
			return true
		}

		if u == user {
			if p == pkg {
				return true
			} else if p == "*" {
				return true
			}
		}
	}

	return false
}
