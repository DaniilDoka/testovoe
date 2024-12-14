create extension if not exists "uuid-ossp";

create table user_status (
    id serial primary key,
    status TEXT
);

insert into user_status(status) values('registered');
insert into user_status(status) values('awaiting_confirmation');

create table "user" (
    id UUID PRIMARY KEY default uuid_generate_v4(),
    email TEXT not null,
    nickname TEXT not null,
    password TEXT not null,
    status int references user_status(id) default 2
);
