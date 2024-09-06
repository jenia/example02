package main

import
(
	"log"
	"net/http"
	"archie/server"
	"archie/image"
)

func main(){
	im := &image.ImageManipulator{Quality: 50}
	s := server.Server{
		Im:im,
	}
	http.HandleFunc("/resize", func(w http.ResponseWriter, r *http.Request) {
		s.Resize(w, r)
	})

	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		s.Convert(w, r)
	})

	http.HandleFunc("/compress", func(w http.ResponseWriter, r *http.Request) {
		s.Compress(w, r)
	})

    log.Fatal(http.ListenAndServe(":8080", nil))
}
