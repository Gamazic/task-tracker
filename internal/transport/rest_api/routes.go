package rest_api

import (
	"net/http"
	"regexp"
	"strconv"
	"tracker_backend/internal/transport/rest_api/register_controller"
	"tracker_backend/internal/transport/rest_api/task_controller"
)

var (
	registerPattern, _       = regexp.Compile("^/api/register/?$")
	taskCollectionPattern, _ = regexp.Compile("^/api/tasks/?$")
	taskObjPattern, _        = regexp.Compile("^/api/tasks/([0-9]+)/?$")
	swaggerPattern, _        = regexp.Compile("^/docs/?")
)

type MainHandler struct {
	RegisterController register_controller.RegisterController
	TaskController     task_controller.TaskController
	SwaggerHandler     http.Handler
}

func (m MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var taskId int
	switch {
	case match(r.URL.Path, swaggerPattern) && r.Method == http.MethodGet:
		http.StripPrefix("/docs", m.SwaggerHandler).ServeHTTP(w, r)
	case match(r.URL.Path, registerPattern) && r.Method == http.MethodPost:
		m.RegisterController.Post(w, r)
	case match(r.URL.Path, taskCollectionPattern) && r.Method == http.MethodPost:
		m.TaskController.Post(w, r)
	case match(r.URL.Path, taskCollectionPattern) && r.Method == http.MethodGet:
		m.TaskController.GetCollection(w, r)
	case match(r.URL.Path, taskObjPattern, &taskId) && r.Method == http.MethodPatch:
		m.TaskController.Patch(w, r, taskId)
	default:
		http.NotFound(w, r)
	}
}

func match(path string, pattern *regexp.Regexp, vars ...any) bool {
	matches := pattern.FindStringSubmatch(path)
	if len(matches) <= 0 {
		return false
	}
	for i, m := range matches[1:] {
		switch p := vars[i].(type) {
		case *string:
			*p = m
		case *int:
			n, err := strconv.Atoi(m)
			if err != nil {
				return false
			}
			*p = n
		default:
			panic("vars must be *string or *int")
		}
	}
	return true
}
