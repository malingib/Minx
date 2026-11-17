# Nine-Month Funding, Delivery and Pilot Capacity Plan

**Status:** Approved planning baseline, subject to supplier quotes and benchmark gates  
**Programme:** 6 months development + 3 months live beta  
**Currency:** Kenyan shillings (KES)  
**Total cash funding envelope:** KES 2,500,000

## 1. Executive baseline

Minx will be delivered through a founder-led, milestone-based programme over nine months. The founder/lead developer is paid; the KES 400,000 allocation is a cash cost and is not an in-kind contribution.

| Measure | Planning baseline |
|---|---:|
| Total project cash budget | KES 2,500,000 |
| Development and professional fees | KES 1,450,000 |
| Development infrastructure/tools (months 1–6) | KES 150,000 |
| Live beta operating budget (months 7–9) | KES 650,000 |
| Central contingency reserve | KES 250,000 |
| Development period | 6 months |
| Live beta period | 3 months |
| Registered-user target at public pilot | 10,000 |
| Monthly active-user planning target | 2,500 |
| Monthly video viewing | 2,500 viewer-hours |
| Peak concurrent broadcast viewers | 150 |
| Initial market | Kenya |

The user and viewing numbers are launch planning limits, not demonstrated capacity or a vendor guarantee. Registration should be staged and may be paused automatically if cost, safety, reliability or support thresholds are exceeded.

## 2. Development and professional fees

| Workstream | Cash allocation |
|---|---:|
| Go backend contractor | KES 350,000 |
| React Native developer | KES 250,000 |
| UX/UI design | KES 100,000 |
| QA and security | KES 150,000 |
| DevOps | KES 100,000 |
| Legal and payment integration support | KES 100,000 |
| Lead developer — founder, paid | KES 400,000 |
| **Total** | **KES 1,450,000** |

These are project allocations, not supplier quotations or a statement of tax treatment. Contracts should specify deliverables, acceptance, IP assignment, payment milestones and support/defect obligations.

## 3. Full project budget

| Budget category | Allocation |
|---|---:|
| Development and professional fees | KES 1,450,000 |
| Development infrastructure, environments and tools (months 1–6) | KES 150,000 |
| Live beta infrastructure and operations (months 7–9) | KES 650,000 |
| Central contingency reserve | KES 250,000 |
| **Total cash budget** | **KES 2,500,000** |

The contingency is centrally held and released only against an approved variance. It is not pre-authorised feature scope or routine operating spend.

## 4. Three-month live beta budget

| Item | Three-month allocation |
|---|---:|
| AWS compute, Core and separate Finance databases, backups, networking and data transfer | KES 180,000 |
| Cloudflare Stream delivery and storage allowance | KES 80,000 |
| LiveKit Cloud Ship plan allowance | KES 20,000 |
| Monitoring, logging, alerting and security tooling | KES 35,000 |
| QA, defect resolution and technical support | KES 130,000 |
| Creator and merchant onboarding | KES 60,000 |
| User support and beta operations | KES 65,000 |
| Payment testing, callbacks and reconciliation operations | KES 35,000 |
| App-store and public-pilot release administration | KES 25,000 |
| Beta-specific operating buffer | KES 20,000 |
| **Total live beta allocation** | **KES 650,000** |

AWS and support allocations are planning envelopes. They must be replaced or reconciled against a region- and configuration-specific AWS Pricing Calculator estimate, provider invoices and actual beta telemetry before public-pilot expansion.

## 5. Media and realtime cost calculation

### 5.1 Cloudflare Stream

Use Cloudflare Stream for broadcast/VOD delivery to the audience. The public pricing page states USD 1 per 1,000 delivered minutes and USD 5 per 1,000 minutes of stored-video capacity per month.

Planning calculation:

- 2,500 viewer-hours/month × 60 = 150,000 delivered minutes/month.
- Three-month beta = 450,000 delivered minutes.
- Delivery: 450,000 / 1,000 × USD 1 = USD 450.
- Reserved stored-video capacity: 10,000 minutes × USD 5 / 1,000 = USD 50/month.
- Three months of storage capacity: USD 150.
- Total estimated Stream charges: USD 600.

For budgeting, use KES 130/USD as a rounded planning conversion (CBK indicative rate was KES 129.98/USD on 17 November 2026): USD 600 ≈ KES 78,000. The KES 80,000 beta line provides a small FX/usage allowance, not unlimited media usage.

The 10,000-minute storage capacity is the total active stored-duration allowance, not 10,000 new minutes every month. It must cover uploaded videos and any retained RTMP/SRT live recordings. Apply retention, deletion and upload-duration policies. Delivery is usage-based; preloading/buffering and cache behaviour can affect billable minutes.

