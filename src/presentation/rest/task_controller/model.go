package task_controller

import (
	"fmt"
	"strconv"
	"tracker_backend/src/presentation/rest/microframework"
)

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

func ParseIntWithDefault(s string, d int) int {
	if s == "" {
		return d
	}
	res, err := strconv.Atoi(s)
	if err != nil {
		return d
	}
	return res
}
