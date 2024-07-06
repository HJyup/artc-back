CREATE TABLE IF NOT EXISTS users
(
    `pk`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `id`          VARCHAR(255) NOT NULL,
    `avatar`      VARCHAR(255),
    `first_name`  VARCHAR(255) NOT NULL,
    `last_name`   VARCHAR(255) NOT NULL,
    `email`       VARCHAR(255) NOT NULL,
    `password`    VARCHAR(255) NOT NULL,
    `speciality_id` INT UNSIGNED NOT NULL,
    `location`    VARCHAR(255) NOT NULL,
    `is_accepted` BOOLEAN NOT NULL DEFAULT FALSE,
    `is_reviewer` BOOLEAN NOT NULL DEFAULT FALSE,
    `createdAt`   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`pk`),
    UNIQUE KEY (`email`),
    FOREIGN KEY (`speciality_id`) REFERENCES specialities(`id`)
);
