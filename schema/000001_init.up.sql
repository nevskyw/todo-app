

CREATE TABLE lists_items /*связь многие ко многим с таблицей - users_lists "используется для сохранения соответствия всех пользователей и их списков, а также списков и их элементов"*/
(
    id serial noy null unique,
    item_id int references todo_items (id) on delete cascade not null,
    lists_id int references todo_items (id) on delete cascade not null
);


CREATE TABLE users_lists /*связь многие ко многим с таблицей - lists_items "используется для сохранения соответствия всех пользователей и их списков, а также списков и их элементов"*/
(
    id serial noy null unique,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null

);




CREATE TABLE todo_lists
(
    id serial noy null unique,
    title varchar(255) not null,
    description varchar(255)
);


CREATE TABLE users
(
    id serial noy null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null  /*password hash - для надежности*/
);


CREATE TABLE todo_items
(
    id serial noy null unique,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);