package models

type Task struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func (t *Task) GetId() string {
	return t.Id
}
