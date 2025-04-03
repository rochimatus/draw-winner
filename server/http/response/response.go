package response

import (
	"encoding/json"
	"github.com/rochimatus/draw-winner/logger"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func StdResponseBody(success bool, data interface{}) ([]byte, error) {
	response := Response{
		success,
		data,
	}

	return json.Marshal(response)
}

func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func ResponseRenderer(writer http.ResponseWriter, response *Response, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if response.Data != nil {
		marshalledResponse, err := response.Marshal()
		if err != nil {
			logger.Error(err, "error in marshalling response")
		}

		if _, err = writer.Write(marshalledResponse); err != nil {
			logger.Error(err, "response writing failed.")
		}
	} else {
		logger.Info("No data is written in HTTP response since data was <nil>")
	}
}

type ErrResponse struct {
	Error string `json:"error"`
}

func BadRequestResponseRenderer(writer http.ResponseWriter, err error) {
	ResponseRenderer(writer, &Response{Success: false, Data: ErrResponse{Error: err.Error()}}, http.StatusBadRequest)
}

func InternalErrorResponseRenderer(writer http.ResponseWriter, err error) {
	ResponseRenderer(writer, &Response{Success: false, Data: ErrResponse{Error: err.Error()}}, http.StatusInternalServerError)
}

func SuccessResponseRender(writer http.ResponseWriter, data interface{}) {
	ResponseRenderer(writer, &Response{Success: true, Data: data}, http.StatusOK)
}

func CreatedResponseRender(writer http.ResponseWriter, data interface{}) {
	ResponseRenderer(writer, &Response{Success: true, Data: data}, http.StatusCreated)
}

// TODO: check no content behavior on existing flow. is it harmless if the Response set to nil or not.
func SuccessWithNoContentResponseRender(writer http.ResponseWriter, _ interface{}) {
	ResponseRenderer(writer, &Response{Success: true, Data: nil}, http.StatusNoContent)
}
