package models

import(
	"database/sql"
	"time"
)

type Tasks struct{
	ID 			int
	Task 		string
	Created 	time.Time
}

type TaskModel struct{
	DB *sql.DB
}

// This will insert a new task into the database.
func (m *TaskModel) Insert(task string) (int, error) {
    stmt := `INSERT INTO tasks (task, created)
    VALUES(?, UTC_TIMESTAMP())`

    result, err := m.DB.Exec(stmt, task)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *TaskModel) Get(id int) (*Tasks, error) {
    return nil, nil
}

// This will return the 10 most recently created snippets.
func (m *TaskModel) Latest() ([]*Tasks, error) {
    return nil, nil
}