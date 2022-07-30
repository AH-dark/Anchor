package services

import (
	"errors"
	"fmt"
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
	"io/ioutil"
	"net/http"
	"path/filepath"
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

// CheckNpmWhiteList 检验仓库信息是否在白名单内
func CheckNpmWhiteList(user string, pkg string) bool {
	if conf.Config.Proxy.Npm.WhiteList == nil || len(conf.Config.Proxy.Npm.WhiteList) == 0 {
		return true
	}

	for _, v := range conf.Config.Proxy.Npm.WhiteList {
		t := strings.Split(v, "/")
		switch len(t) {
		case 1:
			p := t[0]
			if p == pkg {
				return true
			}
		case 2:
			u := strings.TrimPrefix(t[0], "@")
			p := t[1]

			if u == user && p == pkg {
				return true
			}
		}
	}

	return false
}

// GetPackageInfo 从 Registry 获取包信息
func GetPackageInfo(pkg string, version string) (*NpmPackageData, error) {
	if version == "" {
		version = "latest"
	}

	url := fmt.Sprintf("https://registry.npmjs.org/%s/%s", pkg, version)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data, err := UnmarshalNpmRegistryData(body)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

const (
	ErrNoMainFlag = "no main flag"
)

func ParsePackageMainUrl(info *NpmPackageData) (string, error) {
	mainFile := info.Main
	if mainFile == nil {
		return "", errors.New(ErrNoMainFlag)
	}

	id := info.ID
	url := fmt.Sprintf("%s/", id)
	url = filepath.Join(url, *mainFile)

	return url, nil
}
