package http_server

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type ServerWrapper[Req any, Resp any] struct {
	fn func(context.Context, Req) (Resp, error)
}

func NewServerWrapper[Req any, Resp any](fn func(context.Context, Req) (Resp, error)) *ServerWrapper[Req, Resp] {
	return &ServerWrapper[Req, Resp]{fn: fn}
}

func (s *ServerWrapper[Req, Resp]) ServeHTTP(resWriter http.ResponseWriter, httpReq *http.Request, _ map[string]string) {
	ctx := httpReq.Context()

	limitedReader := io.LimitReader(httpReq.Body, 1_000_000)

	var request Req
	err := json.NewDecoder(limitedReader).Decode(&request)
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		writeErrorText(resWriter, "decoding JSON", err)
		return
	}

	response, err := s.fn(ctx, request)
	if err != nil {
		resWriter.WriteHeader(http.StatusInternalServerError)
		writeErrorText(resWriter, "running handler", err)
		return
	}

	rawJSON, err := json.Marshal(response)
	if err != nil {
		resWriter.WriteHeader(http.StatusInternalServerError)
		writeErrorText(resWriter, "encoding JSON", err)
		return
	}

	resWriter.Header().Add("Content-Type", "application/json")
	_, _ = resWriter.Write(rawJSON)
}

func writeErrorText(w http.ResponseWriter, text string, err error) {
	buf := bytes.NewBufferString(text)

	buf.WriteString(": ")
	buf.WriteString(err.Error())
	buf.WriteByte('\n')

	w.Write(buf.Bytes())
}
