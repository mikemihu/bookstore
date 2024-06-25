-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users"
(
    id         UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    email      TEXT UNIQUE                            NOT NULL,
    password   TEXT                                   NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- seeder
INSERT INTO users (id, email, password)
VALUES ('3ceb2c5b-d459-4fdc-a871-bcf9d6dc61cb', 'admin@app.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('18aba2b3-ec74-447d-a9ab-83b8cd6d2b24', 'andre@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('c68964b7-a97e-4d33-b180-c3a6d0f52c09', 'budi@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('87d847d5-a7c5-46ea-9f28-5e3400de4fe8', 'cynthia@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('5be97d24-d8bb-433a-ade6-5e334dd70d20', 'debby@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('a177469d-e980-41d2-8e05-2c877d4e1abd', 'elon@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('f5cabe66-c5f8-43de-aeee-b3e527ba2843', 'fina@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('6b95e7dc-6484-48d1-be44-412d7eae7b08', 'george@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('a6b3231d-f030-40bc-8e95-5a9f0831eaf6', 'hannah@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('0679b780-bdcb-4fb0-992a-7ffe96a532b1', 'isabel@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('773f1ee7-57ae-48c7-babf-082d62cb63dd', 'jason@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('dd37d63c-ac20-41ec-90a8-73ad6d15321f', 'kevin@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('52ab4454-fb3c-4189-b9f4-677cd2929b4c', 'leonard@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.'),
       ('3ef82d2a-1d01-48b4-b7f9-68c0b72fb8c7', 'mike@mail.com',
        '$2a$10$W0HLGiNo8/PpX8z.jWhJ/uQFEACgixvWIrPE8VPOr31o7cYwT8C4.');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users" CASCADE;

DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
