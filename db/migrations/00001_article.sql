-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE article (
  id int NOT NULL AUTO_INCREMENT,
  title varchar(255),
  content varchar(255),
  created_at datetime,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE article;
