-- name: GetRoom :one
SELECT
    "id", "theme"
FROM rooms
WHERE id = $1;