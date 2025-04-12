-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages
(
    id            bigserial
        primary key,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone,
    uuid          text,
    sender_uuid   text,
    receiver_uuid text,
    content       text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages
-- +goose StatementEnd
