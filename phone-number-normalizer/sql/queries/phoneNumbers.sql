-- name: UpdatePhoneNumber :one
UPDATE phone_numbers
SET phone_number = $1
WHERE id = $2;
