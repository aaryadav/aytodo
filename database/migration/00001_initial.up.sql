create TABLE if not exists todo(
    id text primary key,
    title text not null,
    description text not null,
    completed bool not null,
    created_at timestamp default current_timestamp,
    due_date timestamp not null
)