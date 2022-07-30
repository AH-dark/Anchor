package services

import (
	"github.com/AH-dark/Anchor/pkg/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetWordpressRawFile(endpoint string, name string, version string, path string) []byte {
	url := utils.Replace(map[string]string{
		"{{name}}":    name,
		"{{version}}": version,
	}, endpoint)
	url = url + strings.Trim(path, "/")

	resp, err := client.Get(url)
	if err != nil {
		utils.Log().Error("向 %s 发送 GET 请求时错误，%s", url, err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		utils.Log().Error("向 %s 发送 GET 请求时错误，返回状态码 %d", url, resp.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Log().Error("解析 Body 失败，%s", err)
		return nil
	}

	return body
}
