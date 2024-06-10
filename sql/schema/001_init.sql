-- +goose Up

create table users (
    id uuid primary key,
    name text not null unique,
    age int not null,
    password text not null
);

create table todos (
  id uuid primary key,
  title text not null,
  description text not null,
  done bool default false,
  user_id uuid not null references users(id) on delete cascade,
  unique(title, user_id)
);

-- +goose Down
drop table todos;
drop table users;
