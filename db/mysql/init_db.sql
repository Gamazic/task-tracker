CREATE DATABASE tasktracker;

CREATE TABLE tasktracker.`user` (
    username varchar(63) NOT NULL,
    password binary(16) NOT NULL,
    PRIMARY KEY (username)
)
ENGINE=InnoDB;

CREATE TABLE tasktracker.task (
    task_id int(4) NOT NULL AUTO_INCREMENT,
    username varchar(63) NOT NULL,
    stage varchar(20) NOT NULL,
    PRIMARY KEY (task_id),
    FOREIGN KEY (username) REFERENCES user(username)
) ENGINE=InnoDB;
