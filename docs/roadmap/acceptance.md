# Roadmap and Acceptance

## Programme baseline

The initial funded programme is nine months: **six months of development followed by three months of live beta**. The cash funding envelope is KES 2,500,000, including the paid KES 400,000 founder/lead-developer allocation. See the [funding, delivery and pilot capacity plan](funding-capacity-delivery-plan.md) for the budget, provider assumptions, cash-release schedule and capacity limits.

### Months 1–2 — Engineering foundation
Repository/module boundaries; CI/CD and IaC; dev/staging/prod; OIDC; account/profile/organisation model; core DB and physically separate finance DB; outbox/event contracts; observability, secrets and audit; threat model and data inventory. Exit when builds/deploys are reproducible, migrations and service boundaries are tested, and restore drill is completed.

### Months 3–4 — Social, creator and commerce foundations
Profiles, social graph, content, short video, feed, discovery, Creator Studio basics, messaging, notifications, reporting and moderation console. Build seller profiles, catalogue, variants and inventory foundations. Exit when core journeys pass E2E, moderation is operational and low-bandwidth targets are met.

### Months 5–6 — Integration and release candidate
LiveKit interactive sessions, Cloudflare Stream broadcast/VOD, live-shopping pins, cart/orders, PSP integration, separate finance ledger, payables, refunds and reconciliation. Exit when financial blockers pass, provider production approval/legal sign-off exists, and media load/cost tests are accepted. Critical/high finance or security defects block beta entry.

### Month 7 — Closed beta
Invite-only users, selected creators and merchants. Validate payment state transitions, callback idempotency/replay, reconciliation, messaging persistence, media workflows, moderation and support operations.

### Month 8 — Expanded live beta
Expand onboarding only after month-7 acceptance. Test the agreed pilot workload, media cost limits, weak-network behaviour, API/database saturation, incident response and security remediation.

### Month 9 — Final beta and controlled public pilot
Resolve critical/high defects, complete restore and incident-response drills, verify production payment approval and operational readiness, and open controlled public registration only after all release gates pass.

## Initial public-pilot capacity limits

| Metric | Pilot target |
|---|---:|
| Registered users | 10,000 |
| Monthly active users (planning target) | 2,500 |
| Monthly video viewing | 2,500 viewer-hours |
| Peak concurrent broadcast viewers | 150 |
| Initial geography | Kenya |

These are planning targets, not proven capacity guarantees. They must be demonstrated on the exact release candidate and provider tiers. The 150-person broadcast audience is to use Cloudflare Stream delivery; LiveKit is reserved for the interactive stage unless a measured cost and capacity test approves a different topology. The earlier 100,000-account, 25,000-viewer-hour and 500-concurrent-viewer ambition is a later growth objective, not part of this KES 2.5M pilot commitment.

## Launch acceptance

- Account/profile/organisation and cross-tenant tests pass.
- Feed and publishing pass low-bandwidth mobile tests.
- Media join/reconnect/moderation work.
- 150 concurrent broadcast-viewer test passes the agreed soak and quality thresholds.
- 2,500 monthly viewer-hours are modelled, metered and reconciled against provider billing.
- Checkout/order state is correct; finance ledger invariants pass.
- PSP callbacks, idempotency, reconciliation and payout protection pass.
- Messaging persists and recovers after reconnect.
- Trust/safety workflows are staffed.
- Privacy inventory, DPIA/registration assessment and processor agreements complete.
- Backup restore, incident response, monitoring, on-call and support procedures are active.
- Actual infrastructure/media costs are within approved limits or a revised funding decision is documented.
- No unresolved critical/high security or financial correctness defects remain.
