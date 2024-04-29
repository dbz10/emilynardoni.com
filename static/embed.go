package static

import "embed"

//go:embed content html img css
var Static embed.FS
