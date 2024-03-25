create table if not exists `product`(
    id int(11) not null auto_increment primary key,
    picture varchar(64) not null comment '商品图片',
    title varchar(64) not null comment '标题',
    color varchar(64) not null comment '商品颜色',
    size varchar(64) not null comment'尺寸',
    describtion varchar(64) not null comment '商品描述',
    unitprice int(11) not null comment'商品单价',
    catalogue_id int(11) not null comment '分类',
    Create_At timestamp default now() comment '创建时间'
);