-- name: CreateUser :one
insert into users (id, name, age, password) values ($1, $2, $3, $4) returning *;

-- name: GetUserById :one
select * from users where id=$1;


-- name: GetUserByName :one
select * from users where name=$1;
