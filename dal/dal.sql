CREATE DATABASE gopu;

CREATE TABLE gopu.task (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    data VARCHAR(255)
);

CREATE TABLE gopu.subtask (
    task_id BIGINT NOT NULL,
    token VARCHAR(31) NOT NULL,
    is_delivered BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
