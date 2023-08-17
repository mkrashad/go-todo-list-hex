CREATE TABLE
    IF NOT EXISTS public.tasks (
        id serial PRIMARY KEY,
        created_at TIMESTAMP
        WITH
            TIME ZONE,
            updated_at TIMESTAMP
        WITH
            TIME ZONE,
            deleted_at TIMESTAMP
        WITH
            TIME ZONE,
            task_name VARCHAR(300),
            completed BOOL,
            user_id BIGINT
    );