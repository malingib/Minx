# Minx — Investor Proposal

**Investor discussion draft | 18 November 2026**  
**Company:** Minx (independent venture)  
**Initial market:** Kenya, with Africa-first expansion potential

## 1. Executive summary

Minx is a mobile-first, creator-first social entertainment and commerce platform connecting short-form content, livestreaming, communities, messaging, creator monetisation and merchant commerce. Its signature journey allows a viewer to discover a product during a creator livestream, select a variant, check out using M-Pesa, receive an order confirmation and follow seller fulfilment without leaving the live experience.

The programme is structured as six months of product development followed by a three-month controlled beta, then a full commercial launch and three years of commercial operating development. The initial controlled beta targets are 10,000 registered users, 2,500 monthly active users, 2,500 monthly video-viewing hours and 150 peak concurrent livestream viewers. These are goals, not proven results or capacity guarantees.

The current financial model produces a preliminary Base-case funding envelope of **KES 29.70 million**. This is a planning estimate and must be validated against monthly cash timing, provider quotes, tax and regulatory advice before being presented as a firm funding ask.

## 2. Problem and opportunity

Creators often build audiences on social platforms but monetise through disconnected tools. Merchants need trusted creator-led discovery that can translate into measurable orders. Consumers need local content, familiar payments, lower-friction checkout and experiences that work under variable connectivity.

Minx is designed around the combined social, live, community and commerce loop, beginning with a controlled Kenya launch. Market size and demand are not asserted as verified in this draft; they require primary customer discovery, competitor analysis and category-level market sizing.

## 3. Product

Minx comprises eight pillars:
1. **Social:** profiles, short video, posts, feed and discovery.
2. **Live:** creator livestreams, co-hosting, interactive sessions and live shopping.
3. **Market:** merchant catalogues, inventory, carts, checkout, orders and fulfilment.
4. **Connect:** persistent messaging, presence, calls and notifications.
5. **Wallet and earnings:** payment history and contractual creator/merchant payable summaries—not an unauthorised stored-value wallet.
6. **Creator Studio:** publishing, analytics, monetisation settings and shop tools.
7. **Communities:** groups, membership, discussion and events.
8. **Business:** merchant organisations, staff roles, fulfilment and reporting.

The launch will not include a general-purpose customer stored-value wallet, customer-to-customer transfers or transferable/cash-redeemable coins. Creator and merchant earnings are contractual payables settled through appropriately licensed payment partners.

## 4. Differentiation

- Commerce is integrated into the live viewing journey rather than sent to a disconnected destination.
- M-Pesa-native payments and local-first product decisions.
- Creator and merchant tools share discovery and audience infrastructure while financial operations remain isolated.
- Low-bandwidth controls, adaptive media, resumable uploads and network-aware experiences.
- Modular Go architecture with separately deployed critical runtimes and a hard Finance boundary.

These are product and architecture choices, not yet proven competitive advantages. Validation requires user research, creator/merchant pilots, retention and conversion data.

## 5. Technology and controls

The approved baseline is React Native + Expo mobile apps, Go 1.27.4 modular Core API, separately deployed Finance service and database, managed PostgreSQL, NATS JetStream, Redis-compatible ephemeral cache, LiveKit interactive RTC, Cloudflare Stream broadcast/VOD, managed AWS foundation and Cloudflare edge/media. AWS Cape Town service availability, legal/data-transfer position, provider prices and performance must be evidenced before production commitments.

Finance uses double-entry accounting, immutable posted journals, integer minor units, idempotent operations, verified provider webhooks and reconciliation. Orders, payments, creator payables and inventory are separate concepts. No exactly-once event delivery assumption is made.

## 6. Delivery plan

| Stage | Timing | Focus | Gate |
|---|---|---|---|
| Product build | Months 1–6 | Core mobile journeys, identity, social, media, commerce, finance, security and operations foundation | End-to-end acceptance, CI, QA/security evidence and partner sandbox tests |
| Controlled beta | Months 7–9 | Invite cohorts, creators, merchants, payments, moderation, support and unit-economics measurement | KPI, reliability, trust/safety and spend review |
| Commercial launch | Following beta | Public launch, acquisition, creator/merchant activation and support scale-up | Compliance, app release, partner and operational readiness |
| Commercial years 1–3 | 36 months after beta | Retention, monetisation, geographic and category expansion | Quarterly board-approved scale gates |

## 7. Initial programme budget

### Build and controlled beta — KES 2.45M

| Item | KES |
|---|---:|
| Go contractor | 300,000 |
| React Native developer | 250,000 |
| UX/UI | 100,000 |
| QA/security | 150,000 |
| DevOps | 100,000 |
| Legal/payment integration | 100,000 |
| Founder/lead developer remuneration (paid) | 400,000 |
| Development infrastructure/tools | 150,000 |
| Three-month controlled beta | 650,000 |
| Central contingency | 250,000 |
| **Total** | **2,450,000** |

### Commercial launch allowance — KES 5.00M

