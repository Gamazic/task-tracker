package task

import (
	"tracker_backend/internal/domain"
	"tracker_backend/internal/domain/user"
)

type Task struct {
	OwnerTaskId   TaskNumber
	Description   Description
	OwnerUsername user.Username
	Stage         Stage
}

func (t Task) Validate() error {
	return domain.ValidateCascade([]domain.Validatable{
		t.OwnerTaskId,
		t.Description,
		t.OwnerUsername,
		t.Stage,
	})
}
