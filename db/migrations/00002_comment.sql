-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE comment (
  id int NOT NULL AUTO_INCREMENT,
  text varchar(255),
  article_id int,
  created_at datetime,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE comment;
