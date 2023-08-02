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

ALTER TABLE public.tasks
ADD
    CONSTRAINT fk_tasks_users FOREIGN KEY (user_id) REFERENCES public.users (id);