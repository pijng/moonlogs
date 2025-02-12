package controllers

import (
	"log"
	"mime"
	"moonlogs/web"
	"net/http"
	"path"
)

func Web(w http.ResponseWriter, r *http.Request) {
	filePath := "build" + r.URL.Path
	if r.URL.Path == "/" {
		filePath = "build/index.html"
	}

	data, err := web.Assets.ReadFile(filePath)
	if err != nil {
		data, _ = web.Assets.ReadFile("build/index.html")
	}

	var contentType string
	extType := path.Ext(filePath)

	if extType != "" {
		contentType = mime.TypeByExtension(extType)
	} else {
		contentType = http.DetectContentType(data)
	}
	w.Header().Set("Content-Type", contentType)

	_, err = w.Write(data)
	if err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
