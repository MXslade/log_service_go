CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS apps (
    id UUID DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,

    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS app_envs (
    id UUID DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    app_id UUID NOT NULL,

    PRIMARY KEY(id),
    CONSTRAINT fk__app_env__app
        FOREIGN KEY(app_id)
        REFERENCES apps(id) 
        ON DELETE CASCADE
);

