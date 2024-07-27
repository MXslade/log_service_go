    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

    CREATE TABLE IF NOT EXISTS admins (
        id UUID DEFAULT uuid_generate_v4(),
        name VARCHAR(255) NOT NULL,
        password TEXT NOT NULL,

        PRIMARY KEY(id)
    );