### 5.2 LiveKit

The published LiveKit Cloud Ship starting price is USD 50/month. The pricing page lists 150,000 WebRTC participant minutes/month, 250 GB downstream transfer and up to 1,000 concurrent connections on Ship. Three months at the starting subscription price is USD 150, approximately KES 19,500 at the planning FX rate. The KES 20,000 line is a rounded allowance.

Architecture guardrail: use LiveKit for the interactive stage (host, co-hosts/guests and low-latency interactions); use Cloudflare Stream HLS/broadcast delivery for the mass audience. Do not route all 150 broadcast viewers through LiveKit by default. If the product instead uses LiveKit for every viewer, re-estimate participant minutes and downstream GB at the measured bitrate before enabling that topology.

### 5.3 Exchange-rate and price caveats

USD/KES is variable. The planning conversion uses KES 130/USD, close to the CBK indicative KES 129.98/USD published for 17 November 2026. Actual card, bank or provider settlement rates and fees may differ. Recalculate before payment or contract commitment.

Cloudflare and LiveKit prices are published list prices, not negotiated quotes. Taxes, FX spreads, additional storage, recording, transformations, LiveKit overages, support tiers and provider-specific add-ons may change the invoice.

## 6. AWS and managed infrastructure validation

AWS RDS and Fargate are consumption-priced. RDS charges vary by Region, DB class, deployment mode, storage, I/O, backup and data transfer; Fargate varies by vCPU, memory, OS/CPU architecture, ephemeral storage and runtime. Therefore the KES 180,000 three-month beta allowance is not represented as a verified AWS quote.

Before production spend approval, generate and retain an AWS Pricing Calculator estimate for Africa (Cape Town, af-south-1) and a comparison region if the approved architecture gate requires it. The estimate must explicitly include:

- Two physically separate managed PostgreSQL deployments/clusters: Core and Finance.
- ECS Fargate task count, vCPU, memory, architecture and hours for Core API, realtime gateway, media control plane and workers.
- Managed cache, managed NATS JetStream or the selected contracted broker plan.
- Load balancer, NAT/networking, public IPv4, inter-AZ and internet egress.
- Database storage, backup retention, snapshots, PITR and restore-test environment.
- Secrets, logs, metrics, traces and retention.
- Dev/staging/prod separation and any minimum monthly charges.

Keep the Finance database and credentials separate. Do not combine the two databases to make the budget fit. If the calculator estimate exceeds the allocation, reduce nonessential always-on capacity, use controlled beta hours where safe, negotiate credits, or reduce pilot scope; do not remove finance isolation, backups, monitoring or security controls.

## 7. Delivery schedule and cash-release plan

| Period | Delivery focus | Cash release |
|---|---|---:|
| Months 1–2 | Architecture, repo foundations, CI/CD, auth, account/profile/org, database boundaries and mobile shell | KES 500,000 |
| Months 3–4 | Social, content, feed, messaging, creator basics, catalogue and commerce foundations | KES 550,000 |
| Months 5–6 | M-Pesa, Finance ledger, orders, livestreaming, moderation, integration and release candidate | KES 550,000 |
| Month 7 | Closed beta, selected creators/merchants, payment and operational validation | KES 200,000 |
| Month 8 | Expanded live beta, capacity tests, defect fixes and security validation | KES 200,000 |
| Month 9 | Final beta, acceptance, controlled public-pilot release | KES 250,000 |
| Central contingency (held separately) | Approved, evidenced variance only | KES 250,000 |
| **Total** | | **KES 2,500,000** |

Cash releases are ceilings against accepted milestones, not automatic payments. The founder lead-developer allocation is paid against an agreed nine-month work/deliverable schedule and is already included in the professional-fee budget.

## 8. Nine-month delivery gates

### Months 1–2 — Foundation
- Go modular Core foundation and React Native/Expo application shell.
- OIDC, account/profile/organisation and resource authorisation.
- Separate Core and Finance managed database deployments.
- CI/CD, migrations, secrets, observability and dev/staging/prod boundaries.
- Threat model, data inventory and payment-provider onboarding plan.

### Months 3–4 — Core product
- Social graph, feed, posts, short video and engagement.
- Messaging, notifications, reporting, blocking and moderation basics.
- Creator publishing foundations.
- Seller profiles, catalogue, variants and inventory foundations.

