package user

const (
	queryCreate = `INSERT INTO users (tg_id, name, created_at, updated_at) VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`
	queryDelete = `DELETE FROM users WHERE id = $1`
	queryGet    = `SELECT id, tg_id, name, created_at, updated_at FROM users WHERE id = $1 `
	querySet    = `UPDATE users SET name = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`
)
