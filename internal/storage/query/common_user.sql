-- name: GetCommonUser :one
SELECT *
FROM common_user
WHERE id = $1
LIMIT 1;

-- name: GetCommonUserByEmail :one
SELECT *
FROM common_user
WHERE email = $1
LIMIT 1;

-- name: CreateCommonUser :one
INSERT INTO common_user (full_name, cpf, email, cnpj)
values ($1, $2, $3, $4)
returning *;
