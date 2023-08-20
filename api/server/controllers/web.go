package controllers

import (
	"log"
	"moonlogs/web"
	"net/http"
	"strings"
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

	contentType := http.DetectContentType(data)
	if strings.Contains(filePath, ".js") {
		contentType = "text/javascript; charset=UTF-8"
	}

	w.Header().Set("Content-Type", contentType)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}
