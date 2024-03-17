create table if not exists orderforms(
    id int(11) not null auto_increment primary key,
    mainorder_id int(11) not null comment'所属订单id',
    picture varchar(64) not null comment '商品图片',
    title varchar(64) not null comment '标题',
    color varchar(64) not null comment '商品颜色',
    size varchar(64) not null comment'尺寸',
    unitprice int(11) not null comment'商品单价',
    quantity int(11) not null comment '商品数量'
);