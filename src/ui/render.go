package ui

import (
	"io"
	"net/http"
	"os"
)

func Static() http.Handler {
	return http.FileServer(http.Dir("./src/ui/static"))
}

func Render(w http.ResponseWriter, r *http.Request) {
	filePath := "./src/ui/static/index.html" // Update with the path to your HTML file

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
}
