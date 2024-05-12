package models

import "errors"

// Get all todos
func GetAllTodosAdmin() ([]Todo, error) {
	// Query db for all todos
	var todos []Todo

	// Query db for all todos of current user
	rows, err := DB.Query("SELECT * FROM task")
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

// Get all todos
func GetAllUsersAdmin() ([]User, error) {
	// Query db for all todos
	var users []User

	// Query db for all todos of current user
	rows, err := DB.Query("SELECT * FROM user WHERE is_admin = ?", false)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan each row
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.IsAdmin); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	// Check for error in overall query
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// Remove a todo from db
func RemoveTodoAdmin(id int64) error {
	// Exec delete query
	result, err := DB.Exec("DELETE FROM task WHERE id = ?", id)
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
func UpdateTodoAdmin(title string, description string, id int64) error {
	// Exec update query
	result, err := DB.Exec("UPDATE task SET title = ?, description = ? WHERE id = ?", title, description, id)
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
