package compress

import (
	"github.com/AH-dark/Anchor/pkg/conf"
	"github.com/AH-dark/Anchor/pkg/utils"
	"strings"
)

var compressible = []string{
	"application/javascript",
	"text/javascript",
	"text/css",
	"text/html",
	"application/json",
}

func CanBeCompressed(path string, contentType string, minify conf.MinifyType) bool {
	if minify == conf.MinifyNone {
		return false
	} else if minify == conf.MinifyOnlyMin && !utils.FileHasMinSuffix(path) {
		return false
	}

	for _, s := range compressible {
		if strings.Contains(contentType, s) {
			return true
		}
	}

	return false
}
