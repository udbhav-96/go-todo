package models

import(
	"database/sql"
    "errors"
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
    s := &Tasks{}

    err := m.DB.QueryRow("SELECT id, task, created FROM tasks WHERE id = ?", id).Scan(&s.ID, &s.Task, &s.Created)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNoRecord
        } else {
            return nil, err
        }
    }

    return s, nil
}

// This will return the 10 most recently created snippets.
func (m *TaskModel) Latest() ([]*Tasks, error) {
    stmt := `SELECT id, task, created FROM tasks 
             ORDER BY id DESC LIMIT 10`

    rows, err := m.DB.Query(stmt)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    tasks := []*Tasks{}

    for rows.Next(){
        s := &Tasks{}

        err = rows.Scan(&s.ID, &s.Task, &s.Created)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks,s)
    }

    if err = rows.Err(); err != nil{
        return nil, err
    }
    return tasks, nil
}