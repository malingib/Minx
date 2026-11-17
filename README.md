# Minx — Social Super App

**Architecture baseline v1.0 · 17 November 2026**

Minx is an independent, mobile-first, creator-first social entertainment and commerce platform launching in Kenya and expanding across Africa. This repository is the engineering source of truth. The product is not a MobiWave subsidiary or MobiWave-branded product; external services may be procured from MobiWave as vendors without changing Minx's independent identity.

## Product pillars
1. Social — profiles, short video, posts, feed, discovery and social graph.
2. Live — interactive sessions, co-hosting, battles and live shopping.
3. Market — creator/merchant shops, catalogue, inventory, cart, checkout, orders and fulfilment.
4. Connect — persistent messaging, presence, calls and notifications.
5. Wallet & earnings — payment history and contractual creator/merchant payable summaries; not an unauthorised stored-value wallet.
6. Creator Studio — publishing, audience analytics, monetisation configuration and shop tools.
7. Communities — groups, membership, discussion and events.
8. Business — merchant organisations, staff roles, orders, fulfilment and reporting.

## Approved technology baseline
| Layer | Decision |
|---|---|
| Mobile | React Native + Expo, New Architecture, Android and iOS |
| Public website | React + Vite; marketing/documentation only |
| Backend | Go 1.27.4 baseline |
| HTTP | Chi |
| Database access | pgx + sqlc; no ORM |
| Transactional database | Managed PostgreSQL 18 |
| Architecture | Modular monolith with separately deployed critical runtimes |
| Finance | Isolated Go service and separate PostgreSQL cluster |
| Cache | Managed Redis-compatible service; ephemeral state only |
| Event broker | Managed NATS JetStream |
| Interactive media | LiveKit |
| Broadcast/VOD | Cloudflare Stream |
| Media control plane | Go media orchestrator |
| Object storage | Cloudflare R2 for non-sensitive media; regional managed object storage for sensitive evidence/backups |
| Search | PostgreSQL FTS initially; OpenSearch when justified |
| Analytics | Event archive; ClickHouse Cloud when justified |
| Authentication | Managed OIDC; Minx owns account, profile and authorisation model |
| Messaging | PostgreSQL durable store; WebSocket realtime gateway |
| Calls | LiveKit |
| Push | Direct FCM and APNs |
| APIs | REST/JSON + OpenAPI externally; gRPC/protobuf internally where justified |
| Observability | OpenTelemetry, central logs/metrics/traces and error tracking |
| Infrastructure | Managed AWS foundation + Cloudflare edge/media; IaC |
| Initial cloud region | AWS Africa (Cape Town), conditional on verifying required service availability and data-transfer compliance |

## Runtime boundaries
1. **Core API (Go):** profiles, organisations, social graph, content metadata, feed orchestration, catalogues, inventory, carts, orders, communities and creator/merchant configuration.
2. **Realtime gateway (Go):** WebSockets, presence, message fanout, live application events and delivery receipts.
3. **Financial service (Go):** payment orchestration, provider adapters, double-entry ledger, payables, refunds, payouts, reconciliation, risk and audit. Separate DB and credentials.
4. **Media orchestrator (Go):** provider-neutral session lifecycle, short-lived tokens, permissions, recording policy, webhooks and takedowns. Never proxies media bytes.
5. **Worker runtime (Go):** outbox publication, consumers, notifications, fulfilment, media coordination and scheduled jobs.
6. **Intelligence runtime (Python, later):** recommendations, ranking evaluation, feature extraction and fraud models; not transactional source of truth.

## Non-negotiable financial controls
- No general-purpose customer stored-value wallet, customer-to-customer transfers or transferable/cash-redeemable coins at launch.
- Use appropriately licensed payment partners for collections and payouts.
- Creator and merchant earnings are contractual payables, not customer deposits.
- Double-entry ledger: debits equal credits per journal and currency.
- Integer minor units only; never floating-point money.
- Posted journals are immutable; corrections use reversals.
- Idempotency at API, provider callback, consumer and payout layers.
- Provider timeouts are UNKNOWN until queried/reconciled; never blindly retry a payout.
- Legal review of PSP scope, marketplace settlement, gifts, tax, KYC/KYB and AML is a production gate.

## Architecture rules
- Core modules have explicit ownership and contracts; do not create a microservice per domain prematurely.
- Finance is a hard security and database boundary from day one.
- Orders are not payments; creator earnings are not ledger balances; inventory is not catalogue.
- PostgreSQL is transactional truth. Search, cache and analytics are derived projections.
- Redis is ephemeral only. NATS delivery is at-least-once; consumers must be idempotent. Do not assume exactly-once delivery.
- Transactional outbox/inbox for durable cross-domain events.
- Resource-level server authorisation on every request. Tenant isolation is tested; RLS is defence in depth, not the sole control.
- No E2EE in launch baseline. Add only after selecting and auditing a recognised protocol and designing reporting/multi-device recovery.
- Dev, staging and production are isolated environments.
- No suppressed tests, fake success paths or fallbacks that hide broken integrations.

## Delivery stages
- **Stage 0:** repository, CI/CD, IaC, identity, data boundaries, finance isolation, outbox, observability and threat model.
- **Stage 1:** profiles, social graph, posts/video, feed, discovery, messaging, notifications and moderation.
- **Stage 2:** live, live shopping, merchant catalogue, checkout, licensed PSP integration, ledger, payables and reconciliation.
- **Stage 3:** gifts, memberships, affiliate, battles, communities and group calls, each with separate legal/financial/abuse design.
- **Stage 4:** regional payment rails, currencies, languages, shipping/tax, advanced recommendations, events and ads.

## Production gates (not assumed complete)
- Verify actual AWS Cape Town availability for every required managed service.
- Complete Kenya data protection, DPIA/ODPC registration and cross-border transfer assessment as applicable.
- Obtain Kenyan legal advice and confirm PSP authorisation/scope and signed contracts.
- Benchmark LiveKit + Stream end-to-end, including latency, quality and cost.
- Verify managed NATS region, SLA, backup and exit/export path.
- Select OIDC provider after regional, recovery, MFA, pricing and export review.
- Define secondary-region disaster recovery and approve any cross-border replication.

## Documentation status
This README establishes the initial architecture baseline. Detailed technical specifications, ADRs, data/API contracts, finance/media/security/operations specs and acceptance tests are to be maintained under `docs/`. Do not mark vendor, legal, deployment or benchmark gates complete without evidence.

## Change control
Any change to an approved technology or boundary requires an Architecture Decision Record containing context, alternatives, decision, consequences, migration impact, owner and date. The ADR register is authoritative.
