CREATE DATABASE `flock`;
use `flock`;

DROP TABLE IF EXISTS `user`, `event`, `type`, `eventType`, `userEvent`;

CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255),
  `full_name` varchar(255),
  `password` varchar(255),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `event` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `host_id` bigint(20) NOT NULL REFERENCES user(`host_id`),
  `name` varchar(255),
  `cost` tinyint(1),
  `time` bigint(20),
  `lat` float,
  `lng` float,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `eventType` (
  `eventId` bigint(20) NOT NULL REFERENCES event(`eventId`),
  `typeId` bigint(20) NOT NULL REFERENCES type(`typeId`),
  PRIMARY KEY (`eventId`, `typeId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `activityType` (
  `activityId` bigint(20) NOT NULL REFERENCES event(`activityId`),
  `typeId` bigint(20) NOT NULL REFERENCES type(`typeId`),
  PRIMARY KEY (`activityId`, `typeId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `userEvent` (
  `userId` bigint(20) NOT NULL REFERENCES user(`userId`),
  `eventId` bigint(20) NOT NULL REFERENCES event(`evenId`),
  PRIMARY KEY (`userId`, `eventId`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `activity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `eventId` bigint(20) NOT NULL REFERENCES event(`eventId`),
  `placeId` VARCHAR(255) NOT NULL,
  `rating` float,
  `lat` float,
  `lng` float,
  `cost` tinyint,
  `startTime` bigint(20),
  `endTime` bigint(20),
  `desc` varchar(512),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
