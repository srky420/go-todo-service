package models

import "errors"

// Define todo struct
type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int64  `json:"user_id"`
}

// Get all todos
func GetAllTodos(id int64) ([]Todo, error) {

	var todos []Todo

	// Query db for all todos of current user
	rows, err := DB.Query("SELECT * FROM task WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan each row
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.UserId); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	// Check for error in overall query
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

// Insert a todo in db
func AddTodo(title string, description string, userId int64) (int64, error) {
	// Check if title and description are empty
	if title == "" || description == "" {
		return 0, errors.New("title and description fields cannot be empty")
	}

	// Exec insert query
	result, err := DB.Exec("INSERT INTO task (title, description, user_id) VALUES (?, ?, ?)", title, description, userId)
	if err != nil {
		return 0, err
	}

	// Get the id of created todo
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Remove a todo from db
func RemoveTodo(id int64, userId int64) error {
	// Exec delete query
	result, err := DB.Exec("DELETE FROM task WHERE id = ? AND user_id = ?", id, userId)
	if err != nil {
		return err
	}

	// Check affected row
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row != 1 {
		return errors.New("expected to affect 1 row")
	}
	return nil
}

// Update a todo in db
func UpdateTodo(title string, description string, id int64, userId int64) error {
	// Exec update query
	result, err := DB.Exec("UPDATE task SET title = ?, description = ? WHERE id = ? AND user_id = ?", title, description, id, userId)
	if err != nil {
		return err
	}

	// Check affected row
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row != 1 {
		return errors.New("expected to affect 1 row")
	}
	return nil
}
