package token

const (
	queryAdd = "INSERT INTO tokens (u_id, token, created_at, updated_at) VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	queryGet = "SELECT * FROM tokens WHERE token = $1"
)
