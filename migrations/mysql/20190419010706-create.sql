
-- +migrate Up

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema tamabus
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Table `buses`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `buses` (
  `id` INT NOT NULL,
  `name` VARCHAR(255) NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `images`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `images` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `path` VARCHAR(512),
  `base64` BLOB,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `bus_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `image_bus_id_idx` (`bus_id` ASC),
  CONSTRAINT `image_bus_id`
    FOREIGN KEY (`bus_id`)
    REFERENCES `buses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `congestion_log`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `congestion_log` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `latitude` DOUBLE(10,7) NOT NULL,
  `longitude` DOUBLE(10,7) NOT NULL,
  `congestion` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `bus_id` INT NOT NULL,
  `position` INT NOT NULL,
  `complete` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  INDEX `congestion_log_bus_id_idx` (`bus_id` ASC),
  CONSTRAINT `congestion_log_bus_id`
    FOREIGN KEY (`bus_id`)
    REFERENCES `buses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

-- +migrate Down
