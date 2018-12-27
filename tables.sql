CREATE TABLE todos (
    id int(10) primary key auto_increment,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    title varchar(255),
    completed int(1) default 0
);