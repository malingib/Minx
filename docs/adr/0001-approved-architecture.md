# ADR-0001 — Approved Minx Architecture

- **Status:** Accepted
- **Date:** 2026-11-17

## Context
Minx is a mobile-first, creator-first social entertainment and commerce platform launching in Kenya and expanding across Africa.

## Decisions
1. Mobile: React Native + Expo, New Architecture, Android/iOS.
2. Public website: React + Vite, marketing and documentation only.
3. Backend: Go 1.27.4; Chi; pgx + sqlc.
4. Architecture: modular core plus separately deployed realtime, finance, media and worker runtimes.
5. Core database: managed PostgreSQL 18.
6. Finance: separate Go service and separate PostgreSQL cluster, credentials, deployment and audit boundary.
7. Cache: managed Redis-compatible service for ephemeral state only.
8. Events: managed NATS JetStream, transactional outbox/inbox and idempotent consumers.
9. Media: LiveKit for interactive RTC; Cloudflare Stream for broadcast/VOD; Go media control plane.
10. Storage: Cloudflare R2 for non-sensitive media; regional object storage for sensitive evidence/backups.
11. Search: PostgreSQL FTS initially; OpenSearch when justified.
12. Analytics: event archive and ClickHouse Cloud when justified.
13. Auth: managed OIDC; Minx owns profiles, organisations, roles and resource authorisation.
14. Messaging: PostgreSQL durable source of truth; WebSocket delivery.
15. Push: direct FCM and APNs.
16. APIs: REST/JSON + OpenAPI externally; gRPC/protobuf internally where useful.
17. Infrastructure: managed AWS foundation, Cloudflare edge/media, IaC; Cape Town region subject to availability and data-transfer validation.
18. E2EE: deferred; no custom cryptography.
19. Financial product: no general customer stored-value wallet, transfers or redeemable coins at launch; licensed partners handle money movement; creator/merchant earnings are contractual payables.

## Consequences
- Strong operational boundaries without premature microservices.
- Finance isolated from core application data and ordinary service credentials.
- Vendor choices are reference defaults; production procurement remains gated by regional availability, legal review, contract, benchmark and exit-path evidence.
- Events are at-least-once; handlers must be idempotent.
- Search and analytics are projections, not transaction authorities.

## Rejected alternatives
Node.js as primary backend; Rust as default; microservice-per-domain; direct client-to-PSP integration; Redis as durable truth; exactly-once assumptions; self-designed E2EE; full consumer web app at launch.

## Reopening criteria
Reopen only with benchmark evidence, regulatory/legal change, vendor unavailability, material cost/latency issue or a documented security requirement. Any change requires a superseding ADR and migration plan.
