// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

// loggingResponseWriter nos permite capturar el status y el tamaño de la respuesta
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode    int
	contentLength int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	n, err := lrw.ResponseWriter.Write(b)
	lrw.contentLength += n
	return n, err
}

// loggerMiddleware intercepta la petición, mide el tiempo y escribe la línea de log
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Inicializamos con 200 OK por defecto, ya que si el manejador
		// no llama a WriteHeader, Go asume un 200.
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Procesar la petición
		next.ServeHTTP(lrw, r)

		// Calcular el tiempo transcurrido
		duration := time.Since(startTime)

		// Format: Date/Time RemoteAddr Status Method URL (size, time)
		fmt.Printf("%s %s %d %s %s => (%d bytes, %v)\n",
			startTime.Format("2006-01-02 15:04:05"),
			r.RemoteAddr,
			lrw.statusCode,
			r.Method,
			r.URL.Path,
			lrw.contentLength,
			duration,
		)
	})
}

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	http.Handle("/", loggerMiddleware(http.FileServer(http.Dir(*dir))))
	err := http.ListenAndServe(*listen, nil)
	log.Fatalln(err)
}
