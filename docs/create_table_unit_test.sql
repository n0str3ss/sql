CREATE DATABASE `local_unit_test`;

USE `local_unit_test`;

CREATE TABLE `unit_test`
(
    `id`        INT(11)  NOT NULL AUTO_INCREMENT,
    `field_one` INT(11)  NOT NULL,
    `field_two` CHAR(35) NOT NULL,
    PRIMARY KEY (`id`)
)
    COLLATE = 'latin1_swedish_ci'
    ENGINE = InnoDB
    ROW_FORMAT = DYNAMIC
;