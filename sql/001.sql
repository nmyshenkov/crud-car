create database vehicals;

CREATE TABLE `vehicals`.`cars` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `company` varchar(255) NOT NULL,
    `model` varchar(255) NOT NULL,
    `capacity` int(11) NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
