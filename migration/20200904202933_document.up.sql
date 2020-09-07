CREATE TABLE IF NOT EXISTS documents
(
    id          SERIAL PRIMARY KEY,
    name        text,
    description text,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE,
    user_id     INT REFERENCES users (id) NOT NULL,
    deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

INSERT INTO documents (user_id, name, description)
VALUES (1, 'test', 'testadssd1');
INSERT INTO documents (user_id, name, description)
VALUES (1, 'test1', 'testadssd2');
INSERT INTO documents (user_id, name, description)
VALUES (2, 'tes3t', 'tddddestadssd1');
INSERT INTO documents (user_id, name, description)
VALUES (2, 'tesasdasd3t', 'tddddestadssd1');
INSERT INTO documents (user_id, name, description)
VALUES (2, 'teasdasdasdasds3t', 'tddddestadssd1');
