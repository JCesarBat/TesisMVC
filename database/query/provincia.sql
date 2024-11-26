-- name: GetProvincia :one
SELECT * FROM "provincia" WHERE id =$1;

-- name: GetAllProv :many
SELECT * FROM "provincia" ;

-- name: GetProvByName :one
SELECT * FROM "provincia" WHERE name =$1 ;

-- name: InertarProv :one
Insert INTO "provincia" ("name") VALUES ($1) RETURNING *;