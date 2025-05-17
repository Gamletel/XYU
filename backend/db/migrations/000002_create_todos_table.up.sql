create table if not exists todos(
    id serial primary key,
    title varchar(75) not null,
    text text,
    image text,
    user_id integer references users(id) on delete cascade,
    created_at timestamptz not null default now()
)