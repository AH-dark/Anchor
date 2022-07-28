// Package utils; From https://github.com/cloudreve/Cloudreve/blob/master/pkg/util/path.go
package utils

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// DotPathToStandardPath 将","分割的路径转换为标准路径
func DotPathToStandardPath(path string) string {
	return "/" + strings.Replace(path, ",", "/", -1)
}

// FillSlash 给路径补全`/`
func FillSlash(path string) string {
	if path == "/" {
		return path
	}
	return path + "/"
}

// RemoveSlash 移除路径最后的`/`
func RemoveSlash(path string) string {
	if len(path) > 1 {
		return strings.TrimSuffix(path, "/")
	}
	return path
}

// SplitPath 分割路径为列表
func SplitPath(path string) []string {
	if len(path) == 0 || path[0] != '/' {
		return []string{}
	}

	if path == "/" {
		return []string{"/"}
	}

	pathSplit := strings.Split(path, "/")
	pathSplit[0] = "/"
	return pathSplit
}

// FormSlash 将path中的反斜杠'\'替换为'/'
func FormSlash(old string) string {
	return path.Clean(strings.ReplaceAll(old, "\\", "/"))
}

// RelativePath 获取相对可执行文件的路径
func RelativePath(name string) string {
	if filepath.IsAbs(name) {
		return name
	}
	e, _ := os.Executable()
	return filepath.Join(filepath.Dir(e), name)
}

// FileHasMinSuffix 文件第二段后缀是否为`min`
func FileHasMinSuffix(path string) bool {
	p := strings.Split(path, "/")
	filename := p[len(p)-1]
	f := strings.Split(filename, ".")
	return len(f) > 2 && f[len(f)-2] == "min"
}

// RemoveMinSuffix 去除第二段`min`后缀
func RemoveMinSuffix(path string) string {
	if !FileHasMinSuffix(path) {
		return path
	}

	arr := strings.Split(path, "/")
	ext := Extension(path)
	last := arr[len(arr)-1]
	for strings.HasSuffix(last, ".min."+ext) {
		last = strings.Replace(last, ".min."+ext, "."+ext, 1)
	}
	arr[len(arr)-1] = last

	return strings.Join(arr, "/")
}

// Extension 获取文件扩展名
func Extension(path string) string {
	p := strings.Split(path, "/")
	filename := p[len(p)-1]
	f := strings.Split(filename, ".")
	if len(f) <= 1 {
		return ""
	}

	return f[len(f)-1]
}