### Months 5–6 — Integration and release candidate
- Cart, orders and M-Pesa provider integration.
- Finance ledger, payment callbacks, idempotency, reconciliation and payout protections.
- LiveKit interactive stage and Cloudflare Stream broadcast/VOD integration.
- Live-shopping product pinning.
- End-to-end, security, restore and media cost tests.
- Development exit requires a release candidate; unresolved critical/high finance or security defects block beta.

### Month 7 — Closed beta
- Invite-only cohort; staged creator and merchant onboarding.
- Validate payment state transitions, refunds, duplicate callbacks, messaging and media workflows.
- Daily reconciliation and support review.

### Month 8 — Expanded live beta
- Expand cohort only after month-7 gates pass.
- Test 150 concurrent broadcast viewers and 2,500 monthly viewer-hours as controlled workload objectives.
- Review cost per viewer-hour, stream startup/rebuffering, API latency, DB saturation and support workload.

### Month 9 — Final beta and public pilot
- Resolve critical/high defects and complete security remediation.
- Restore and incident-response drills.
- Confirm payment-provider production approval and operational readiness.
- Open controlled public registration only after acceptance gates pass.

## 9. Capacity limits and automatic controls

| Metric | Pilot target / ceiling |
|---|---:|
| Registered users | 10,000 |
| Monthly active users (planning target) | 2,500 |
| Monthly video viewing | 2,500 viewer-hours |
| Peak concurrent broadcast viewers | 150 |
| Initial geography | Kenya |

These are not simultaneous guarantees. Monthly viewing volume and peak concurrency are separate dimensions. The test must include realistic session duration, bitrate, device mix and network conditions.

Required controls:
- Registration cohorts and the ability to pause invitations.
- Per-account upload duration and storage quotas.
- Stream delivery and storage alerts at 50%, 75%, 90% and 100% of budget/quota.
- Media data-saver defaults, adaptive bitrate, autoplay controls and thumbnail-first feed.
- Cost alerts and daily provider usage review during beta.
- Separate load test for interactive LiveKit participants versus Cloudflare Stream broadcast viewers.
- Kill switches for live creation, recording, uploads and nonessential background work.
- No automatic scale-up that can create unbounded spend; alert and require an operator-approved budget increase.

## 10. Capacity benchmark and release gates

The 10,000 registered / 2,500 MAU / 2,500 viewer-hours / 150 concurrent baseline must be demonstrated on the exact release candidate and paid provider tiers.

Minimum beta gate:
- 150 concurrent broadcast viewers sustained for an agreed soak window, with join success, startup time, rebuffering, bitrate and cost recorded.
- 2,500 viewer-hours/month modelled and metered; reconcile billed minutes to product analytics.
- Separate interactive-room tests for 2, 5 and 10 participants; do not infer this capacity from the broadcast viewer test.
- Core API, messaging and checkout load tests with 2,500 MAU workload assumptions and burst traffic.
- Finance invariants, idempotency, callback replay, reconciliation and tenant isolation pass.
- No unresolved critical/high security or financial correctness defects.
- Restore, incident response, moderation, support and cost-alert procedures are operational.

Scale beyond this baseline only after a new capacity and cost review. The earlier 100,000-account / 25,000-viewer-hour / 500-concurrent ambition remains a later growth objective and is not included in this KES 2.5M pilot commitment.

## 11. Assumptions, exclusions and approval controls

- KES 2.5M is the complete proposed cash envelope, including paid founder lead-development fees.
- The KES 250,000 contingency is unspent unless formally approved.
- No founder salary is assumed beyond the stated KES 400,000 lead-developer allocation.
- Advanced battles/tournaments, sophisticated recommendation models, complex memberships, automated affiliate campaigns and multi-region deployment remain deferred unless scope is explicitly traded.
- Payment provider transaction fees, settlement charges, taxes and refunds must be separately identified once commercial terms are received; do not silently absorb unknown fees into engineering estimates.
- App-store fees, legal registration/filing fees and any third-party minimums must be verified before commitment.
- No supplier quote is implied by a planning allocation. Maintain invoices, calculator exports, contracts and actual usage reports as the financial evidence pack.

## 12. Source references

- Cloudflare Stream pricing and billing dimensions: https://developers.cloudflare.com/stream/pricing/
- LiveKit Cloud plans and included WebRTC/bandwidth quotas: https://livekit.com/pricing
- AWS RDS for PostgreSQL pricing and calculator: https://aws.amazon.com/rds/postgresql/pricing/
- AWS Fargate pricing and calculator: https://aws.amazon.com/fargate/pricing/
- Central Bank of Kenya indicative rates: https://www.centralbank.go.ke/rates/forex-exchange-rates/

