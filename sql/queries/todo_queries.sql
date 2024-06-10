-- name: CreateTodo :one
insert into todos (id, title, description, user_id) values ($1, $2, $3, $4) returning *;

-- name: GetTodoById :one
select * from todos where id=$1;

-- name: UpdateTodo :one
update todos set done=true where id=$1 returning *;
