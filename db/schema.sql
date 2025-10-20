CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT UNIQUE,
    email TEXT UNIQUE NOT NULL,
    password  TEXT NOT NULL,
    useragent TEXT,
    picture TEXT ,
    phone_number TEXT,
    bio TEXT,
    role TEXT DEFAULT 'user',
    google_id TEXT UNIQUE ,
    login_option TEXT DEFAULT 'signup',
    revoked BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

