CREATE TABLE "user" (
   username varchar(63) NOT NULL PRIMARY KEY,
   "password" bytea NOT NULL
);

CREATE TABLE task (
  username varchar(63) NOT NULL,
  task_number int NOT NULL,
  description varchar(255) NOT NULL,
  stage varchar(20) NOT NULL,
  PRIMARY KEY (username, task_number)
);