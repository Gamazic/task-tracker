package task_controller

import (
	"fmt"
	"tracker_backend/src/presentation/rest/microframework"
)

const usernameHeaderKey = "Username"

type UsernameHeader string

func (u UsernameHeader) Validate() error {
	if u == "" {
		return fmt.Errorf("%w: non empty header '%s' is required",
			microframework.ValidationErr, usernameHeaderKey)
	}
	return nil
}

type TaskPostRequestModel struct {
	Description string `json:"description"`
}

func (t TaskPostRequestModel) Validate() error {
	if t.Description == "" {
		return fmt.Errorf("%w: non empy field 'description' is required",
			microframework.ValidationErr)
	}
	return nil
}

var availableStages = map[string]struct{}{
	"todo":        {},
	"in_progress": {},
	"done":        {},
}

type TaskPatchRequestModel struct {
	Stage string `json:"stage"`
}

func (t TaskPatchRequestModel) Validate() error {
	if t.Stage == "" {
		return fmt.Errorf("%w: non empy field 'stage' is required",
			microframework.ValidationErr)
	}
	_, ok := availableStages[t.Stage]
	if !ok {
		stages := ""
		for k := range availableStages {
			stages += k + ","
		}
		stages = stages[:len(stages)-1]
		return fmt.Errorf("%w: stage should be one of %s",
			microframework.ValidationErr, stages)
	}
	return nil
}

type TaskIdModel struct {
	TaskId int `json:"task_id"`
}

type TaskModel struct {
	TaskId      int    `json:"task_id"`
	Description string `json:"description"`
	Stage       string `json:"stage"`
}
