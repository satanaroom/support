CREATE TABLE support_items
(
    id     serial not null unique,
    number int not null unique,
    name   varchar(255) not null,
    date   timestamp not null
);