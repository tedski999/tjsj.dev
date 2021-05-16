package webstats

import "net/http"

type StatsResponseWriter struct {
	http.ResponseWriter
	status, length int
	isHeaderWritten bool
}

func (w *StatsResponseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	if !w.isHeaderWritten {
		w.WriteHeader(http.StatusOK)
	}
	w.length += n
	return n, err
}

func (w *StatsResponseWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	if !w.isHeaderWritten {
		w.isHeaderWritten = true
		w.status = status
	}
}
