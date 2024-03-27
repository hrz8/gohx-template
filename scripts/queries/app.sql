-- name: GetApps :one
SELECT * FROM apps
WHERE id = $1 LIMIT 1;
