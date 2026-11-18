# Core database

Core persistence is a managed PostgreSQL cluster. Finance uses its own independent cluster and credentials; Core must never connect to or write Finance tables.

## Initial migration

- `migrations/000001_identity_organisations.up.sql` establishes account, profile and organisation ownership boundaries.
- The matching `.down.sql` is for controlled development rollback only. Never use destructive down migrations as a production recovery strategy.
- The application runtime must use a non-owner database role without `BYPASSRLS`. The migration role owns schema changes and is not used by the API.
- Authenticated request transactions set `SET LOCAL minx.account_id = '<verified account UUID>'` before executing tenant-scoped queries. Never derive this value from an untrusted request body or header.
- RLS is defence in depth. Application-level actor/resource/action checks remain mandatory, with cross-organisation IDOR tests.

## Query generation

The sqlc configuration generates pgx/v5 query bindings from the migration schema and `db/queries`. Generated files are committed so application builds do not require sqlc at runtime.

## Migration discipline

Migrations are append-only after they reach a shared environment. Apply them through a dedicated migration identity, verify the resulting schema version, and test both clean database creation and upgrade from the previous released schema. No production database is connected or modified by this repository change.
