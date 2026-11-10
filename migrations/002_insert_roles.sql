INSERT INTO roles (name, description)
VALUES
    ('ADMIN', 'Full access to TribBox'),
    ('MODERATOR', 'Moderation permissions'),
    ('PLAYER', 'Standard player permissions')
ON CONFLICT (name) DO NOTHING;