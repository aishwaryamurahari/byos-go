package server

import (
	"mime"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// ServeStaticFile attempts to read and return a file from the ./static directory
func ServeStaticFile(requestPath string) (string, string, int) {
	// URL decode in case of spaces, %20, etc.
	cleanPath, err := url.PathUnescape(requestPath)
	if err != nil {
		return "text/plain", "Bad Request", 400
	}

	// Prevent directory traversal like ../../etc/passwd
	if strings.Contains(cleanPath, "..") {
		return "text/plain", "Forbidden", 403
	}

	// Remove /static/ prefix and join with local static directory
	relPath := strings.TrimPrefix(cleanPath, "/static/")
	fullPath := filepath.Join("static", relPath)

	// Read the file
	data, err := os.ReadFile(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "text/plain", "File Not Found", 404
		}
		return "text/plain", "Internal Server Error", 500
	}

	// Guess content type from file extension
	ext := filepath.Ext(fullPath)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return contentType, string(data), 200
}
