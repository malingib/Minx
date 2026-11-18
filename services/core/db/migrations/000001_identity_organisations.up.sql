CREATE SCHEMA IF NOT EXISTS identity;
CREATE SCHEMA IF NOT EXISTS profiles;
CREATE SCHEMA IF NOT EXISTS organisations;

CREATE TABLE identity.accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    auth_subject TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL DEFAULT 'ACTIVE'
        CHECK (status IN ('ACTIVE', 'SUSPENDED', 'CLOSED')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE profiles.profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES identity.accounts(id),
    handle TEXT NOT NULL
        CHECK (handle = lower(handle) AND handle ~ '^[a-z0-9_]{3,30}$'),
    display_name TEXT NOT NULL CHECK (length(trim(display_name)) BETWEEN 1 AND 80),
    profile_type TEXT NOT NULL
        CHECK (profile_type IN ('PERSONAL', 'CREATOR', 'BUSINESS')),
    visibility TEXT NOT NULL DEFAULT 'PUBLIC'
        CHECK (visibility IN ('PUBLIC', 'PRIVATE')),
    status TEXT NOT NULL DEFAULT 'ACTIVE'
        CHECK (status IN ('ACTIVE', 'SUSPENDED', 'CLOSED')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX profiles_handle_lower_uq ON profiles.profiles (lower(handle));
CREATE INDEX profiles_account_id_idx ON profiles.profiles (account_id);

CREATE TABLE organisations.organisations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    legal_name TEXT,
    display_name TEXT NOT NULL CHECK (length(trim(display_name)) BETWEEN 1 AND 120),
    organisation_type TEXT NOT NULL
        CHECK (organisation_type IN ('BUSINESS', 'CREATOR_TEAM', 'NONPROFIT', 'COMMUNITY', 'OTHER')),
    country_code CHAR(2) NOT NULL,
    status TEXT NOT NULL DEFAULT 'ACTIVE'
        CHECK (status IN ('ACTIVE', 'SUSPENDED', 'CLOSED')),
    created_by UUID NOT NULL REFERENCES identity.accounts(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE organisations.roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organisation_id UUID NOT NULL REFERENCES organisations.organisations(id),
    name TEXT NOT NULL,
    permissions JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (organisation_id, name),
    UNIQUE (id, organisation_id)
);

CREATE TABLE organisations.memberships (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organisation_id UUID NOT NULL REFERENCES organisations.organisations(id),
    account_id UUID NOT NULL REFERENCES identity.accounts(id),
    role_id UUID NOT NULL,
    status TEXT NOT NULL DEFAULT 'ACTIVE'
        CHECK (status IN ('INVITED', 'ACTIVE', 'SUSPENDED', 'REMOVED')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (organisation_id, account_id),
    FOREIGN KEY (role_id, organisation_id)
        REFERENCES organisations.roles(id, organisation_id)
);
CREATE INDEX memberships_account_id_idx ON organisations.memberships (account_id);

CREATE TABLE profiles.operators (
    profile_id UUID NOT NULL REFERENCES profiles.profiles(id),
    account_id UUID NOT NULL REFERENCES identity.accounts(id),
    permission_set JSONB NOT NULL DEFAULT '{}'::jsonb,
    status TEXT NOT NULL DEFAULT 'ACTIVE'
        CHECK (status IN ('ACTIVE', 'SUSPENDED', 'REMOVED')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (profile_id, account_id)
);

-- The runtime sets this transaction-local value only after authenticating the actor.
-- The migration owner remains responsible for schema changes and is not the app role.
ALTER TABLE identity.accounts ENABLE ROW LEVEL SECURITY;
ALTER TABLE profiles.profiles ENABLE ROW LEVEL SECURITY;
ALTER TABLE organisations.organisations ENABLE ROW LEVEL SECURITY;
ALTER TABLE organisations.roles ENABLE ROW LEVEL SECURITY;
ALTER TABLE organisations.memberships ENABLE ROW LEVEL SECURITY;
ALTER TABLE profiles.operators ENABLE ROW LEVEL SECURITY;

CREATE POLICY accounts_self_read ON identity.accounts
    FOR SELECT USING (id = nullif(current_setting('minx.account_id', true), '')::uuid);

CREATE POLICY profiles_public_or_owner_read ON profiles.profiles
    FOR SELECT USING (
        (visibility = 'PUBLIC' AND status = 'ACTIVE')
        OR account_id = nullif(current_setting('minx.account_id', true), '')::uuid
    );
CREATE POLICY profiles_owner_insert ON profiles.profiles
    FOR INSERT WITH CHECK (
        account_id = nullif(current_setting('minx.account_id', true), '')::uuid
    );
CREATE POLICY profiles_owner_update ON profiles.profiles
    FOR UPDATE USING (
        account_id = nullif(current_setting('minx.account_id', true), '')::uuid
    ) WITH CHECK (
        account_id = nullif(current_setting('minx.account_id', true), '')::uuid
    );

CREATE POLICY organisations_member_read ON organisations.organisations
    FOR SELECT USING (
        EXISTS (
            SELECT 1 FROM organisations.memberships m
            WHERE m.organisation_id = organisations.organisations.id
              AND m.account_id = nullif(current_setting('minx.account_id', true), '')::uuid
              AND m.status = 'ACTIVE'
        )
    );
CREATE POLICY organisations_creator_insert ON organisations.organisations
    FOR INSERT WITH CHECK (
        created_by = nullif(current_setting('minx.account_id', true), '')::uuid
    );

CREATE POLICY organisation_roles_member_read ON organisations.roles
    FOR SELECT USING (
        EXISTS (
            SELECT 1 FROM organisations.memberships m
            WHERE m.organisation_id = roles.organisation_id
              AND m.account_id = nullif(current_setting('minx.account_id', true), '')::uuid
              AND m.status = 'ACTIVE'
        )
    );
CREATE POLICY memberships_self_read ON organisations.memberships
    FOR SELECT USING (
        account_id = nullif(current_setting('minx.account_id', true), '')::uuid
    );
CREATE POLICY memberships_self_insert ON organisations.memberships
    FOR INSERT WITH CHECK (
        account_id = nullif(current_setting('minx.account_id', true), '')::uuid
    );

CREATE POLICY profile_operators_owner_or_self_read ON profiles.operators
    FOR SELECT USING (
        account_id = nullif(current_setting('minx.account_id', true), '')::uuid
        OR EXISTS (
            SELECT 1 FROM profiles.profiles p
            WHERE p.id = operators.profile_id
              AND p.account_id = nullif(current_setting('minx.account_id', true), '')::uuid
        )
    );
