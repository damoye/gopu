CREATE TABLE Task (
    ID bigint AUTO_INCREMENT PRIMARY KEY,
    Message varchar(255)
);

CREATE TABLE Subtask (
    ID bigint AUTO_INCREMENT PRIMARY KEY,
    Client string NOT NULL INDEX,
    TaskID bigint NOT NULL INDEX,
    TTL timestamp NOT NULL INDEX,
    CreatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP INDEX,
    DeliveredAt timestamp NOT NULL DEFAULT 0 INDEX,
);