| Item | KES |
|---|---:|
| Public launch marketing/acquisition | 2,500,000 |
| Creator and merchant activation | 1,000,000 |
| Launch support/moderation/operations setup | 750,000 |
| Legal/compliance/app release/launch contingency | 750,000 |
| **Total** | **5,000,000** |

The launch allowance is a planning assumption, not a vendor-backed campaign budget. It must be tied to channel-level acquisition, creator activation and conversion plans.

## 8. Three-year Base-case outlook

| KES millions | Year 1 | Year 2 | Year 3 |
|---|---:|---:|---:|
| Platform revenue | 6.204 | 33.720 | 135.720 |
| Operating costs | 18.437 | 38.906 | 89.896 |
| Operating surplus/(deficit) | (12.233) | (5.186) | 45.824 |

Revenue consists of marketplace commission, Minx share of creator monetisation, merchant subscriptions and advertising. Costs include payment processing, media delivery/storage, staff, marketing, core infrastructure, support/moderation and G&A/legal/compliance.

**Important reconciliation:** Year 1 marketplace GMV is KES 50.40M and the 5% commission is KES 2.52M. Year 2 marketplace GMV is KES 252.00M and commission is KES 12.60M; Year 3 marketplace GMV is KES 972.00M and commission is KES 48.60M. The detailed assumptions and formulas are documented in [the financial model](financial-model-36-month.md) and [monthly cash flow](monthly-base-case.csv).

## 9. Indicative funding bridge

| Component | KES |
|---|---:|
| Build and controlled beta | 2,450,000 |
| One-off commercial launch | 5,000,000 |
| Peak cumulative operating deficit (36-month monthly proxy) | 18,540,715 |
| 20% working-capital reserve | 3,708,143 |
| **Indicative planning envelope** | **29,698,858** |

This replaces the earlier KES 28.35M estimate, which understated the reconciled operating deficits. The KES 29.70M amount is not yet a firm raise target. A dated cash forecast must model hiring start dates, customer collections, provider settlement timing, supplier payment terms, tax, app-store charges, FX and contingency.

## 10. Funding and investor structure

The recommended approach is milestone-based financing. Do not set a valuation or equity percentage before diligence and negotiation. Potential structures include:
- Priced equity after valuation, cap-table and governance diligence.
- A convertible instrument where valuation is deferred, subject to Kenyan legal and tax review.
- Milestone tranches with clear release gates and investor protections.
- Strategic commercial partnerships for distribution, creators, merchants, payment rails or media expertise.

No investor return, dividend, valuation, equity allocation or exit date is promised in this proposal. Rights, reserved matters, board representation, information rights, dilution and exit provisions require negotiation and counsel review.

## 11. Milestones and release gates

A. Build foundation: approved architecture, signed scopes, IP assignment, company-controlled repositories/cloud accounts, CI and threat model.  
B. Feature-complete beta: end-to-end critical journeys, verified payment sandbox/webhooks and no unresolved critical security defects.  
C. Controlled beta: weekly cohort KPIs, reliability, moderation, spend and unit-economics reporting.  
D. Commercial launch: legal/compliance sign-off, app release, creator/merchant supply and support readiness.  
E. Scale extension: additional capital only after repeat usage, monetisation, CAC/payback and contribution margin are evidenced.

## 12. Key risks

- **Acquisition and retention:** cohort-led GTM, channel attribution and retention gates.
- **Monetisation:** validate buyer conversion, creator payer rates, merchant retention and ad demand.
- **Media cost:** adaptive bitrate, usage budgets, alerts, provider terms and delivery benchmarking.
- **Payments and regulation:** licensed partners, isolated ledger, reconciliation and Kenyan legal advice.
- **Trust and safety:** reporting, moderation, evidence, appeals, age policies and incident response.
- **Security and privacy:** least privilege, resource-level authorisation, data minimisation, testing and incident response.
- **Execution and runway:** milestone SOWs, independent QA/security, monthly cash forecast and tranche gates.

## 13. Governance and reporting

Maintain a board-approved budget, monthly management accounts, quarterly investor reporting, an auditable cap table, documented approval matrix and a risk register with owners and triggers. Investor information rights do not imply access to customer personal data. Company IP, source code, domains, cloud accounts and payment-provider relationships should remain under company-controlled accounts.

## 14. Due diligence

The investor room should include corporate and statutory records, cap table, shareholder agreements, IP assignments, product requirements, architecture and ADRs, repository/CI/cloud ownership evidence, security and privacy documentation, commercial pipeline, financial model and actuals, partner contracts and quotations, legal/tax advice, team agreements, beta KPIs and risk register. A folder index and checklist are committed in `docs/investors/`.

## 15. Disclaimer

This document is a discussion draft for planning purposes, not legal, tax, investment or regulatory advice, and not an offer to sell securities or a solicitation to invest. Forecasts are illustrative, sensitive to assumptions and not guarantees. Obtain qualified Kenyan legal, tax and regulatory advice before fundraising, launch or making public claims.
