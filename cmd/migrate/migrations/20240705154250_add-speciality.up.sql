CREATE TABLE IF NOT EXISTS specialities
(
    `id`    INT UNSIGNED NOT NULL,
    `name`  VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`name`)
);

INSERT INTO specialities (id, name) VALUES
                                        (1, 'musician'),
                                        (2, 'actor'),
                                        (3, 'visual_artist'),
                                        (4, 'writer'),
                                        (5, 'designer'),
                                        (6, 'dancer'),
                                        (7, 'photographer'),
                                        (8, 'filmmaker');
