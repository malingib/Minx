\\set ON_ERROR_STOP on

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
 ('20000000-0000-4000-8000-000000000002', '10000000-0000-4000-8000-000000000001', 'private_a', 'Private A', 'PERSONAL', 'PRIVATE');

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

SAVEPOINT before_spoofed_profile;
\set ON_ERROR_STOP off
INSERT INTO profiles.profiles (account_id, handle, display_name, profile_type, visibility)
VALUES ('10000000-0000-4000-8000-000000000001', 'spoofed_a', 'Spoofed', 'PERSONAL', 'PUBLIC');
\if :ERROR
  ROLLBACK TO SAVEPOINT before_spoofed_profile;
  \echo 'PASS: cross-account profile insert rejected by RLS'
\else
  ROLLBACK TO SAVEPOINT before_spoofed_profile;
  \echo 'FAIL: cross-account profile insert unexpectedly succeeded'
  \quit 1
\endif
\set ON_ERROR_STOP on

SELECT count(*) = 1 AS public_profile_visible
FROM profiles.profiles WHERE handle = 'creator_a' \gset
\if :public_profile_visible
\else
  \echo 'FAIL: public profile was not visible'
  \quit 1
\endif

SELECT count(*) = 0 AS private_profile_hidden
FROM profiles.profiles WHERE handle = 'private_a' \gset
\if :private_profile_hidden
\else
  \echo 'FAIL: private profile leaked across accounts'
  \quit 1
\endif

SELECT count(*) = 0 AS organisation_hidden
FROM organisations.organisations
WHERE id = '30000000-0000-4000-8000-000000000001' \gset
\if :organisation_hidden
\else
  \echo 'FAIL: organisation leaked to a non-member'
  \quit 1
\endif

SELECT count(*) = 0 AS account_hidden
FROM identity.accounts
WHERE id = '10000000-0000-4000-8000-000000000001' \gset
\if :account_hidden
\else
  \echo 'FAIL: account row leaked across accounts'
  \quit 1
\endif

ROLLBACK;
RESET ROLE;
DROP ROLE minx_app;
