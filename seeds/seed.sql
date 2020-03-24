SET FOREIGN_KEY_CHECKS = 0;

INSERT INTO `shinkan`.`circles`
    (`name`, `about`, `catch_copy`, `description`, `circle_category_id`, `email`, `twitter`, `url`, `eyecatch`, `updated_at`)
    VALUES
    ('airou', 'プログラミング学習、教育', 'ここに格好いいキャッチコピー入る', '各にメンターを用意しプログラミング学習を進めていきます。このサークルでは webをメインにデザインからサーバーサイドまで幅広く学ぶことができます', '3', 'airou@sample.com', 'airou_tus', 'https://airou.work', 'https://shinkan-circle-images.s3-ap-northeast-1.amazonaws.com/airou/eyecatch/airou_eyecatch.png', NOW()),
    ('金町食文化研究会', 'ゆる ～く食べ歩く', 'ここに格好いいキャッチコピー入る', 'モテる男は美味しい店を知っている。モテる女は美味しいものを食べている。理科大には食文化研究会がある。さぁ、君たちはどうする？', '3', 'shokubunka2020@gmail.com', 'hungryresearch', '', '', '1999-01-01 00:00:00');

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
