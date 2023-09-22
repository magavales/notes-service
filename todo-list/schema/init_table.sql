CREATE TABLE tasks (
    id bigserial primary key not null,
    header varchar(64) not null,
    description text not null,
    date timestamp not null,
    status varchar(16) not null
)