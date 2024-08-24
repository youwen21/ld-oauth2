package asset

import (
	"embed"
)

//go:embed dist
var Dist embed.FS

//go:embed dist
var Static embed.FS

//go:embed favicon_io
var Favicon embed.FS
