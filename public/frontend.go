package public

import "embed"

//go:embed xindex.html dist static
var Content embed.FS
