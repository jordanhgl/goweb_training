CREATE TABLE users (
    id int AUTO_INCREMENT COMMENT `标识`,
    name varchar(30) not null default "",
    city_id int COMMENT '所在城市id',

);
