DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'STATUS') THEN
        CREATE TYPE STATUS AS ENUM
        (
            'UNAVAILABLE',
            'HIDDEN'
        );
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS menu_items (
    id              VARCHAR                      UNIQUE PRIMARY KEY NOT NULL,
    status          STATUS                       NOT NULL,
    available_at    TIMESTAMP WITH TIME ZONE     NULL
);