package inmemory

import (
	"sync"
	"tracker_backend/src/application/task/task_command"
	"tracker_backend/src/application/task/task_query"
	userUsecase "tracker_backend/src/application/user"
	"tracker_backend/src/domain"
)

type taskInDb struct {
	TaskId      int
	Username    string
	Description string
	Stage       string
}

type Db struct {
	userTasks map[string][]int
	tasks     []taskInDb
	rwLock    sync.RWMutex
}

func NewDb() *Db {
	return &Db{
		userTasks: make(map[string][]int),
		tasks:     make([]taskInDb, 0),
	}
}

func (d *Db) SaveIncrId(task domain.Task) (domain.TaskId, error) {
	username := string(task.GetOwnerUsername())
	d.rwLock.Lock()
	nextId := len(d.tasks)
	d.tasks = append(d.tasks, taskInDb{
		TaskId:      nextId,
		Username:    username,
		Description: string(task.GetDescription()),
		Stage:       string(task.GetStage()),
	})
	d.rwLock.Unlock()

	d.saveIfNotExist(username)

	d.rwLock.Lock()
	d.userTasks[username] = append(d.userTasks[username], nextId)
	d.rwLock.Unlock()
	return domain.TaskId(nextId), nil
}

func (d *Db) ChangeStage(taskId domain.TaskId, stage domain.Stage,
	ensureTaskOwnership domain.TaskOwnershipParams) error {
	d.rwLock.Lock()
	defer d.rwLock.Unlock()

	if int(taskId) >= len(d.tasks) {
		return task_command.ErrTaskNotFound
	}
	task := &d.tasks[int(taskId)]
	if task.Username != string(ensureTaskOwnership.TaskOwnerUsername) {
		return domain.ErrOpNotAllowed
	}
	task.Stage = string(stage)
	return nil
}

func (d *Db) FetchOwnerTasks(query task_query.OwnerTasksQuery) ([]task_query.TaskResult, error) {
	tasks := d.fetchStoreOwnerTasks(query.OwnerUsername,
		query.Pagination.Limit, query.Pagination.Offset)
	tasksResult := make([]task_query.TaskResult, len(tasks))
	for i := range tasks {
		tasksResult[i] = task_query.TaskResult{
			TaskId:      tasks[i].TaskId,
			Description: tasks[i].Description,
			Stage:       tasks[i].Stage,
		}
	}
	return tasksResult, nil
}

func (d *Db) fetchStoreOwnerTasks(username string,
	limit int, offset int) []taskInDb {
	d.rwLock.RLock()
	defer d.rwLock.RUnlock()
	taskIds, ok := d.userTasks[username]
	if !ok {
		return []taskInDb{}
	}

	offset = min(offset, len(taskIds))
	if limit != 0 && offset != 0 {
		limit = offset + limit
		limit = min(len(taskIds), limit)
		taskIds = taskIds[offset:limit]
	} else if offset != 0 {
		taskIds = taskIds[offset:]
	}

	tasks := make([]taskInDb, len(taskIds))
	for i := range taskIds {
		tasks[i] = d.tasks[taskIds[i]]
	}
	return tasks
}

func (d *Db) SaveIfNotExist(user domain.User) error {
	username := string(user.GetUsername())
	wasExist := d.saveIfNotExist(username)
	if wasExist {
		return userUsecase.ErrUserAlreadyExist
	}
	return nil
}

func (d *Db) saveIfNotExist(username string) (wasExist bool) {
	d.rwLock.Lock()
	defer d.rwLock.Unlock()
	_, ok := d.userTasks[username]
	if ok {
		return true
	}
	d.userTasks[username] = make([]int, 0)
	return false
}
