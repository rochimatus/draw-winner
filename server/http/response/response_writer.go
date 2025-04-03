package response

import (
	"bufio"
	"bytes"
	"errors"
	"net"
	"net/http"
)

// ResponseWriter acts as an adapter for `http.ResponseWriter` type to store response status
// statusCode and responseSize.
type Writer struct {
	http.ResponseWriter
	StatusCode   int
	ResponseSize int
	ResponseBody bytes.Buffer
}

// NewResponseWriter returns a new `ResponseWriter` type by decorating `http.ResponseWriter` type.
func NewResponseWriter(w http.ResponseWriter) *Writer {
	return &Writer{
		ResponseWriter: w,
	}
}

// WriteHeader overrides `http.ResponseWriter` type.
func (writer *Writer) WriteHeader(code int) {
	writer.StatusCode = code
	writer.ResponseWriter.WriteHeader(code)
}

// Overrides `http.ResponseWriter` type.
func (writer *Writer) Write(body []byte) (int, error) {
	if writer.Code() == 0 {
		writer.WriteHeader(http.StatusOK)
	}

	var err error

	writer.ResponseBody.Write(body)
	writer.ResponseSize, err = writer.ResponseWriter.Write(body)

	return writer.ResponseSize, err
}

// Flush Overrides `http.Flusher` type.
func (writer *Writer) Flush() {
	if flusher, ok := writer.ResponseWriter.(http.Flusher); ok {
		if writer.Code() == 0 {
			writer.WriteHeader(http.StatusOK)
		}

		flusher.Flush()
	}
}

// Hijack Overrides `http.Hijacker` type.
func (writer *Writer) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := writer.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("the hijacker interface is not supported")
	}

	return hijacker.Hijack()
}

// Code Returns response status code.
func (writer *Writer) Code() int {
	return writer.StatusCode
}

// Size Returns response responseSize.
func (writer *Writer) Size() int {
	return writer.ResponseSize
}
