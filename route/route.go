package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Kafka"))
}

func Route() {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	r.Route("/test", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Kafka"))
		})
		r.Get(`/`, Test)
	})

	http.ListenAndServe(":3000", r)

}
