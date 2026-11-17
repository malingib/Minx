# Minx System Design

## Runtime map
- **Core API (Go):** profiles, organisations, social graph, content metadata, feed, discovery, catalogues, inventory, carts, orders, communities and creator/merchant configuration.
- **Realtime gateway (Go):** authenticated WebSockets, presence, message fanout, receipts and live application events.
- **Finance (Go):** payment orchestration, adapters, ledger, payables, refunds, payouts, reconciliation, risk and financial audit. Separate database and credentials.
- **Media orchestrator (Go):** provider-neutral session lifecycle, participant permissions, short-lived tokens, recording, webhooks, quotas and moderation. Does not proxy media.
- **Workers (Go):** outbox publication, event consumers, notifications, fulfilment, media coordination and scheduled tasks.
- **Intelligence (Python, later):** ranking, recommendations, feature extraction and fraud models; no transactional ownership.

## Domain ownership
Identity; profiles; organisations; social graph; content; media metadata; feed; live; messaging; communities; catalogue; inventory; checkout; orders; fulfilment; creator economy; finance; notifications; trust and safety; search; analytics; platform operations. Each domain owns its writes and schema. Cross-domain access uses explicit interfaces or versioned events.

## Data topology
- Core managed PostgreSQL 18 cluster with domain-owned schemas and restricted roles.
- Separate finance PostgreSQL cluster.
- Managed Redis-compatible cache for presence, rate limits, ephemeral coordination and hot projections.
- NATS JetStream durable event streams.
- R2 for non-sensitive media; regional object storage for sensitive evidence and backups.
- PostgreSQL full-text search initially; OpenSearch later if relevance or scale requires it.
- ClickHouse Cloud for analytical workloads when justified.

## Transaction boundaries
A business transaction commits to its owning database and writes an outbox record in the same transaction. A publisher transfers the event to JetStream. Consumers persist inbox/deduplication state and their domain change atomically. No cross-database distributed transaction is assumed.

## Commerce/payment sequence
1. Core validates cart, stock, delivery and price.
2. Core creates order with immutable price/tax/fee snapshot and AWAITING_PAYMENT state.
3. Finance creates an idempotent payment intent and submits to licensed provider.
4. Provider callback/status query is verified, deduplicated and durably recorded.
5. Finance posts balanced journals and emits confirmed/failed event.
6. Core consumes event idempotently and updates order.
7. Fulfilment begins only after authoritative confirmation.

## Failure isolation
Provider outage must not stop browsing; broker outage accumulates outbox rows; push outage does not prevent in-app notification retrieval; search outage degrades discovery, not checkout; media outage must not corrupt session metadata; analytics outage never blocks finance or commerce writes.
