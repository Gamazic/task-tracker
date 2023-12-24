package microframework

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func defaultJson(msg string) string {
	return fmt.Sprintf(`{"message":"%s"}`, msg)
}

type ResponseBuilder struct {
	status  int
	headers map[string]string
	body    []byte
	errBody error

	writer http.ResponseWriter
}

func NewResponseBuilder(w http.ResponseWriter) ResponseBuilder {
	return ResponseBuilder{writer: w}
}

func (r ResponseBuilder) BuildHeader(h map[string]string) ResponseBuilder {
	r.headers = h
	return r
}

func (r ResponseBuilder) BuildStatus(s int) ResponseBuilder {
	r.status = s
	return r
}

func (r ResponseBuilder) BuildBody(body any) ResponseBuilder {
	rawBody, err := json.Marshal(body)
	if err != nil {
		r.errBody = err
		return r
	}
	r.body = rawBody
	return r
}

func (r ResponseBuilder) BuildBodyPlainMsg(msg string) ResponseBuilder {
	r.body = []byte(msg)
	return r
}

func (r ResponseBuilder) BuildBodyNestedMsg(msg string) ResponseBuilder {
	r.body = []byte(defaultJson(msg))
	return r
}

func (r ResponseBuilder) Send() error {
	if r.errBody != nil {
		return r.errBody
	}
	r.writer.Header().Set("Content-Type", "application/json")
	for k, v := range r.headers {
		r.writer.Header().Set(k, v)
	}
	r.writer.WriteHeader(r.status)
	n, err := r.writer.Write(r.body)
	defer r.writer.Write([]byte("\n"))
	if err != nil {
		return err
	}
	if n != len(r.body) {
		return fmt.Errorf("expected %d bytes written, wrote %d", len(r.body), n)
	}
	return nil
}

func SendInternalServerError(w http.ResponseWriter) {
	NewResponseBuilder(w).
		BuildStatus(http.StatusInternalServerError).
		BuildBodyNestedMsg("Internal Server Error").
		Send()
}

func SendValidationError(w http.ResponseWriter, err error) {
	NewResponseBuilder(w).
		BuildStatus(http.StatusInternalServerError).
		BuildBodyNestedMsg(err.Error()).
		Send()
}

func SendForbidden(w http.ResponseWriter) {
	NewResponseBuilder(w).
		BuildStatus(http.StatusForbidden).
		BuildBodyNestedMsg("Forbidden").
		Send()
}
