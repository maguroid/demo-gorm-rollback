-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `foos` (
        id int unsigned NOT NULL AUTO_INCREMENT,
        description varchar(255) NOT NULL,
        PRIMARY KEY (id)
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `foo`;
-- +goose StatementEnd
