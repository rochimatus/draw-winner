package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/rochimatus/draw-winner/server/http/response"
	"github.com/rochimatus/draw-winner/util"
)

func (h Handler) Draw(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query().Get("names")
	if queryParams == "" {
		response.BadRequestResponseRenderer(writer, errors.New("no names provided"))
		return
	}

	names := util.SliceFilter(func(str string) bool { return str != "" }, strings.Split(queryParams, ","))

	result, err := h.service.Draw(names)
	if err != nil {
		response.BadRequestResponseRenderer(writer, err)
		return
	}

	response.SuccessResponseRender(writer, result)
}
