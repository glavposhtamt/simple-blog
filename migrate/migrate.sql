CREATE DATABASE IF NOT EXISTS `go-cms`;

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `auth_key` varchar(255) NOT NULL,
  `access_token` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX ixPass ( `password` ),
  INDEX ixKey ( `auth_key` ),
  INDEX ixToken ( `access_token` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `post` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `post` text NOT NULL,
  `date` DATETIME NOT NULL,
  `author` int(10) unsigned NOT NULL,
  `thumbnail` varchar(255) NOT NULL,
   `categories` int (10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`),
  INDEX ixDate(`date`),
  FOREIGN KEY fk_cat(`categories`) REFERENCES categories(`id`),
  FOREIGN KEY fk_user(`author`) REFERENCES users(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
