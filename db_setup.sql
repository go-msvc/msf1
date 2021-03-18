CREATE DATABASE stock;
CREATE USER 'stock'@'localhost' IDENTIFIED BY 'P@$$w0rd';
GRANT ALL ON stock.* TO 'stock'@'localhost';

USE stock;

DROP TABLE IF EXISTS `company`;
CREATE TABLE IF NOT EXISTS `company` (
    `company_id` INT(11) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) NOT NULL,
    `info` VARCHAR(256) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `date_modified` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`company_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Rollback
-- DROP TABLE company;


DROP TABLE IF EXISTS `location`;
CREATE TABLE IF NOT EXISTS `location` (
    `location_id` INT(11) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) NOT NULL,
    `info` VARCHAR(256) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `date_modified` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX location_index (`name`),
    PRIMARY KEY (`location_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Rollback
-- DROP TABLE location;


DROP TABLE IF EXISTS `stock_status`;
CREATE TABLE IF NOT EXISTS `stock_status` (
    `stock_status_id` INT(11) NOT NULL AUTO_INCREMENT,
    `status` VARCHAR(64) NOT NULL,
    `description` VARCHAR(256) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `date_modified` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`stock_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Rollback
-- DROP TABLE stock_status;

DROP TABLE IF EXISTS `owner`;
CREATE TABLE IF NOT EXISTS `owner` (
    `owner_id` INT(11) NOT NULL AUTO_INCREMENT,
    `company_id` INT(11) NOT NULL,
    `merchant_reference` VARCHAR(64) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `date_modified` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`owner_id`),
    INDEX company_index (`company_id`),
    INDEX merchant_reference_index (`merchant_reference`),
    INDEX date_modified_index (`date_modified`),
    FOREIGN KEY(company_id) REFERENCES company(company_id),
    CONSTRAINT uniqueOwner UNIQUE (`company_id`, `merchant_reference`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Rollback
-- DROP TABLE owner;

DROP TABLE IF EXISTS `inventory`;
CREATE TABLE IF NOT EXISTS `inventory` (
    `inventory_id` INT(11) NOT NULL AUTO_INCREMENT,
    `owner_id` INT(11) NOT NULL,
    `product_reference` VARCHAR(64) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `date_modified` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`inventory_id`),
    INDEX company_id_index (`owner_id`),
    INDEX product_reference_index (`product_reference`),
    INDEX date_modified_index (`date_modified`),
    FOREIGN KEY(owner_id) REFERENCES owner(owner_id),
    CONSTRAINT uniqueInventory UNIQUE (`product_reference`, `owner_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Rollback
-- DROP TABLE inventory;


DROP TABLE IF EXISTS `stock`;
CREATE TABLE IF NOT EXISTS `stock` (
    `stock_id` INT(11) NOT NULL AUTO_INCREMENT,
    `inventory_id` INT(11) NOT NULL,
    `location_id` INT(11) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `date_modified` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`stock_id`),
    INDEX inventory_index (`inventory_id`),
    INDEX location_index (`location_id`),
    INDEX date_modified_index (`date_modified`),
    FOREIGN KEY(inventory_id) REFERENCES inventory(inventory_id),
    FOREIGN KEY(location_id) REFERENCES location(location_id),
    CONSTRAINT uniqueStock UNIQUE (`inventory_id`, `location_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Rollback
-- DROP TABLE stock;


DROP TABLE IF EXISTS `stock_level`;
CREATE TABLE IF NOT EXISTS `stock_level` (
    `stock_id` INT(11) NOT NULL,
    `stock_status_id` INT(11) NOT NULL,
    `quantity` INT(11) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `date_modified` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`stock_id`, `stock_status_id`),
    INDEX date_modified_index (`date_modified`),
    FOREIGN KEY(stock_status_id) REFERENCES stock_status(stock_status_id),
    FOREIGN KEY(stock_id) REFERENCES stock(stock_id),
    CONSTRAINT uniqueStockLevel UNIQUE (`stock_id`, `stock_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DELETE FROM `location`;
INSERT INTO `location` SET location_id=1, name="CPT", info="go-msvc Cape Town DC";
INSERT INTO `location` SET location_id=3, name="JHB", info="go-msvc Johannesburg DC";

DELETE FROM `company`;
INSERT INTO `company` SET company_id=1, name="go-msvc", info="go-msvc";

DELETE FROM `owner`;
INSERT INTO `owner` SET owner_id=1, company_id=1, merchant_reference="go-msvc";

