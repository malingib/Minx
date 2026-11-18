-- name: CreateAccount :one
INSERT INTO identity.accounts (id, auth_subject)
VALUES ($1, $2)
RETURNING id, auth_subject, status, created_at, updated_at;

-- name: GetAccountByAuthSubject :one
SELECT id, auth_subject, status, created_at, updated_at
FROM identity.accounts
WHERE auth_subject = $1;

-- name: CreateProfile :one
INSERT INTO profiles.profiles (
    id, account_id, handle, display_name, profile_type, visibility
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, account_id, handle, display_name, profile_type, visibility, status,
          created_at, updated_at;

-- name: GetProfileByHandle :one
SELECT id, account_id, handle, display_name, profile_type, visibility, status,
       created_at, updated_at
FROM profiles.profiles
WHERE handle = lower($1) AND status = 'ACTIVE';

-- name: ListProfilesByAccount :many
SELECT id, account_id, handle, display_name, profile_type, visibility, status,
       created_at, updated_at
FROM profiles.profiles
WHERE account_id = $1
ORDER BY created_at, id;
