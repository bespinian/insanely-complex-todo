package models

type Task struct {
	Id       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Complete bool   `json:"complete" bson:"complete"`
}

func (t *Task) GetId() string {
	return t.Id
}
