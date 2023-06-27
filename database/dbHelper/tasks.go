package dbHelper

import (
	"aytodo/models"
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateTask(db sqlx.Ext, id, title, description string, completed bool) error {
	SQL := `INSERT INTO todo(ID, Title, Description, Completed) VALUES ($1,$2,$3,$4)`
	_, err := db.Query(SQL, id, title, description, completed)
	log.Println("Added task.")
	if err != nil {
		return err
	}
	return nil
}
func AllTasks(db sqlx.Ext) ([]models.Task, error) {
	SQL := `SELECT id, title, description, completed from todo`
	var tasks []models.Task
	rows, err := db.Query(SQL)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func OrderedTasks(db sqlx.Ext) ([]models.Task, error) {
	SQL := `SELECT * FROM todo ORDER BY created_at ASC;`
	var tasks []models.Task
	rows, err := db.Query(SQL)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func OrderedTasksDue(db sqlx.Ext) ([]models.Task, error) {
	SQL := `SELECT * FROM todo ORDER BY due_date DESC;`
	var tasks []models.Task
	rows, err := db.Query(SQL)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func CompletedTasks(db sqlx.Ext) ([]models.Task, error) {
	SQL := `SELECT * FROM todo WHERE completed = true;`
	var tasks []models.Task
	rows, err := db.Query(SQL)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func GetTaskById(db sqlx.Ext, id string) (models.Task, error) {
	SQL := `SELECT * FROM tasks WHERE completed = true;`
	rows, err := db.Query(SQL, id)
	if err != nil {
		return models.Task{}, err
	}
	var task models.Task
	rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	return task, nil
}
func UpdateTask(db sqlx.Ext, id string) {
	SQL := `UPDATE todo SET Completed = true WHERE ID = $1`
	_, err := db.Query(SQL, id)
	if err != nil {
		return
	}
}
