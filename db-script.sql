create table profiles(
    id int(10) not null auto_increment,
    name varchar(200) not null,
    email varchar(200) default null,
    age int default 0,
    primary key (id)
);