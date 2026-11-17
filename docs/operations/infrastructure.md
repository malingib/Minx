# Infrastructure and Operations

## Reference deployment
Cloudflare provides DNS, WAF, edge cache, public delivery and R2. AWS Africa (Cape Town) is the initial application/data region for managed container compute, PostgreSQL, cache, secrets and telemetry, subject to confirming service availability. LiveKit Cloud provides interactive media; Cloudflare Stream provides broadcast/VOD; managed NATS JetStream provides durable events; ClickHouse Cloud is activated when analytics workload justifies it.

## Compute
AWS ECS on Fargate is the default managed container runtime. Do not use request-only serverless for persistent WebSocket gateways, media orchestration or durable workers unless connection duration, execution and networking limits are proven compatible. No Kubernetes at initial launch. Containers are deployment artifacts, not a mandatory local developer dependency.

## Databases and cache
Managed PostgreSQL 18; finance in separate cluster. Multi-AZ production where available, encrypted backups and PITR. No app superuser/schema-owner credentials. Connection pooling, query budgets, statement timeouts, controlled migrations and restore drills. Managed Redis-compatible cache is ephemeral with TTL and rebuild path.

## Broker and storage
NATS JetStream streams define retention, replication, durable consumers, ack policy, DLQ, replay, quotas and backup. Broker outage accumulates outbox rows. R2 stores non-sensitive media; approved regional storage stores sensitive evidence and backup copies. Private buckets by default, lifecycle policies, checksums and content-type verification.

## Environments
Dev, staging and production are isolated: separate cloud accounts/projects, DBs, OIDC tenants, signing keys, PSP credentials, buckets, broker credentials and analytics datasets. Staging uses provider sandbox credentials only.

## Observability
OpenTelemetry traces/metrics and correlated structured logs with PII redaction. Error tracking for mobile/backend. Dashboards for API latency, DB saturation, broker lag, websocket connections, media quality, payment state age, reconciliation exceptions, payout failure, moderation backlog and cost.

## Initial SLO/recovery objectives
Targets to validate: core API 99.9% monthly; finance API 99.95% after provider exclusions are defined; core RPO <=5 minutes; finance RPO <=1 minute where supported; core RTO <=60 minutes; finance RTO <=30 minutes. Quarterly restore/failover drills. Provider outages are reported separately.

## Incident response
Severity, on-call owner, incident commander, security/privacy escalation, customer communications, evidence preservation, rollback, recovery validation and post-incident review. Finance incidents require hold/freeze capability and reconciliation before payouts resume.
