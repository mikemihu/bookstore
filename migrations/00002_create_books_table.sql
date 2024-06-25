-- +goose Up
-- +goose StatementBegin
CREATE TABLE "books"
(
    id         UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    isbn       VARCHAR(13) UNIQUE                     NOT NULL,
    author     TEXT                                   NOT NULL,
    title      TEXT                                   NOT NULL,
    subtitle   TEXT                                   NOT NULL,
    price      NUMERIC                                NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- seeder (dummy data)
INSERT INTO books (id, isbn, author, title, subtitle, price)
VALUES ('6037389a-03c8-45a1-b2c0-791ff2cc0d6b', '9780804137386', 'Greg McKeown',
        'Essentialism', 'The Disciplined Pursuit of Less', 400000),
       ('0275a3e7-6922-4071-9e61-c2a2216529fe', '9781443442312', 'Angela Duckworth',
        'Grit', 'The Power of Passion and Perseverance', 350000),
       ('9ef088b3-9e1b-46bf-8176-6cb84499d783', '9781847941831', 'James Clear',
        'Atomic Habits', 'The life-changing million copy bestseller', 360000),
       ('8d115de0-9278-4003-9308-e0c728fc8489', '9780857197689', 'Morgan Housel',
        'The Psychology of Money', 'Timeless lessons on wealth, greed, and happiness', 315000),
       ('062c9a67-768d-4fc2-9c61-d49d7b4b548c', '9780743269513', 'Stephen R. Covey',
        '7 Basic Habits of Highly Effective People', 'Powerful Lessons in Personal Change', 375000),
       ('48941f76-2652-49c3-a443-afdee02d1bff', '9780374533557', 'Daniel Kahneman',
        'Thinking, fast and slow', '', 390000),
       ('62b1d43f-329b-46f8-b5e0-f9d1b8a1dfb7', '9781591847816', 'Ryan Holiday',
        'Ego Is the Enemy', 'The Fight to Master Our Greatest Opponent', 260000),
       ('01b316fe-bed3-48e1-9c92-f58bb7c1061e', '9780751532715', 'Robert T. Kiyosaki',
        'Rich Dad, Poor Dad', '', 190000),
       ('e8660258-aef4-4024-8139-3581b140804e', '9780307352149', 'Susan Cain',
        'Quiet', 'The Power of Introverts in a World That Can''t Stop Talking', 230000),
       ('7371450f-f30e-46db-9957-3b3c6cbb773b', '9780525429562', 'Adam M. Grant',
        'Originals', 'How Non-Conformists Move the World', 235000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "books" CASCADE;
-- +goose StatementEnd
