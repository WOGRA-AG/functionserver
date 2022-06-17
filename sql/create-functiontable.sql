CREATE TABLE `functions` (
  `name` varchar(200) NOT NULL,
  `botid` varchar(200) NOT NULL,
  `code` longtext,
  `version` int(11) DEFAULT NULL,
  PRIMARY KEY (`name`,`botid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `appids` (
  `appid` VARCHAR(64) NOT NULL,
  `superuser` TINYINT NULL DEFAULT 0,
  `owner` VARCHAR(255) NULL,
  `contact` VARCHAR(255) NULL,
  PRIMARY KEY (`AppId`));

  CREATE TABLE `appid_bots` (
  `appid` VARCHAR(64) NOT NULL,
  `botid` VARCHAR(64) NOT NULL,
  PRIMARY KEY (`appid`, `botid`));

  ALTER TABLE `appid_bots` 
ADD CONSTRAINT `FK_APPIDS`
  FOREIGN KEY (`appid`)
  REFERENCES `appids` (`AppId`)
  ON DELETE CASCADE
  ON UPDATE NO ACTION;

  CREATE TABLE `key_value_store` (
  `botId` VARCHAR(64) NOT NULL,
  `identifier` VARCHAR(64) NOT NULL,
  `value` LONGTEXT NULL,
  PRIMARY KEY (`botId`, `key`));