create table if not exists `users`(
    id int(11) not null auto_increment primary key,
    username varchar(64) not null comment '用户名',
    password varchar(256) not null comment '密码',
    phone varchar(64) not null comment '电话号码',
    gender int(2) not null comment'性别',
    age int(2) not null comment '年龄',
    refreshedAt timestamp not null default now() comment'创建时间'
);
