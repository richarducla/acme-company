CREATE TABLE "users" (
    id         serial primary key,
    username   varchar(20) not null unique,
    email      varchar(50) not null unique,
    password   varchar(100) not null,
    status     bool not null default true,
    created_at timestamptz not null default (now()),
    updated_at timestamptz not null default (now())
);