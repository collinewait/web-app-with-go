package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GzipMiddleware struct {
	Next http.Handler
}

func (gm *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}

	encoding := r.Header.Get("Accept-Encoding")
	if !strings.Contains(encoding, "gzip") {
		gm.Next.ServeHTTP(w, r)
		return
	}
	w.Header().Add("Content-Encoding", "gzip")
	gzipwriter := gzip.NewWriter(w)
	defer gzipwriter.Close()

	grw := gzipResponseWriter{
		ResponseWriter: w,
		Writer:         gzipwriter,
	}
	gm.Next.ServeHTTP(grw, r)
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (gzw gzipResponseWriter) Write(data []byte) (int, error) {
	return gzw.Writer.Write(data)
}
