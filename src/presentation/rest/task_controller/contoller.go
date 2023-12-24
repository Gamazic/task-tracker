package task_controller

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"tracker_backend/src/application"
	taskUsecase "tracker_backend/src/application/task"
	permissionDomain "tracker_backend/src/domain/permission"
	"tracker_backend/src/domain/task"
	userDomain "tracker_backend/src/domain/user"
	"tracker_backend/src/presentation/rest/microframework"
)

type TaskController struct {
	CreateTaskFactory    AbsCreateFactory
	ChangeStageFactory   AbsChangeStageFactory
	GetOwnerTasksFactory AbsGetOwnerTasksFactory
	Logger               microframework.Logger
}

func (t TaskController) GetCollection(w http.ResponseWriter, r *http.Request) {
	credentials, err := microframework.GetCredentials(r)
	if err != nil {
		microframework.SendValidationError(w, err)
		return
	}
	ctx := r.Context()
	factoryDeps := CredentialCtxDeps{
		Ctx:      ctx,
		Username: credentials.Username,
		Password: credentials.Password,
	}
	getOwnerTasksUsecase, err := t.GetOwnerTasksFactory.Build(factoryDeps)
	if err != nil {
		t.Logger.Errorf("tasks get building change stage: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	queryParams := taskUsecase.OwnerTasksQuery{
		OwnerUsername: credentials.Username,
	}
	tasks, err := getOwnerTasksUsecase.Execute(queryParams)
	if errors.Is(err, application.ErrProvidingId) {
		microframework.SendForbidden(w)
		return
	}
	if errors.Is(err, userDomain.ErrInvalidUsername) {
		microframework.SendValidationError(w, err)
		return
	}
	if errors.Is(err, permissionDomain.ErrOpNotAllowed) {
		microframework.SendForbidden(w)
		return
	}
	if err != nil {
		t.Logger.Errorf("tasks get usecase call: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	responseBody := make([]TaskModel, len(tasks))
	for i := range tasks {
		responseBody[i] = TaskModel{
			TaskId:      tasks[i].TaskNumber,
			Description: tasks[i].Description,
			Stage:       tasks[i].Stage,
		}
	}
	t.Logger.LogIfErr(microframework.NewResponseBuilder(w).
		BuildStatus(http.StatusOK).
		BuildBody(responseBody).
		Send())
}

func (t TaskController) Post(w http.ResponseWriter, r *http.Request) {
	credentials, err := microframework.GetCredentials(r)
	if err != nil {
		microframework.SendValidationError(w, err)
		return
	}
	var body TaskPostRequestModel
	err = microframework.ReadValidate(r.Body, &body)
	if errors.Is(err, io.EOF) {
		microframework.SendValidationError(w, errors.New("empty body"))
		return
	}
	if errors.Is(err, microframework.ValidationErr) {
		microframework.SendValidationError(w, err)
		return
	}
	if err != nil {
		t.Logger.Errorf("task post parsing: %s", err)
		microframework.SendValidationError(w, errors.New("bad body"))
		return
	}
	ctx := r.Context()
	factoryDeps := CredentialCtxDeps{
		Ctx:      ctx,
		Username: credentials.Username,
		Password: credentials.Password,
	}
	createUsecase, err := t.CreateTaskFactory.Build(factoryDeps)
	if err != nil {
		t.Logger.Errorf("task post building: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	taskDto := taskUsecase.TaskInCreate{
		Description:   body.Description,
		OwnerUsername: credentials.Username,
	}
	createdTask, err := createUsecase.Execute(taskDto)
	if errors.Is(err, application.ErrProvidingId) {
		microframework.SendForbidden(w)
		return
	}
	if errors.Is(err, task.ErrInvalidDescription) || errors.Is(err, userDomain.ErrInvalidUsername) {
		microframework.SendValidationError(w, err)
		return
	}
	if errors.Is(err, application.ErrProvidingId) {
		microframework.SendForbidden(w)
		return
	}
	if err != nil {
		t.Logger.Errorf("task post usecase call: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	responseBody := TaskModel{
		TaskId:      createdTask.TaskNumber,
		Description: taskDto.Description,
		Stage:       createdTask.Stage,
	}
	t.Logger.LogIfErr(microframework.NewResponseBuilder(w).
		BuildStatus(http.StatusCreated).
		BuildBody(responseBody).
		Send())
}

func (t TaskController) Patch(w http.ResponseWriter, r *http.Request, taskId int) {
	credentials, err := microframework.GetCredentials(r)
	if err != nil {
		microframework.SendValidationError(w, err)
		return
	}
	var body TaskPatchRequestModel
	err = microframework.ReadValidate(r.Body, &body)
	if errors.Is(err, io.EOF) {
		microframework.SendValidationError(w, errors.New("empty body"))
		return
	}
	if errors.Is(err, microframework.ValidationErr) {
		microframework.SendValidationError(w, err)
		return
	}
	if err != nil {
		t.Logger.Errorf("task post parsing: %s", err)
		microframework.SendValidationError(w, fmt.Errorf("bad body"))
		return
	}
	ctx := r.Context()
	factoryDeps := CredentialCtxDeps{
		Ctx:      ctx,
		Username: credentials.Username,
		Password: credentials.Password,
	}
	changeStageUsecase, err := t.ChangeStageFactory.Build(factoryDeps)
	if err != nil {
		t.Logger.Errorf("task patch building change stage: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	taskDto := taskUsecase.TaskInStageChange{
		TaskNumber:    taskId,
		TargetStage:   body.Stage,
		OwnerUsername: credentials.Username,
	}
	err = changeStageUsecase.Execute(taskDto)
	if errors.Is(err, task.ErrInvalidStage) {
		microframework.SendValidationError(w, err)
		return
	}
	if errors.Is(err, application.ErrProvidingId) {
		microframework.SendForbidden(w)
		return
	}
	if errors.Is(err, permissionDomain.ErrOpNotAllowed) {
		microframework.SendForbidden(w)
		return
	}
	if errors.Is(err, taskUsecase.ErrTaskDoesntExist) {
		t.Logger.LogIfErr(microframework.NewResponseBuilder(w).
			BuildStatus(http.StatusNotFound).
			BuildBodyNestedMsg("task with specified id doesnt find").
			Send())
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
