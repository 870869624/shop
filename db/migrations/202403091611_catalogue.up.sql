create table if not exists `catalogue`(
    id int(11) not null auto_increment primary key,
    cataloguename varchar(64) not null comment '商品名称',
    icon varchar(64) not null comment '图标',
    `order` int(11) not null comment'顺序'
);