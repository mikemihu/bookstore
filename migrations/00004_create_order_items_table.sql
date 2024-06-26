-- +goose Up
-- +goose StatementBegin
CREATE TABLE "order_items"
(
    id         UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    order_id   UUID                                   NOT NULL REFERENCES orders (id),
    book_id    UUID                                   NOT NULL REFERENCES books (id),
    qty        INTEGER                                NOT NULL,
    price      NUMERIC                                NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- seeder (dummy data)
INSERT INTO order_items (id, order_id, book_id, qty, price)
VALUES ('60b6fe65-1f58-4f6e-bed0-169ad3ce6c95', '9ed9b22b-ea2f-4c9e-ad67-7344c72a38ca',
        '48941f76-2652-49c3-a443-afdee02d1bff', 1, 390000),
       ('51f25213-8c0d-4690-951a-6d0794cba8b1', '9ed9b22b-ea2f-4c9e-ad67-7344c72a38ca',
        'e8660258-aef4-4024-8139-3581b140804e', 2, 230000),
       ('37410c24-36b7-4aa1-a798-57efce29ad00', '524d41bf-dfbe-44dc-a8a9-7f1c74788ad0',
        '6037389a-03c8-45a1-b2c0-791ff2cc0d6b', 1, 400000),
       ('9afdf73b-19a0-42a0-a35f-0c3620a87c9d', '833ae9a5-2482-43d1-b737-c08d21bdef90',
        '7371450f-f30e-46db-9957-3b3c6cbb773b', 1, 235000),
       ('d05b10f6-894d-442f-9fe5-8f838118f25b', '833ae9a5-2482-43d1-b737-c08d21bdef90',
        'e8660258-aef4-4024-8139-3581b140804e', 1, 230000),
       ('62e0edc2-9e8d-4ed3-9066-b52c7509c17f', 'ea8edfc4-1485-461b-802b-ba080b5b4902',
        '01b316fe-bed3-48e1-9c92-f58bb7c1061e', 1, 190000),
       ('2432c587-b6b4-4355-b9ca-276212015707', '864ccbc3-07d4-41d6-bb5a-31201231d29e',
        '62b1d43f-329b-46f8-b5e0-f9d1b8a1dfb7', 1, 260000),
       ('63dfae0e-325c-4ab6-b533-f66e7db49820', '1c5d77e1-4b0c-4355-81c4-02781d5de85f',
        '8d115de0-9278-4003-9308-e0c728fc8489', 1, 315000),
       ('e8d29ee1-569d-492c-b36f-da465fcb8302', '1c5d77e1-4b0c-4355-81c4-02781d5de85f',
        '062c9a67-768d-4fc2-9c61-d49d7b4b548c', 1, 375000),
       ('43622894-58a4-49d8-a22d-5a0962ba6f4b', '901841d8-c2d1-4aec-8200-43d9f3e1fb2b',
        '9ef088b3-9e1b-46bf-8176-6cb84499d783', 10, 360000);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "order_items" CASCADE;
-- +goose StatementEnd
