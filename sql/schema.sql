CREATE TABLE IF NOT EXISTS zombies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(40)
);

CREATE TABLE IF NOT EXISTS wizzes (
    id SERIAL PRIMARY KEY,
    content TEXT,
    created_at TIMESTAMP,

    zombie_id INTEGER NOT NULL REFERENCES zombies (id),
    CONSTRAINT wizzes_zombie_fk FOREIGN KEY (zombie_id) REFERENCES zombies (id)
);
