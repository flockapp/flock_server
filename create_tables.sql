use flock;

DROP TABLE IF EXISTS `user`, `event`, `type`, `eventType`;

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
  `time` DATETIME,
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
