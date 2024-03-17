create table if not exists `cart`(
    id int(11) not null auto_increment primary key,
    productid int(64) not null comment '商品id',
    quantity int(11) not null comment'商品数量',
    userid int(11) not null comment'用户id'
);