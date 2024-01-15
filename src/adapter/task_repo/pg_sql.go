package task_repo

import (
	"context"
	"database/sql"
	"fmt"
	taskUsecase "tracker_backend/src/application/task"
)

type PgDbAdapter struct {
	TaskTable string

	ConnPool *sql.DB
	Ctx      context.Context
}

func (p *PgDbAdapter) SaveIncrOwnerTaskNumber(taskDto taskUsecase.TaskSaveDto) (int, error) {
	query := fmt.Sprintf(
		`INSERT INTO "%s" (username, task_number, description, stage)
			SELECT $1, (SELECT coalesce(max(task_number),-1) FROM "%s"
			WHERE username=$2)+1, $3, $4 RETURNING task_number;`,
		p.TaskTable, p.TaskTable)
	rows, err := p.ConnPool.QueryContext(p.Ctx, query,
		taskDto.OwnerUsername, taskDto.OwnerUsername,
		taskDto.Description, taskDto.Stage)
	if err != nil {
		return 0, fmt.Errorf("pg SaveIncrOwnerTaskNumber query: %w", err)
	}
	defer rows.Close()
	rows.Next()
	var newId int
	err = rows.Scan(&newId)
	if err != nil {
		return 0, fmt.Errorf("pg SaveIncrOwnerTaskNumber scan: %w", err)
	}
	return newId, nil
}

func (p *PgDbAdapter) ChangeStage(changeDto taskUsecase.ChangeStageDto,
) (taskExist bool, err error) {
	query := fmt.Sprintf(`UPDATE "%s" SET stage=$1 WHERE username=$2 AND task_number=$3;`,
		p.TaskTable)
	res, err := p.ConnPool.ExecContext(p.Ctx, query,
		changeDto.TargetStage, changeDto.TaskOwnerUsername,
		changeDto.TaskNumber)
	if err != nil {
		return false, fmt.Errorf("pg ChangeStage query: %w", err)
	}
	rowsUpd, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("pg ChangeStage rows affected: %w", err)
	}
	if rowsUpd == 0 {
		return false, nil
	}
	return true, nil
}

func (p *PgDbAdapter) FetchOwnerTasks(queryDto taskUsecase.OwnerTasksQuery) ([]taskUsecase.TaskResult, error) {
	query := fmt.Sprintf(
		`SELECT task_number, description, stage FROM "%s" WHERE username=$1;`,
		p.TaskTable)
	rows, err := p.ConnPool.QueryContext(p.Ctx, query, queryDto.OwnerUsername)
	if err != nil {
		return nil, fmt.Errorf("pg FetchOwnerTasks query: %w", err)
	}
	results := make([]taskUsecase.TaskResult, 0)
	var res taskUsecase.TaskResult
	for rows.Next() {
		err = rows.Scan(&res.TaskNumber, &res.Description, &res.Stage)
		if err != nil {
			return nil, fmt.Errorf("pg FetchOwnerTasks scan: %w", err)
		}
		results = append(results, res)
	}
	return results, nil
}
