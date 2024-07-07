CREATE TABLE IF NOT EXISTS portfolio_media
(
    `id`           BIGINT AUTO_INCREMENT,
    `portfolio_id` BIGINT                  NOT NULL,
    `media_url`    VARCHAR(255)            NOT NULL,
    `media_type`   ENUM ('photo', 'video') NOT NULL,
    `description`  TEXT,
    `createdAt`    TIMESTAMP               NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    FOREIGN KEY (portfolio_id) REFERENCES portfolios (id)
);
