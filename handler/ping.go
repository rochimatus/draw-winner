package handler

import (
	"net/http"

	"github.com/rochimatus/draw-winner/logger"
	"github.com/rochimatus/draw-winner/server/http/response"
)

func (h Handler) Ping(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusOK)

	responseData := response.PingResponse("pong")

	responseBody, err := response.StdResponseBody(true, responseData)
	if err != nil {
		logger.Error(err, "error in marshalling response")
		return
	}

	if _, err = writer.Write(responseBody); err != nil {
		logger.Error(err, "response writing failed")
		return
	}
}
