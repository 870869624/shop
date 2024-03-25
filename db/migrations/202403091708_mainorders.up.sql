create table if not exists `maindrders`(
    id int(11) not null auto_increment primary key,
    date varchar(64) not null comment '创建日期',
    ordernumber varchar(64) not null comment '订单号',
    totalprice int(64) not null comment '商品实付款', 
    userid int(11) not null comment'用户id',
    Create_At timestamp default now() comment '创建时间'
);
