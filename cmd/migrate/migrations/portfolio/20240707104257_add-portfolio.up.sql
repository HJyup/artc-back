CREATE TABLE IF NOT EXISTS portfolios
(
    `id`            BIGINT AUTO_INCREMENT,
    `user_id`       VARCHAR(255),
    `description`   TEXT,
    `education`     VARCHAR(255),
    `skills`        TEXT      NOT NULL,
    `exhibitions`   TEXT,
    `awards`        TEXT,
    `instagram_url` VARCHAR(255),
    `spotlight_url` VARCHAR(255),
    `createdAt`     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    DATETIME           DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY (`user_id`),
    FOREIGN KEY (`user_id`) REFERENCES users (`id`)
);
