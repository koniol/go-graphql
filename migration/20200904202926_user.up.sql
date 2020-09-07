CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     email VARCHAR(255),
                                     password text not null,
                                     created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                                     updated_at TIMESTAMP WITH TIME ZONE,
                                     deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

INSERT INTO users (email, password) VALUES ('admin@admin.pl', 'admin');
INSERT INTO users (email, password) VALUES ('admin1@admin.pl', 'admin1');