CREATE ROLE postgres;

\c postgres;

create table "users"
(
    "id" SERIAL PRIMARY KEY,
    "username" character varying(255) NOT NULL,
    "email" character varying(255) NOT NULL,
    "password" character varying(255) NOT NULL,
    "created_at" timestamp(0),
    "updated_at" timestamp(0)
);

create table "messages" (
    "id" SERIAL PRIMARY KEY,
    "message" TEXT NOT NULL,
    "sender" INT NOT NULL,
    "receiver" INT NOT NULL,
    "deleted_at" TIMESTAMP,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp(0),
    FOREIGN KEY (sender) REFERENCES "users" (id),
    FOREIGN KEY (receiver) REFERENCES "users" (id)
);

INSERT INTO "users" ("username", "email", "password", "created_at", "updated_at")
VALUES
    ('tom', 'tom@example.com', '$2a$10$z1Unfo0B28zjo85ulUQ7Y.VQ1zy9biPGAC7XSJeZ.EByCDaX8E.L6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('rayou', 'rayou@example.com', '$2a$10$z1Unfo0B28zjo85ulUQ7Y.VQ1zy9biPGAC7XSJeZ.EByCDaX8E.L6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('hassan', 'hassan@example.com', '$2a$10$z1Unfo0B28zjo85ulUQ7Y.VQ1zy9biPGAC7XSJeZ.EByCDaX8E.L6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('hassannou', 'hassan@example.com', '$2a$10$z1Unfo0B28zjo85ulUQ7Y.VQ1zy9biPGAC7XSJeZ.EByCDaX8E.L6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('hassaninou', 'hassan@example.com', '$2a$10$z1Unfo0B28zjo85ulUQ7Y.VQ1zy9biPGAC7XSJeZ.EByCDaX8E.L6', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO "messages" ("message", "sender", "receiver", "created_at", "deleted_at")
VALUES
    ('Salut rayou', 1, 2, CURRENT_TIMESTAMP, NULL),
    ('Salut Tom ', 2, 1, CURRENT_TIMESTAMP, NULL);


INSERT INTO "messages" ("message", "sender", "receiver", "created_at", "deleted_at")
VALUES
    ('Ca va?', 1, 2, CURRENT_TIMESTAMP, NULL);

INSERT INTO "messages" ("message", "sender", "receiver", "created_at", "deleted_at")
VALUES
    ('Hassan ???????', 1, 3, CURRENT_TIMESTAMP, NULL),
    ('Ca roule et toi ? :)', 2, 1, CURRENT_TIMESTAMP, NULL),
    ('Oui ???', 2, 1, CURRENT_TIMESTAMP, NULL);


-- Créer une fonction qui sera appelée par le déclencheur
CREATE OR REPLACE FUNCTION notify_new_message()
RETURNS TRIGGER AS $$
BEGIN
    -- Insérer les données nécessaires dans une file d'attente (en utilisant pg_notify)
    PERFORM pg_notify(
        'new_message',
        json_build_object(
            'table', TG_TABLE_NAME,
            'sender', NEW.sender,
            'receiver', NEW.receiver,
            'message', NEW.message,
            'created_at', NEW.created_at
        )::text
    );

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Créer un déclencheur qui appelle la fonction pour chaque insertion dans la table messages
CREATE TRIGGER messages_after_insert
    AFTER INSERT ON messages
    FOR EACH ROW
    EXECUTE FUNCTION notify_new_message();