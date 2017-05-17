package httplog

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"time"
)

type logResponseWriter struct {
	w       http.ResponseWriter
	status  int
	written int
}

func (l *logResponseWriter) Header() http.Header {
	return l.w.Header()
}

func (l *logResponseWriter) Write(data []byte) (int, error) {
	n, err := l.w.Write(data)
	l.written += n
	return n, err
}

func (l *logResponseWriter) WriteHeader(status int) {
	l.w.WriteHeader(status)
	l.status = status
}

func (l *logResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return l.w.(http.Hijacker).Hijack()
}

type logHandler struct {
	handler http.Handler
}

// New ...
func New(handler http.Handler) http.Handler {
	return &logHandler{handler: handler}
}

func (h *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ts := time.Now()
	logger := logResponseWriter{w: w, status: http.StatusOK}
	h.handler.ServeHTTP(&logger, r)
	username := "-"
	if r.URL.User != nil && r.URL.User.Username() != "" {
		username = r.URL.User.Username()
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}
	log.Printf("%s - %s [%s] \"%s %s %s\" %d %d",
		host,
		username,
		ts.Format("02/Jan/2006:15:04:05 -0700"),
		r.Method,
		r.RequestURI,
		r.Proto,
		logger.status,
		logger.written,
	)
}
