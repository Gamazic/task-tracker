package task_controller

import (
	"errors"
	"fmt"
	"net/http"
	"tracker_backend/src/application"
	"tracker_backend/src/application/task/task_command"
	"tracker_backend/src/application/task/task_query"
	"tracker_backend/src/domain"
	taskFactory "tracker_backend/src/factory/task"
	"tracker_backend/src/infrastructure"
	"tracker_backend/src/presentation/rest/microframework"
)

type TaskHandler struct {
	CreateTaskFactory    taskFactory.AbsCreateFactory
	ChangeStageFactory   taskFactory.AbsChangeStageFactory
	GetOwnerTasksFactory taskFactory.AbsGetOwnerTasksFactory
	Logger               infrastructure.Logger
}

func (t TaskHandler) GetCollection(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get(usernameHeaderKey)
	limit := ParseIntWithDefault(r.URL.Query().Get("limit"), 0)
	offset := ParseIntWithDefault(r.URL.Query().Get("offset"), 0)

	err := UsernameHeader(username).Validate()
	if err != nil {
		microframework.SendValidationError(w, err)
		return
	}
	ctx := r.Context()
	factoryDeps := taskFactory.GetOwnerTasksDeps{
		Ctx:               ctx,
		RequesterUsername: username,
	}
	getOwnerTasksUsecase, err := t.GetOwnerTasksFactory.Build(factoryDeps)
	if err != nil {
		t.Logger.Errorf("task patch building change stage: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	queryParams := task_query.OwnerTasksQuery{
		OwnerUsername: username,
		Pagination: application.Pagination{
			Limit:  limit,
			Offset: offset,
		},
	}
	tasks, err := getOwnerTasksUsecase.Execute(queryParams)
	if errors.Is(err, domain.ErrInvalidUsername) {
		microframework.SendValidationError(w, err)
		return
	}
	if errors.Is(err, domain.ErrOpNotAllowed) {
		microframework.SendForbidden(w)
		return
	}
	if err != nil {
		t.Logger.Errorf("task patch usecase call: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	responseBody := make([]TaskModel, len(tasks.Tasks))
	for i := range tasks.Tasks {
		responseBody[i] = TaskModel{
			TaskId:      tasks.Tasks[i].TaskId,
			Description: tasks.Tasks[i].Description,
			Stage:       tasks.Tasks[i].Stage,
		}
	}
	t.Logger.LogIfErr(microframework.NewResponseBuilder(w).
		BuildStatus(http.StatusOK).
		BuildBody(responseBody).
		Send())
}

func (t TaskHandler) Post(w http.ResponseWriter, r *http.Request) {
	var body TaskPostRequestModel
	err := microframework.ReadValidate(r.Body, &body)
	if errors.Is(err, microframework.ValidationErr) {
		microframework.SendValidationError(w, err)
		return
	} else if err != nil {
		t.Logger.Errorf("task post parsing: %s", err)
		microframework.SendValidationError(w, errors.New("bad body"))
		return
	}
	username := r.Header.Get(usernameHeaderKey)
	err = UsernameHeader(username).Validate()
	if err != nil {
		microframework.SendValidationError(w, err)
		return
	}
	ctx := r.Context()
	createUsecase, err := t.CreateTaskFactory.Build(taskFactory.CreateDeps{Ctx: ctx})
	if err != nil {
		t.Logger.Errorf("task post building: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	task := task_command.TaskInCreate{
		Description:   body.Description,
		OwnerUsername: username,
	}
	createdTask, err := createUsecase.Execute(task)
	if errors.Is(err, domain.ErrInvalidDescription) || errors.Is(err, domain.ErrInvalidUsername) {
		microframework.SendValidationError(w, err)
		return
	} else if err != nil {
		t.Logger.Errorf("task post usecase call: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	responseBody := TaskModel{
		TaskId:      int(createdTask.TaskId),
		Description: task.Description,
		Stage:       string(createdTask.Stage),
	}
	t.Logger.LogIfErr(microframework.NewResponseBuilder(w).
		BuildStatus(http.StatusCreated).
		BuildBody(responseBody).
		Send())
}

func (t TaskHandler) Patch(w http.ResponseWriter, r *http.Request, taskId int) {
	var body TaskPatchRequestModel
	err := microframework.ReadValidate(r.Body, &body)
	if errors.Is(err, microframework.ValidationErr) {
		microframework.SendValidationError(w, err)
		return
	}
	if err != nil {
		t.Logger.Errorf("task post parsing: %s", err)
		microframework.SendValidationError(w, fmt.Errorf("bad body"))
		return
	}
	username := r.Header.Get(usernameHeaderKey)
	err = UsernameHeader(username).Validate()
	if err != nil {
		microframework.SendValidationError(w, err)
		return
	}
	ctx := r.Context()
	factoryDeps := taskFactory.ChangeStageDeps{
		Ctx:               ctx,
		RequesterUsername: username,
	}
	changeStageUsecase, err := t.ChangeStageFactory.Build(factoryDeps)
	if err != nil {
		t.Logger.Errorf("task patch building change stage: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	task := task_command.TaskInStageChange{
		TaskId:      taskId,
		TargetStage: body.Stage,
	}
	err = changeStageUsecase.Execute(task)
	if errors.Is(err, domain.ErrInvalidStage) {
		microframework.SendValidationError(w, err)
		return
	}
	if errors.Is(err, domain.ErrOpNotAllowed) {
		microframework.SendForbidden(w)
		return
	}
	if errors.Is(err, task_command.ErrTaskNotFound) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		t.Logger.Errorf("task patch usecase call: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	responseBody := TaskIdModel{
		TaskId: taskId,
	}
	t.Logger.LogIfErr(microframework.NewResponseBuilder(w).
		BuildStatus(http.StatusOK).
		BuildBody(responseBody).
		Send())
}
