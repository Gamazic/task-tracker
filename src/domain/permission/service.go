package permission

type PermissionService struct{}

func (p PermissionService) HaveAccess(user UserRoleIdentity, task TaskOwnershipParams) bool {
	return task.TaskOwnerUsername == user.Username
}

func (p PermissionService) UserRoleToTask(user UserRoleIdentity) TaskOwnershipParams {
	return TaskOwnershipParams{
		TaskOwnerUsername: user.Username,
	}
}
