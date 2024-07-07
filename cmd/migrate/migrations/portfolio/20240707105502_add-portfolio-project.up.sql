CREATE TABLE IF NOT EXISTS portfolio_projects
(
    `id`                  BIGINT AUTO_INCREMENT,
    `portfolio_id`        BIGINT       NOT NULL,
    `project_name`        VARCHAR(255) NOT NULL,
    `project_description` TEXT,
    `project_link`        VARCHAR(255),
    `start_date`          DATE,
    `end_date`            DATE,
    `art_style`           VARCHAR(255),
    `createdAt`           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`          DATETIME              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`portfolio_id`) REFERENCES portfolios (`id`)
);
