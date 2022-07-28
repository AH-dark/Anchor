package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDotPathToStandardPath(t *testing.T) {
	asserts := assert.New(t)

	asserts.Equal("/", DotPathToStandardPath(""))
	asserts.Equal("/目录", DotPathToStandardPath("目录"))
	asserts.Equal("/目录/目录2", DotPathToStandardPath("目录,目录2"))
}

func TestFillSlash(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal("/", FillSlash("/"))
	asserts.Equal("/", FillSlash(""))
	asserts.Equal("/123/", FillSlash("/123"))
}

func TestRemoveSlash(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal("/", RemoveSlash("/"))
	asserts.Equal("/123/1236", RemoveSlash("/123/1236"))
	asserts.Equal("/123/1236", RemoveSlash("/123/1236/"))
}

func TestSplitPath(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal([]string{}, SplitPath(""))
	asserts.Equal([]string{}, SplitPath("1"))
	asserts.Equal([]string{"/"}, SplitPath("/"))
	asserts.Equal([]string{"/", "123", "321"}, SplitPath("/123/321"))
}

func TestFormSlash(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal("/", FormSlash("\\"))
	asserts.Equal("/a/b/c/d", FormSlash("\\a\\b\\c\\d"))
}

func TestFileHasMinSuffix(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal(false, FileHasMinSuffix("/test/test2/testfile.js"))
	asserts.Equal(false, FileHasMinSuffix("testfile.js"))
	asserts.Equal(true, FileHasMinSuffix("/test/test2/testfileb.min.js"))
	asserts.Equal(true, FileHasMinSuffix("testfileb.min.js"))
	asserts.Equal(false, FileHasMinSuffix("/test/test2/testfileb.aaaaa.bbbbb.ccccc.js"))
}

func TestRemoveMinSuffix(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal("/test/test2/testfile.js", RemoveMinSuffix("/test/test2/testfile.js"))
	asserts.Equal("testfile.js", RemoveMinSuffix("testfile.js"))
	asserts.Equal("/test/test2/testfileb.js", RemoveMinSuffix("/test/test2/testfileb.min.js"))
	asserts.Equal("testfileb.js", RemoveMinSuffix("testfileb.min.js"))
	asserts.Equal("/test/test2/testfileb.aaaaa.bbbbb.ccccc.js", RemoveMinSuffix("/test/test2/testfileb.aaaaa.bbbbb.ccccc.js"))
	asserts.Equal("/source", RemoveMinSuffix("/source"))
}

func TestExtension(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal("html", Extension("/aaa/bbb/index.html"))
	asserts.Equal("js", Extension("/aaa/bbb/jq.min.js"))
	asserts.Equal("css", Extension("style.css"))
}
