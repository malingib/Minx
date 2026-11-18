\set ON_ERROR_STOP on

CREATE ROLE minx_app NOLOGIN;
GRANT USAGE ON SCHEMA identity, profiles, organisations TO minx_app;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA identity, profiles, organisations TO minx_app;

SET ROLE minx_app;
BEGIN;

SET LOCAL minx.account_id = '10000000-0000-4000-8000-000000000001';
INSERT INTO identity.accounts (id, auth_subject)
VALUES ('10000000-0000-4000-8000-000000000001', 'issuer|subject-a');

INSERT INTO profiles.profiles (id, account_id, handle, display_name, profile_type, visibility)
VALUES
 ('20000000-0000-4000-8000-000000000001', '10000000-0000-4000-8000-000000000001', 'creator_a', 'Creator A', 'CREATOR', 'PUBLIC'),
 ('20000000-0000-4000-8000-000000000001', '10000000-0000-4000-8000-000000000001', 'private_a', 'Private A', 'PERSONAL', 'PRIVATE');

INSERT INTO organisations.organisations (
 id, display_name, organisation_type, country_code, created_by
) VALUES (
 '30000000-0000-4000-8000-000000000001', 'A Studio', 'CREATOR_TEAM', 'KE',
 '10000000-0000-4000-8000-000000000001'
);

INSERT INTO organisations.roles (id, organisation_id, name, permissions)
VALUES (
 '40000000-0000-4000-8000-000000000001',
 '30000000-0000-4000-8000-000000000001',
 'Owner', '{"organisation.manage": true}'::jsonb
);

INSERT INTO organisations.memberships (organisation_id, account_id, role_id)
VALUES (
 '30000000-0000-4000-8000-000000000001',
 '10000000-0000-4000-8000-000000000001',
 '40000000-0000-4000-8000-000000000001'
);

SET LOCAL minx.account_id = '10000000-0000-4000-8000-000000000002';
INSERT INTO identity.accounts (id, auth_subject)
VALUES ('10000000-0000-4000-8000-000000000002', 'issuer|subject-b');

DO $$
BEGIN
    IF (SELECT count(*) FROM profiles.profiles WHERE handle = 'creator_a') <> 1 THEN
        RAISE EXCEPTION 'public profile should be visible to another account';
    END IF;
    IF (SELECT count(*) FROM profiles.profiles WHERE handle = 'private_a') <> 0 THEN
        RAISE EXCEPTION 'private profile leaked across accounts';
    END IF;
    IF (SELECT count(*) FROM organisations.organisations WHERE id = '30000000-0000-4000-8000-000000000001') <> 0 THEN
        RAISE EXCEPTION 'organisation leaked to non-member';
    END IF;
    IF (SELECT count(*) FROM identity.accounts WHERE id = '10000000-0000-4000-8000-000000000001') <> 0 THEN
        RAISE EXCEPTION 'account row leaked across accounts';
    END IF;
END $$;

ROLLBACK;
RESET ROLE;
DROP ROLE minx_app;
