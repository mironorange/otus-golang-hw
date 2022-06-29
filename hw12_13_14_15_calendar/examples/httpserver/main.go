package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

type HttpServerHandler struct {

}

type StatusRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.StatusCode = status
	r.ResponseWriter.WriteHeader(status)
}

func (h *HttpServerHandler) Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello, World!"))
}

// (/) IP клиента;
// (/) дата и время запроса;
// (/) метод, path и версия HTTP;
// (/) код ответа;
// (/) latency (время обработки запроса, посчитанное, например, с помощью middleware);
// (/) user agent, если есть.
func Latency(handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &StatusRecorder{
			ResponseWriter: w,
			StatusCode:     200,
		}
		handleFunc(recorder, r)
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		log.Printf("%s %s %s %s %d %s \"%s\"", ip, r.Method, r.RequestURI, r.Proto, recorder.StatusCode, time.Since(start), r.UserAgent())
	}
}

func main() {
	handler := &HttpServerHandler{}

	mux := http.NewServeMux()
	mux.HandleFunc("/", Latency(handler.Hello))

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	//ctx := context.Background()
	//if err := server.Shutdown(ctx); err != nil {
	//	log.Fatal(err)
	//}
}
