package rest

import (
	"net/http"
	"regexp"
	"strconv"
	"tracker_backend/src/presentation/rest/task_controller"
	"tracker_backend/src/presentation/rest/user_controller"
)

var (
	userPattern, _           = regexp.Compile("^/api/user/?$")
	taskCollectionPattern, _ = regexp.Compile("^/api/task/?$")
	taskObjPattern, _        = regexp.Compile("^/api/task/([0-9]+)/?$")
	swaggerPattern, _        = regexp.Compile("^/docs")
)

type MainHandler struct {
	UserHandler user_controller.UserHandler
	TaskHandler task_controller.TaskHandler
	SwaggerDir  string
}

func (m MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var taskId int
	switch {
	case match(r.URL.Path, swaggerPattern) && r.Method == http.MethodGet:
		http.StripPrefix("/docs", http.FileServer(http.Dir(m.SwaggerDir))).ServeHTTP(w, r)
	case match(r.URL.Path, userPattern) && r.Method == http.MethodPost:
		m.UserHandler.Post(w, r)
	case match(r.URL.Path, taskCollectionPattern) && r.Method == http.MethodPost:
		m.TaskHandler.Post(w, r)
	case match(r.URL.Path, taskCollectionPattern) && r.Method == http.MethodGet:
		m.TaskHandler.GetCollection(w, r)
	case match(r.URL.Path, taskObjPattern, &taskId) && r.Method == http.MethodPatch:
		m.TaskHandler.Patch(w, r, taskId)
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
