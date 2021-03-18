package model

import "github.com/go-msvc/msf/model"

type StockLevel struct {
	model.Item
	Stock       `uniq:"uniqueStockLevel"`
	StockStatus `uniq:"uniqueStockLevel"`
	Quantity    int
}

/*
CURRENT msf modeled structure:

CREATE TABLE IF NOT EXISTS `stock_level` (
	`stock_level_id` INT(11) NOT NULL AUTO_INCREMENT, 				<-- get rid of this because uniq reference combination
	`stock_id` INT(11) NOT NULL,
	`stock_status_id` INT(11) NOT NULL,
	`quantity` INT(11) NOT NULL,
	PRIMARY KEY (`stock_level_id`),									<-- change to refs: PRIMARY KEY (`stock_id`, `stock_status_id`)
	FOREIGN KEY(stock_id) REFERENCES stock(stock_id),
	FOREIGN KEY(stock_status_id) REFERENCES stock_status(stock_status_id),
	CONSTRAINT uniqueStockLevel UNIQUE (`stock_id`,`stock_status_id`),
	CONSTRAINT  UNIQUE (`quantity`)									<-- get rid of this... where from??? no name too!
) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/

/*
EXPECTED STRUCTURE:
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
*/
