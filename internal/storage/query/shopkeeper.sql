-- name: GetShopKeeperUser :one
SELECT cu.*, su.cnpj
FROM shoopkeepers su
         JOIN public.common_user cu on cu.id = su.common_user_id
WHERE common_user_id = $1
LIMIT 1;

-- name: CreateNewShopKeeperUser :one
WITH user_sub AS ( INSERT INTO common_user (full_name, cpf, email) values ($1, $2, $3) RETURNING * ),
     shopkeeper_sub AS ( INSERT INTO shoopkeepers (common_user_id, cnpj) values ((SELECT id FROM user_sub), $4) RETURNING *)
SELECT user_sub.*, shopkeeper_sub.cnpj
from user_sub, shopkeeper_sub;

