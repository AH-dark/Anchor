package compress

import (
	"github.com/AH-dark/Anchor/pkg/utils"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"regexp"
)

var m *minify.M

func init() {
	m = minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.AddFunc("application/json", json.Minify)
}

func Bytes(data []byte, contentType string) []byte {
	d, err := m.Bytes(contentType, data)
	if err != nil {
		utils.Log().Error("压缩 %s 文件时错误，%s", contentType, err)
		return nil
	}

	return d
}
