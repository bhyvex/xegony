package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type TaskRepository struct {
	stor storage.Storage
}

func (g *TaskRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *TaskRepository) Get(taskId int64) (task *model.Task, err error) {
	if taskId == 0 {
		err = fmt.Errorf("Invalid Task ID")
		return
	}
	task, err = g.stor.GetTask(taskId)
	return
}

func (g *TaskRepository) Create(task *model.Task) (err error) {
	if task == nil {
		err = fmt.Errorf("Empty task")
		return
	}
	schema, err := task.NewSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}
	task.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(task))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	err = g.stor.CreateTask(task)
	if err != nil {
		return
	}
	return
}

func (g *TaskRepository) Edit(taskId int64, task *model.Task) (err error) {
	schema, err := task.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(task))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}

	err = g.stor.EditTask(taskId, task)
	if err != nil {
		return
	}
	return
}

func (g *TaskRepository) Delete(taskId int64) (err error) {
	err = g.stor.DeleteTask(taskId)
	if err != nil {
		return
	}
	return
}

func (g *TaskRepository) List() (tasks []*model.Task, err error) {
	tasks, err = g.stor.ListTask()
	if err != nil {
		return
	}
	return
}