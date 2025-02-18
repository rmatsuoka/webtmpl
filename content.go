package webtmpl

import (
	"embed"
	"io/fs"

	"github.com/rmatsuoka/webtmpl/internal/x/must"
)

//go:embed _content
var content embed.FS

func Content() fs.FS {
	return must.Do(fs.Sub(content, "_content"))
}
