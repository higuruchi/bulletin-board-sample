DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
    `val` VARCHAR(100)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `message` WRITE;
INSERT INTO `message` VALUES ("first message");
INSERT INTO `message` VALUES ("second message");
INSERT INTO `message` VALUES ("third message");
UNLOCK TABLES;
