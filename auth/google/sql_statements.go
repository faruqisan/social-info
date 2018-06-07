package google

var (
	insertAccessToken = "INSERT INTO access_tokens (access_token, refresh_token, token_type, expiry) VALUES(?,?,?,?)"
	getAccessToken    = "SELECT * FROM access_tokens LIMIT 1"
)
