package frpc

import (
	"embed"

	"github.com/xinda/desk/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
