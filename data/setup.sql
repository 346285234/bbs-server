drop table topic;
drop table post;

create table topic (
    id          serial primary key,
    name        text,
    author      text,
    intro       text,
    content     text
);
create table post (
    id          serial primary key,
    content     text,
    topic_id   integer references topic(id)
);