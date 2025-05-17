create table if not exists users
(
    id      serial primary key,
    email   varchar(100) unique not null,
    login   varchar(40) unique not null,
    name    varchar(40),
    surname varchar(40),
    avatar  text,
    password text not null,
    created_at timestamptz not null default now()
);

create index idx_users_email on users(email);
create index idx_users_login on users(login);