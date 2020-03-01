create table `account` (
    `id` bigint unsigned auto_increment,
    `account_name` varchar(255) not null,
    `password` char(32) not null,
    `email` varchar(255) not null,
    primary key(`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='账户信息';