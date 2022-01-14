CREATE TABLE `gocery`.`items` (
    `id` int AUTO_INCREMENT,
    `eancode` text,
    `name` text,
    `amount_description` text,
    `min_amount` int,
    `count` int
     NOT NULL
     DEFAULT '0',
     PRIMARY KEY (id));
