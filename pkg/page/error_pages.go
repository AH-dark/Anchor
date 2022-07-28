package page

import (
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
)

func ErrorPage(title string) string {
	const html = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf8"/>
<title>{{title}} - {{siteName}}</title>
<style>
body div {
  text-align: center;
  width: 100%;
}
</style>
</head>
<body>
<div>
<h1>{{title}}</h1>
<hr/>
<p>{{siteName}}</p>
</div>
</body>`

	return utils.Replace(map[string]string{
		"{{title}}":    title,
		"{{siteName}}": conf.Config.System.Name,
	}, html)
}

func NotFound() string {
	return ErrorPage("404 Not Found")
}
