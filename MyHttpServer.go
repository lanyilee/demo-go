package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandle(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("the time is:" + tm))
}
func timeHandle2(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("the time is:" + tm))
	}
	return http.HandlerFunc(fn)
}
func timeHandle3(format string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("the time is:" + tm))
	})
}

func timeHandle4(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("the time is:" + tm))
	}
}

func main() {
	mux := http.NewServeMux()
	th := timeHandle4(time.RFC1123)
	mux.HandleFunc("/time", th)
	th2 := timeHandle2(time.RFC822)
	mux.Handle("/time2", th2)
	log.Println("Listening...")
	http.ListenAndServe("127.0.0.1:8000", mux)

}
