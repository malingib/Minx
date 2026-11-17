# Product Requirements

## Product definition
Minx is an independent mobile-first social entertainment and commerce platform, launching in Kenya and expanding across Africa. Its core proposition is creator-first entertainment natively connected to shopping, communication, communities and monetisation.

## Product pillars
- Social: profiles, posts, short video, feed, discovery, comments, reactions, saves and sharing.
- Live: live sessions, co-hosting, audience participation, battles and live shopping.
- Market: creator and merchant shops, catalogue, variants, inventory, cart, checkout, orders, refunds and fulfilment.
- Connect: direct/group messaging, attachments, delivery/read state, presence and calls.
- Wallet/Earnings: payment history and creator/merchant payable summaries; the UI must not imply unauthorised e-money issuance.
- Creator Studio: publishing, content management, analytics, shop tools and monetisation settings.
- Communities: groups, roles, events, discussion and moderation.
- Business: organisations, staff roles, products, orders, fulfilment, settlement reports and integrations.

## Identity model
Account is the authentication/recovery boundary. Profile is a public persona. Organisation is a legal/business/creator entity with members and roles. One account may own/manage multiple profiles and organisations. A profile is not an organisation and cannot own legal settlement obligations. Mode switching changes presentation only; authorisation is always server-side.

## Pilot scope
Registration/recovery; account/profile/organisation foundations; follow graph; feed/posts/video; creator publishing; live and live-shopping pilot; merchant catalogues/cart/checkout/orders; licensed-partner payment collection; payable reporting; basic messaging/notifications; moderation/reporting/support; analytics and observability.

## Expansion scope
Battles, gifts, memberships, affiliate, communities, group calls, regional payment rails, currencies, languages, events, advertising and partner integrations. Each monetisation feature receives its own legal, accounting, fraud and state-machine design before activation.

## Non-functional requirements
Low-bandwidth operation; no duplicate money movement; strong organisation isolation; auditable privileged actions; degraded operation during provider outages; accessible/localised UI; currency-aware formatting; data export/deletion/retention; business and system instrumentation.

## Success metrics
DAU/MAU; active creators and retention; video start/rebuffer; live sessions/concurrent viewers/viewer-minutes; live product-click and purchase conversion; GMV/repeat purchase/refunds; creator earnings paid; payment success/reconciliation exceptions; message latency; moderation response; cost per active user and viewer-minute.
