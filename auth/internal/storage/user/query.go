package user

const (
	queryAdd = "INSERT INTO users (tg_id, login, password, created_at, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	queryGet = "SELECT * FROM users WHERE login = $1"
)
