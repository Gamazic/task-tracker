package domain

import "errors"

type Task struct {
	taskId        TaskId
	description   Description
	ownerUsername Username
	stage         Stage
	isConfirmedId bool
}

func NewTask(description string, ownerUsername string) (Task, error) {
	descriptionDomain, err := NewDescription(description)
	if err != nil {
		return Task{}, err
	}
	ownerUsernameDomain, err := NewUsername(ownerUsername)
	if err != nil {
		return Task{}, err
	}
	return Task{
		taskId:        EmptyTaskId,
		description:   descriptionDomain,
		ownerUsername: ownerUsernameDomain,
		stage:         ToDo,
		isConfirmedId: false,
	}, nil
}

func (t Task) GetConfirmedId() (id TaskId, confirmed bool) {
	if !t.isConfirmedId {
		return EmptyTaskId, false
	}
	return t.taskId, true
}

func (t *Task) SetConfirmedNewId(newId TaskId) error {
	err := newId.Validate()
	if err != nil {
		return err
	}
	t.taskId = newId
	t.isConfirmedId = true
	return nil
}

func (t Task) GetStage() Stage {
	return t.stage
}

var ErrInvalidTaskId = errors.New("invalid task id")

type TaskId int

func NewTaskId(i int) (TaskId, error) {
	taskId := TaskId(i)
	err := taskId.Validate()
	if err != nil {
		return EmptyTaskId, err
	}
	return taskId, nil
}
func (t TaskId) Validate() error {
	if t < 0 {
		return ErrInvalidTaskId
	}
	return nil
}

const EmptyTaskId = TaskId(0)

var ErrInvalidStage = errors.New("invalid stage")

type Stage string

const (
	ToDo       = Stage("todo")
	InProgress = Stage("in_progress")
	Done       = Stage("done")
)

func NewStage(s string) (Stage, error) {
	switch stage := Stage(s); stage {
	case ToDo:
		return stage, nil
	case InProgress:
		return stage, nil
	case Done:
		return stage, nil
	default:
		return Stage(""), ErrInvalidStage
	}
}

var ErrInvalidDescription = errors.New("invalid description")

type Description string

func NewDescription(s string) (Description, error) {
	if s == "" {
		return Description(""), ErrInvalidDescription
	}
	return Description(s), nil
}
