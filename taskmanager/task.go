package taskmanager

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

// Task is the structure of the task.
type Task struct {
	ID        int64     `datastore:"-" json:"ID"`
	Title     string    `datastore:"title" json:"title"`
	Status    Status    `datastore:"status" json:"status"`
	CreatedAt time.Time `datastore:"createdAt" json:"createdAt"`
}

func newTask(title string) *Task {
	return &Task{
		Title:     title,
		Status:    ToDo,
		CreatedAt: time.Now(),
	}
}

func (t Task) add() error {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		return err
	}

	newKey := datastore.IncompleteKey(os.Getenv("MY_CODE"), nil)
	_, err = client.Put(ctx, newKey, &t)
	return err
}

func getAllTask() ([]Task, error) {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		return nil, err
	}
	var t []Task

	q := datastore.NewQuery(os.Getenv("MY_CODE"))
	keys, err := client.GetAll(ctx, q, &t)
	if err != nil {
		return nil, err
	}

	for i, key := range keys {
		t[i].ID = key.ID
	}

	return t, nil
}
