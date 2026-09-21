UPDATE users
SET email = $2, password = $3, name = $4, updated_at = now()
WHERE id = $1
RETURNING id, email, password, name, created_at, updated_at;
