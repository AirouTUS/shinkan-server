use `shinkan_test`;

SET NAMES utf8;

INSERT INTO `shinkan_test`.`circle_types` (`name`) VALUES ('葛飾');
INSERT INTO `shinkan_test`.`circle_types` (`name`) VALUES ('神楽坂');
INSERT INTO `shinkan_test`.`circle_types` (`name`) VALUES ('野田');
INSERT INTO `shinkan_test`.`circle_types` (`name`) VALUES ('インカレ');
INSERT INTO `shinkan_test`.`circle_types` (`name`) VALUES ('機関');
INSERT INTO `shinkan_test`.`circle_types` (`name`) VALUES ('体育局');
INSERT INTO `shinkan_test`.`circle_types` (`name`) VALUES ('公認');
INSERT INTO `shinkan_test`.`circle_types` (`name`) VALUES ('届出');
INSERT INTO `shinkan_test`.`circle_types` (`name`) VALUES ('有理会');

INSERT INTO `shinkan_test`.`circle_categories` (`name`) VALUES ('委員会');
INSERT INTO `shinkan_test`.`circle_categories` (`name`) VALUES ('体育会系');
INSERT INTO `shinkan_test`.`circle_categories` (`name`) VALUES ('文化系');

INSERT INTO `shinkan`.`circles`
    (`name`, `about`, `catch_copy`, `location`, `work_time`, `members_number`, `cost`, `description`, `circle_category_id`, `email`, `twitter`, `url`, `eyecatch`, `updated_at`)
    VALUES
    ("name1", "about1", "catch_copy1", "location1", "work_time1", "members_number1", "cost1", "description1", "1", "1@example.com", "twitter1", "url1", "eyecatch1", "2010-01-01 12:00:00"),
    ("name2", "about2", "catch_copy2", "location2", "work_time2", "members_number2", "cost2", "description2", "2", "2@example.com", "twitter2", "url2", "eyecatch2", "2020-02-02 22:00:00");

INSERT INTO `shinkan`.`circles_circle_types`
    (`circle_type_id`, `circle_id`)
    VALUES
    ('1', '1'),
    ('4', '1'),
    ('9', '1'),
    ('1', '2'),
    ('8', '2');

INSERT INTO `shinkan`.`circle_images`
    (`url`, `circle_id`)
    VALUES
    ('https://hoge.com/hoge1.png', '1'),
    ('https://hoge.com/hoge2.png', '1'),
    ('https://hoge.com/hoge3.png', '2');