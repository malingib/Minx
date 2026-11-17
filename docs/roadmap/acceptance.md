# Roadmap and Acceptance

## Stage 0 — Engineering foundation
Repository/module boundaries; CI/CD and IaC; dev/staging/prod; OIDC; account/profile/organisation model; core DB and isolated finance DB; outbox/event contracts; observability, secrets and audit; threat model and data inventory. Exit when builds/deploys are reproducible, migrations and service boundaries tested, and restore drill completed.

## Stage 1 — Social and creator
Profiles, social graph, content, short video, feed, discovery, Creator Studio basics, messaging, notifications, reporting and moderation console. Exit when core journeys pass E2E, moderation is operational and low-bandwidth targets are met.

## Stage 2 — Live and creator commerce
LiveKit sessions, Stream broadcast/VOD, live-shopping pins, shops/catalogue/inventory/cart/orders, PSP integration, ledger, payables, refunds and reconciliation. Exit when financial blockers pass, provider production approval/legal sign-off exists, and media load/cost tests are accepted.

## Stage 3 — Creator economy and communities
Gifts, memberships, affiliate, communities, battles and group calls. Each receives independent legal, financial, abuse and state-machine design. Exit when earning calculations are reproducible and disputes/fraud operations exist.

## Stage 4 — Regional ecosystem
Regional payment rails, currencies, languages, shipping/tax, merchant integrations, events, advanced recommendations, ads and partner ecosystem. Country launch requires country-specific legal and operational approval.

## Launch acceptance
- Account/profile/organisation and cross-tenant tests pass.
- Feed and publishing pass low-bandwidth mobile tests.
- Media join/reconnect/moderation work.
- Checkout/order state correct; finance ledger invariants pass.
- PSP callbacks, idempotency, reconciliation and payout protection pass.
- Messaging persists and recovers after reconnect.
- Trust/safety workflows are staffed.
- Privacy inventory, DPIA/registration assessment and processor agreements complete.
- Backup restore, incident response, monitoring, on-call and support procedures are active.
