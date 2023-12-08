package permission

type PermissionService struct{}

func (p PermissionService) CanRead(user UserRoleParams, task TaskOwnershipParams) bool {
	return task.TaskOwnerUsername == user.Username
}

func (p PermissionService) UserRoleToTask(user UserRoleParams) TaskOwnershipParams {
	return TaskOwnershipParams{
		TaskOwnerUsername: user.Username,
	}
}
