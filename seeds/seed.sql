SET FOREIGN_KEY_CHECKS = 0;

INSERT INTO `shinkan`.`circles`
    (`name`, `about`, `catch_copy`, `description`, `circle_category_id`, `email`, `twitter`, `url`)
    VALUES
    ('airou', 'プログラミング学習、教育', 'ここに格好いいキャッチコピー入る', '各にメンターを用意しプログラミング学習を進めていきます。このサークルでは webをメインにデザインからサーバーサイドまで幅広く学ぶことができます', '3', 'airou@sample.com', 'airou_tus', 'https://airou.work'),
    ('金町食文化研究会', 'ゆる ～く食べ歩く', 'ここに格好いいキャッチコピー入る', 'モテる男は美味しい店を知っている。モテる女は美味しいものを食べている。理科大には食文化研究会がある。さぁ、君たちはどうする？', '3', 'shokubunka2020@gmail.com', 'hungryresearch', '');

INSERT INTO `shinkan`.`circles_circle_types`
    (`circle_type_id`, `circle_id`)
    VALUES
    ('1', '1'),
    ('4', '1'),
    ('9', '1'),
    ('1', '2'),
    ('8', '2');
