SELECT id, email, password, name, created_at, updated_at
FROM users
WHERE id = $1;
