package models

// Define todo struct
type Todo struct {
	ID          int64
	Title       string
	Description string
}

// Get all todos
func GetAllTodos(id int64) ([]Todo, error) {

	var todos []Todo

	// Query db for all todos of current user
	rows, err := DB.Query("SELECT * FROM todo WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan each row
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
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
	// Exec insert query
	result, err := DB.Exec("INSERT INTO todo (title, description, user_id) VALUES (?, ?, ?)", title, description, userId)
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
func RemoveTodo(id int64) error {
	// Exec delete query
	_, err := DB.Exec("DELETE FROM todo WHERE id = ?", id)
	return err
}

// Update a todo in db
func UpdateTodo(title string, description string, id int64) error {
	// Exec update query
	_, err := DB.Exec("UPDATE todo SET title = ?, description = ? WHERE id = ?", title, description, id)
	return err
}
