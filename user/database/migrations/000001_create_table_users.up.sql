CREATE TABLE IF NOT EXISTS public.users
(
    id         serial PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    first_name   VARCHAR(300),
    last_name   VARCHAR(300),
    email      VARCHAR(300),
    user_name      VARCHAR(300),
    password      VARCHAR(300)
);