package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Structure to store photo information
type Photo struct {
	Path     string
	Filename string
}

// Data to pass to the template
type PageData struct {
	Photos     []Photo
	CurrentDir string
}

// Data to pass to the single photo view
type SinglePhotoData struct {
	Photo      Photo
	CurrentDir string
}

// Supported image file extensions
var validExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}

//go:embed static templates
var content embed.FS

func main() {

	staticFS, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files (CSS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	// Serve photo directory
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("pictures"))))

	// Main page handler
	http.HandleFunc("/", indexHandler)

	// Single photo view handler
	http.HandleFunc("/view/", viewHandler)

	fmt.Println("Server started: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Main page handler (photo list display)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Get photos
	photos, err := getPhotos("pictures")
	if err != nil {
		http.Error(w, "Failed to load photos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare template data
	data := PageData{
		Photos:     photos,
		CurrentDir: "pictures",
	}

	// Render template
	tmpl, err := template.ParseFS(content, "templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
	}
}

// Single photo view handler
func viewHandler(w http.ResponseWriter, r *http.Request) {
	photoPath := strings.TrimPrefix(r.URL.Path, "/view/")
	
	// Check for invalid paths
	if !strings.HasPrefix(photoPath, "pictures/") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	// Create photo information
	photo := Photo{
		Path:     photoPath,
		Filename: filepath.Base(photoPath),
	}

	// Prepare template data
	data := SinglePhotoData{
		Photo:      photo,
		CurrentDir: filepath.Dir(photoPath),
	}

	// Render template
	tmpl, err := template.ParseFS(content, "templates/view.html")
	if err != nil {
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
	}
}

// Get a list of photo files from the specified directory
func getPhotos(dir string) ([]Photo, error) {
	var photos []Photo

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check file extension
		ext := strings.ToLower(filepath.Ext(path))
		if validExtensions[ext] {
			photo := Photo{
				Path:     path,
				Filename: info.Name(),
			}
			photos = append(photos, photo)
		}

		return nil
	})

	return photos, err
}