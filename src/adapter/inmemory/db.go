package inmemory

import (
	"sync"
	taskUsecase "tracker_backend/src/application/task"
)

type taskInDb struct {
	TaskId      int
	Username    string
	Description string
	Stage       string
	Deleted     bool
}

type Db struct {
	userTasks map[string][]taskInDb
	rwLock    sync.RWMutex
}

func NewDb() *Db {
	return &Db{
		userTasks: make(map[string][]taskInDb),
	}
}

func (d *Db) SaveIncrOwnerTaskNumber(taskDto taskUsecase.TaskSaveDto) (int, error) {
	d.rwLock.Lock()
	defer d.rwLock.Unlock()
	nextId := len(d.userTasks[taskDto.OwnerUsername])
	d.userTasks[taskDto.OwnerUsername] = append(d.userTasks[taskDto.OwnerUsername], taskInDb{
		TaskId:      nextId,
		Username:    taskDto.OwnerUsername,
		Description: taskDto.Description,
		Stage:       taskDto.Stage,
	})
	return nextId, nil
}

func (d *Db) ChangeStage(changeStageDto taskUsecase.ChangeStageDto) (taskExist bool, ownershipMatched bool, err error) {
	d.rwLock.Lock()
	defer d.rwLock.Unlock()
	username := changeStageDto.TaskOwnerUsername
	if changeStageDto.TaskNumber >= len(d.userTasks[username]) {
		return false, true, nil
	}
	d.userTasks[username][changeStageDto.TaskNumber].Stage = changeStageDto.TargetStage
	return true, true, nil
}

func (d *Db) FetchOwnerTasks(taskDto taskUsecase.OwnerTasksQuery) ([]taskUsecase.TaskResult, error) {
	d.rwLock.RLock()
	tasks := d.fetchStoreOwnerTasks(taskDto.OwnerUsername)
	d.rwLock.RUnlock()
	tasksResult := make([]taskUsecase.TaskResult, len(tasks))
	for i := range tasks {
		tasksResult[i] = taskUsecase.TaskResult{
			TaskNumber:  tasks[i].TaskId,
			Description: tasks[i].Description,
			Stage:       tasks[i].Stage,
		}
	}
	return tasksResult, nil
}

func (d *Db) fetchStoreOwnerTasks(username string) []taskInDb {
	tasks, ok := d.userTasks[username]
	if !ok {
		return []taskInDb{}
	}
	return tasks
}

func (d *Db) saveCheckFreeUsername(username string) (bool, error) {
	_, ok := d.userTasks[username]
	if ok {
		return false, nil
	}
	d.userTasks[username] = make([]taskInDb, 0)
	return true, nil
}
