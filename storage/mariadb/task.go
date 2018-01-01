package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	taskFields = `duration, title, description, reward, rewardid, cashreward, xpreward, rewardmethod, startzone, minlevel, maxlevel, repeatable`
	taskBinds  = `:duration, :title, :description, :reward, :rewardid, :cashreward, :xpreward, :rewardmethod, :startzone, :minlevel, :maxlevel, :repeatable`
	taskSets   = `duration=:duration, title=:title, description=:description, reward=:reward, rewardid=:rewardid, cashreward=:cashreward, xpreward=:xpreward, rewardmethod=:rewardmethod, startzone=:startzone, minlevel=:minlevel, maxlevel=:maxlevel, repeatable=:repeatable`
)

//GetTask will grab data from storage
func (s *Storage) GetTask(taskID int64) (task *model.Task, err error) {
	task = &model.Task{}
	err = s.db.Get(task, fmt.Sprintf("SELECT id, %s FROM tasks WHERE id = ?", taskFields), taskID)
	if err != nil {
		return
	}
	return
}

//CreateTask will grab data from storage
func (s *Storage) CreateTask(task *model.Task) (err error) {
	if task == nil {
		err = fmt.Errorf("Must provide task")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO tasks(%s)
		VALUES (%s)`, taskFields, taskBinds), task)
	if err != nil {
		return
	}
	taskID, err := result.LastInsertId()
	if err != nil {
		return
	}
	task.ID = taskID
	return
}

//ListTask will grab data from storage
func (s *Storage) ListTask() (tasks []*model.Task, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM tasks ORDER BY ID ASC`, taskFields))
	if err != nil {
		return
	}

	for rows.Next() {
		task := model.Task{}
		if err = rows.StructScan(&task); err != nil {
			return
		}
		tasks = append(tasks, &task)
	}
	return
}

//EditTask will grab data from storage
func (s *Storage) EditTask(taskID int64, task *model.Task) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE tasks SET %s WHERE id = :id`, taskSets), task)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}

//DeleteTask will grab data from storage
func (s *Storage) DeleteTask(taskID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM tasks WHERE id = ?`, taskID)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}
