package server

import (
	"io"
	"net/http"
	"strconv"
)

type Resizer interface {
	Resize(img []byte, x, y int) ([]byte, error)
	Convert(img []byte) ([]byte, error)
	Compress(img []byte) ([]byte, error)
}

// TODO: consider using a JSON error message
type Server struct {
	Im Resizer
}

func (s *Server) Resize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Unable to parse multipart form", http.StatusBadRequest)
		return
	}

	x := r.FormValue("x")
	y := r.FormValue("y")
	xInt, err := strconv.Atoi(x)
	if err != nil {
		http.Error(w, "Invalid value for x", http.StatusBadRequest)
		return
	}

	yInt, err := strconv.Atoi(y)
	if err != nil {
		http.Error(w, "Invalid value for y", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read image file", http.StatusInternalServerError)
		return
	}
	
	resizedImage, err := s.Im.Resize(imageBytes, xInt, yInt)
	if err != nil {
		http.Error(w, "Unable to resize image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")

	if _, err := w.Write(resizedImage); err != nil {
		http.Error(w, "Unable to write image to response", http.StatusInternalServerError)
	}
}

func (s *Server) Convert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Unable to parse multipart form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read image file", http.StatusInternalServerError)
		return
	}
	
	convertedImage, err := s.Im.Convert(imageBytes)
	if err != nil {
		http.Error(w, "Unable to resize image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")

	if _, err := w.Write(convertedImage); err != nil {
		http.Error(w, "Unable to write image to response", http.StatusInternalServerError)
	}
}

func (s *Server) Compress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Unable to parse multipart form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read image file", http.StatusInternalServerError)
		return
	}
	
	compressedImage, err := s.Im.Convert(imageBytes)
	if err != nil {
		http.Error(w, "Unable to resize image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")

	if _, err := w.Write(compressedImage); err != nil {
		http.Error(w, "Unable to write image to response", http.StatusInternalServerError)
	}
}
