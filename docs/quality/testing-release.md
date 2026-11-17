# Testing and Release Engineering

## CI gates
Formatting/lint; static analysis; build/type checks; unit tests; API/event contract compatibility; migration validation; integration tests; dependency/secret/security scanning; Android/iOS builds; critical E2E; performance regression checks; artifact provenance. No suppressed tests, ignored failures or fallback paths that hide broken integrations.

## Test layers
Unit domain rules/state machines/calculations; integration PostgreSQL/cache/broker/provider/object storage; contract OpenAPI/protobuf/event schemas; E2E registration, publishing, live join, checkout, payment, fulfilment, messaging and moderation; security IDOR, tenant isolation, auth bypass, replay, privilege escalation, injection and secret leakage; resilience broker outage, DB failover, worker restart, provider timeout, delayed callbacks and event replay; mobile low memory, older Android, background/resume, poor network and upgrades.

## Finance release blockers
No unbalanced journal, duplicate collection/payout, lost confirmed payment, cross-tenant access, silent journal edit, payout from unavailable balance or hidden reconciliation exception.

## Media load tests
Test 1/10/100/500/1,000/10,000 viewers where plan permits. Measure join success, latency, startup, rebuffer, packet loss, reconnect, CPU, battery, data use and cost per viewer-minute. Test host disconnect, network change, late join, moderation and recording.

## Release process
PR -> required CI -> review -> staging deploy -> migration rehearsal -> integration/E2E -> security/operations approval -> canary -> gradual rollout -> smoke checks -> monitor -> promote or rollback.

## Migration rules
Backward-compatible expand/contract. No destructive migration in same release as dependent code removal. Backup before high-risk migrations. Test rollback/forward repair. Finance schema changes require finance owner approval and reconciliation verification.

## Mobile release
Internal distribution, closed beta, staged rollout, crash monitoring and explicit API compatibility window. Forced upgrade only for security or protocol break.
