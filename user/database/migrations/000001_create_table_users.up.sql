CREATE TABLE IF NOT EXISTS public.users
(
    id         serial PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    firstname   VARCHAR(300),
    lastname   VARCHAR(300),
    email      VARCHAR(300)
);