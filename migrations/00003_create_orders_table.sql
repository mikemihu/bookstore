-- +goose Up
-- +goose StatementBegin
CREATE TABLE "orders"
(
    id          UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    user_id     UUID                                   NOT NULL REFERENCES users (id),
    total_qty   INTEGER                                NOT NULL,
    total_price NUMERIC                                NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- seeder (dummy data)
INSERT INTO orders (id, user_id, total_qty, total_price)
VALUES ('9ed9b22b-ea2f-4c9e-ad67-7344c72a38ca', '18aba2b3-ec74-447d-a9ab-83b8cd6d2b24', 3, 850000),
       ('524d41bf-dfbe-44dc-a8a9-7f1c74788ad0', '18aba2b3-ec74-447d-a9ab-83b8cd6d2b24', 1, 400000),
       ('833ae9a5-2482-43d1-b737-c08d21bdef90', 'c68964b7-a97e-4d33-b180-c3a6d0f52c09', 2, 465000),
       ('ea8edfc4-1485-461b-802b-ba080b5b4902', '87d847d5-a7c5-46ea-9f28-5e3400de4fe8', 1, 190000),
       ('864ccbc3-07d4-41d6-bb5a-31201231d29e', '87d847d5-a7c5-46ea-9f28-5e3400de4fe8', 1, 260000),
       ('1c5d77e1-4b0c-4355-81c4-02781d5de85f', '87d847d5-a7c5-46ea-9f28-5e3400de4fe8', 2, 690000),
       ('901841d8-c2d1-4aec-8200-43d9f3e1fb2b', 'a177469d-e980-41d2-8e05-2c877d4e1abd', 10, 3600000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "orders" CASCADE;
-- +goose StatementEnd
