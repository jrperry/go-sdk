package iland

import (
	"fmt"
	"time"
)

type Task struct {
	ID           string `json:"uuid"`
	Operation    string `json:"operation"`
	Description  string `json:"operation_description"`
	Type         string `json:"task_type"`
	Status       string `json:"status"`
	Progress     int    `json:"progress"`
	Active       bool   `json:"active"`
	Synced       bool   `json:"synced"`
	Message      string `json:"message"`
	UserName     string `json:"username"`
	UserFullName string `json:"user_full_name"`
	EntityID     string `json:"entity_uuid"`
	EntityName   string `json:"entity_name"`
	OrgID        string `json:"org_uuid"`
	CompanyID    string `json:"company_id"`
	LocationID   string `json:"location_id"`
	StartTime    int    `json:"start_time"`
	EndTime      int    `json:"end_time"`
}

type taskService struct {
	client *client
}

func (s *taskService) Get(taskID string) (Task, error) {
	task := Task{}
	err := s.client.getObject(fmt.Sprintf("/v1/tasks/%s", taskID), &task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) Track(taskID string) (Task, error) {
	for {
		time.Sleep(time.Second * 5)
		task, err := s.Get(taskID)
		if err != nil {
			return Task{}, err
		}
		if !task.Active && task.Synced {
			return task, nil
		}
	}
}
