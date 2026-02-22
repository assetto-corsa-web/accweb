package frontend

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var Content embed.FS

type DistFS struct {
	root fs.FS
}

func (d *DistFS) Open(name string) (fs.File, error) {
	return d.root.Open("dist/" + name)
}

func NewDistFS(root fs.FS) *DistFS {
	return &DistFS{
		root: root,
	}
}
