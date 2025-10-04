package task

const (
	queryCreate = `INSERT INTO tasks (u_id, title, text, status, deadline, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`
	queryDelete = `DELETE FROM tasks WHERE id = $1`
	queryGet    = `SELECT id, u_id, title, text, deadline, created_at, updated_at FROM tasks WHERE id = $1`
	queryGetAll = `SELECT id, title, status FROM tasks WHERE u_id = $1`
	querySet    = `UPDATE tasks SET title = $2, text = $3, status = $4, deadline = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1`
)
