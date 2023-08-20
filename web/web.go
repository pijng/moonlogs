package web

import "embed"

//go:embed build/*
var Assets embed.FS
