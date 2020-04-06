SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE SCHEMA IF NOT EXISTS `shinkan_test` DEFAULT CHARACTER SET utf8;
USE `shinkan_test`;

CREATE TABLE IF NOT EXISTS `shinkan_test`.`circle_categories` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `shinkan_test`.`circles` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `about` VARCHAR(45) NULL,
  `catch_copy` VARCHAR(45) NULL,
  `cost` VARCHAR(256) NULL,
  `location` VARCHAR(256) NULL,
  `work_time` VARCHAR(256) NULL,
  `members_number` VARCHAR(256) NULL,
  `description` VARCHAR(5000) NULL,
  `circle_category_id` INT NOT NULL,
  `email` VARCHAR(255) NULL,
  `twitter` VARCHAR(45) NULL,
  `url` VARCHAR(255) NULL,
  `eyecatch` VARCHAR(10000) NULL,
  `updated_at` TIMESTAMP NULL DEFAULT '1999-09-09 09:09:09',
  PRIMARY KEY (`id`),
  INDEX `fk_circle_circle_category_idx` (`circle_category_id` ASC),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC),
  CONSTRAINT `fk_circle_circle_category`
    FOREIGN KEY (`circle_category_id`)
    REFERENCES `shinkan_test`.`circle_categories` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `shinkan_test`.`circle_types` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `shinkan_test`.`circles_circle_types` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `circle_type_id` INT NOT NULL,
  `circle_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_circles_circle_types_circle_type1_idx` (`circle_type_id` ASC),
  INDEX `fk_circles_circle_types_circle1_idx` (`circle_id` ASC),
  CONSTRAINT `fk_circles_circle_types_circle_type1`
    FOREIGN KEY (`circle_type_id`)
    REFERENCES `shinkan_test`.`circle_types` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_circles_circle_types_circle1`
    FOREIGN KEY (`circle_id`)
    REFERENCES `shinkan_test`.`circles` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `shinkan_test`.`circle_images` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `url` VARCHAR(10000) NOT NULL,
  `circle_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_circle_image_circle1_idx` (`circle_id` ASC),
  CONSTRAINT `fk_circle_image_circle1`
    FOREIGN KEY (`circle_id`)
    REFERENCES `shinkan_test`.`circles` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;