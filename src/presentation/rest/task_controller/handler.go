package task_controller

import (
	"errors"
	"fmt"
	"net/http"
	taskUsecase "tracker_backend/src/application/task"
	permissionDomain "tracker_backend/src/domain/permission"
	"tracker_backend/src/domain/task"
	userDomain "tracker_backend/src/domain/user"
	"tracker_backend/src/factory"
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
	queryParams := taskUsecase.OwnerTasksQuery{
		OwnerUsername: username,
	}
	tasks, err := getOwnerTasksUsecase.Execute(queryParams)
	if errors.Is(err, userDomain.ErrInvalidUsername) {
		microframework.SendValidationError(w, err)
		return
	}
	if errors.Is(err, permissionDomain.ErrOpNotAllowed) {
		microframework.SendForbidden(w)
		return
	}
	if err != nil {
		t.Logger.Errorf("task patch usecase call: %s", err)
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
	createUsecase, err := t.CreateTaskFactory.Build(factory.CtxDeps{Ctx: ctx})
	if err != nil {
		t.Logger.Errorf("task post building: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	taskDto := taskUsecase.TaskInCreate{
		Description:   body.Description,
		OwnerUsername: username,
	}
	createdTask, err := createUsecase.Execute(taskDto)
	if errors.Is(err, task.ErrInvalidDescription) || errors.Is(err, userDomain.ErrInvalidUsername) {
		microframework.SendValidationError(w, err)
		return
	}
	//if errors.Is(err, userUsecase.ErrUserDoesntExist) {
	//	t.Logger.LogIfErr(microframework.NewResponseBuilder(w).
	//		BuildStatus(http.StatusBadRequest).
	//		BuildBodyNestedMsg("provided user does not exist").
	//		Send())
	//	return
	//}
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
	taskDto := taskUsecase.TaskInStageChange{
		TaskNumber:  taskId,
		TargetStage: body.Stage,
	}
	err = changeStageUsecase.Execute(taskDto)
	if errors.Is(err, task.ErrInvalidStage) {
		microframework.SendValidationError(w, err)
		return
	}
	if errors.Is(err, permissionDomain.ErrOpNotAllowed) {
		microframework.SendForbidden(w)
		return
	}
	if errors.Is(err, taskUsecase.ErrTaskDoesntExist) {
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
