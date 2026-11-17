# Performance Benchmarks and Capacity Validation

**Status:** Proposed validation plan  
**Applies to:** Minx core API, realtime, finance, media control plane, workers, mobile clients and managed infrastructure  
**Decision rule:** These are acceptance hypotheses for testing, not claims of achieved performance or vendor guarantees. Record the tested commit, region, plan/tier, configuration, dataset, client mix and raw results for every run.

## 1. Objectives

Validate user experience, correctness under concurrency, resilience, operational capacity and unit economics before public launch. A benchmark is not accepted from a single happy-path run. Run warm-up, steady-state, burst, soak and failure-injection phases; repeat material tests at least three times and report p50, p95 and p99.

## 2. Initial service-level objectives (provisional)

| Surface | Initial target | Measurement |
|---|---:|---|
| Core read API | p95 <= 300 ms; p99 <= 800 ms | Server-side latency, excluding client network |
| Core write API | p95 <= 500 ms; p99 <= 1 s | Successful validated writes |
| Feed first useful render | p75 <= 1.5 s on defined 4G profile | Client telemetry, cold and warm cache |
| Message send acknowledgement | p95 <= 500 ms within region | Durable acceptance, not just socket write |
| Message delivery | p95 <= 1 s within region | Recipient delivery acknowledgement |
| WebSocket reconnect | p95 <= 3 s after network restoration | Client-observed |
| Interactive media glass-to-glass | p95 <= 500 ms | Same-region controlled test; confirm measurement method |
| Broadcast playback start | p95 <= 2.5 s | Join request to first decoded frame |
| Broadcast rebuffering | < 1% of watched time | Session-weighted and viewer-weighted reports |
| Checkout orchestration | p95 <= 1 s excluding PSP processing | Minx-side request handling |
| Finance posting | p95 <= 500 ms after verified provider outcome | Atomic journal commit |
| Worker event lag | p95 <= 5 s at steady state | occurred_at to consumer completion |

Targets must be segmented by geography, device class, network profile and provider. Do not average away poor-tail performance or PSP latency.

## 3. Workload model

Use synthetic, privacy-safe data with production-shaped cardinality and skew. Publish the exact workload profile with each test.

- Accounts: 100k baseline, then 1M and 10M scale projections.
- Profiles: model multiple profiles per account and creator/merchant organisations.
- Social graph: skewed creator-follow distributions; include high-fanout accounts.
- Feed: mixed read/write workload, pagination, ranking projection misses and cold caches.
- Messaging: one-to-one and group conversations, durable history, receipts and reconnect replay.
- Live: small interactive rooms plus broadcast audiences; test chat, reactions, pinned products and moderator actions separately.
- Commerce: catalogue browsing, inventory contention, cart updates, checkout, callback duplication and order status.
- Finance: concurrent payment intents, duplicate callbacks, reversals, refunds and payout state transitions.
- Workers: outbox backlog, retries, poison messages, consumer restarts and replay.

Use realistic payload sizes, media bitrates, geographic distribution and diurnal traffic patterns. Document assumptions; never turn a capacity estimate into a forecast without evidence.

## 4. Test matrix

### 4.1 Core API and PostgreSQL

Run 50, 100, 250, 500 and 1,000 concurrent virtual users, increasing only while error rate and database health remain within guardrails.

Measure throughput, latency percentiles, saturation, pool wait, query plans, lock waits, deadlocks, replication lag, cache hit ratio and storage growth. Include hot-row contention (popular product inventory), feed fan-out, high-follower publishing, pagination, tenant isolation and migration under load.

Pass conditions:
- No cross-account, cross-profile or cross-organisation data exposure.
- No lost acknowledged writes or duplicate business effects.
- Error rate < 0.1% for steady-state non-provider requests.
- No unbounded connection growth, lock pile-ups or uncontrolled queue growth.
- Database remains below agreed CPU, IOPS, connection and storage headroom thresholds.

### 4.2 Realtime and messaging

Test 100, 500, 1,000, 5,000 and 10,000 concurrent sockets per environment as capacity permits. Test fanout to 1, 10, 100 and 1,000 recipients, group history replay, duplicate delivery, reconnect storms, token expiry and gateway rolling restart.

Pass conditions:
- Acknowledged messages are durably stored before success is returned.
- Reconnect resumes from a cursor without gaps; duplicates are safely deduplicated.
- Presence and typing expire after TTL and are never treated as durable state.
- Backpressure is bounded and visible; slow consumers do not exhaust memory.
- Authentication and resource authorisation are rechecked on sensitive actions.

### 4.3 Media and live shopping

Test interactive rooms at 2, 5, 10 and 25 participants; broadcast audiences at 100, 500, 1,000, 5,000 and 10,000 viewers where provider plan and budget permit. Test hybrid interactive-stage/broadcast-relay handoff, weak uplink, packet loss, mobile background/foreground, device rotation, token expiry, stream interruption and provider webhook replay.

Capture join success, time to first frame, glass-to-glass latency, bitrate adaptation, packet loss, jitter, reconnect success, rebuffering, battery, data consumed per viewer-minute and total cost per viewer-minute.

The proposed 500-viewer relay threshold is a configurable hypothesis. Choose the actual transition threshold only after comparing quality and cost across room sizes and geographies. Verify recording, retention, takedown and moderation paths during the test.

### 4.4 Finance and commerce correctness

Performance is secondary to correctness. Run concurrency and fault tests against a provider simulator before live credentials are used.

Required adversarial cases:
- Same idempotency key submitted concurrently with identical and conflicting payloads.
- Duplicate, delayed, reordered and invalid-signature provider callbacks.
- Provider timeout after request acceptance; outcome remains UNKNOWN until queried or reconciled.
- Worker crash before and after journal commit.
- Two buyers competing for the final inventory unit.
- Refund after partial fulfilment; reversal after settlement; payout retry after ambiguous timeout.
- Duplicate event delivery and replay after consumer deployment.

Pass conditions:
- Every posted journal balances by currency; amounts use integer minor units.
- Posted entries are immutable; corrections use reversal/replacement.
- One provider transaction cannot produce duplicate financial effects.
- Inventory cannot be oversold where reservation is required.
- No blind retry of an ambiguous payout.
- Order, payment, entitlement and payable state remain separate and reconcilable.
- Every exception is observable and has an operator workflow.

### 4.5 Workers and event broker

Test steady event rate, 2x and 5x bursts, 30-minute backlog recovery, consumer restart, broker node loss, redelivery, schema-version mismatch and poison-message quarantine. Verify outbox/inbox atomicity and idempotent effects.

Pass conditions:
- No event loss after a committed outbox transaction.
- Duplicate delivery does not duplicate side effects.
- Lag returns to baseline within the declared recovery window.
- Poison events are isolated, inspectable and replayable after correction.
- Retention and storage growth remain within the approved budget.

### 4.6 Mobile and low-bandwidth

Test low-end Android devices, current iOS devices and representative older OS versions. Network profiles: stable Wi-Fi, 4G, constrained 3G, high latency, 2–5% packet loss, intermittent connectivity and offline transitions.

Measure cold start, memory, CPU, battery, data per feed session, upload resume, crash-free sessions, image/video startup and accessibility interaction. Verify data saver, autoplay-off, thumbnail-first feed, resumable uploads, offline drafts and reduced background refresh.

## 5. Resilience and recovery drills

Run controlled failure injection in non-production:
- Kill one API task and one realtime gateway during traffic.
- Restart workers while the outbox has pending events.
- Simulate Redis loss; durable business flows must continue with degraded ephemeral features.
- Simulate broker unavailability; outbox backlog grows boundedly and drains after recovery.
- Simulate primary database failover and verify connection recovery.
- Restore core and finance databases from backup into isolated environments.
- Simulate media-provider outage and verify clear user-facing degraded states.
- Simulate PSP timeout and callback outage; reconcile without duplicate charge or payout.

Before production, agree and test RPO/RTO separately for core and finance. A successful backup job is not a restore test.

## 6. Cost and capacity model

Report cost in the units that map to product decisions:
- Core API: cost per 1,000 successful requests.
- Realtime: cost per 1,000 concurrent socket-hours and per delivered message.
- Media: cost per viewer-minute, creator-hour and recorded GB-month.
- Storage: cost per active user and per uploaded GB.
- Finance: cost per 1,000 payment attempts and reconciled transactions.
- Notifications: cost per 1,000 delivered notifications.
- Observability: ingestion and retention cost per active user.

For each test, report provider, region, service tier, committed spend, variable spend, egress, storage, peak usage and projected headroom. Include at least 30% operational headroom in initial capacity planning; revisit after measured traffic. No spend commitment until pricing, egress and minimum charges are confirmed.

## 7. Benchmark report template

Each run must record:
1. Run ID, date, tested commit SHA and environment.
2. Region, provider plan, service versions and configuration.
3. Dataset size, workload mix, client/device/network profile.
4. Warm-up, steady-state, burst, soak and fault-injection duration.
5. Throughput and p50/p95/p99 latency; errors and timeouts.
6. CPU, memory, DB connections, IOPS, locks, queue lag and network egress.
7. Correctness assertions and invariant-check results.
8. Cost breakdown and cost per unit.
9. Known limitations, regressions and raw artifact links.
10. Decision: PASS, FAIL or INCONCLUSIVE, with owner and remediation.

## 8. Promotion gates

- **Local/dev:** correctness and repeatable smoke tests.
- **Staging:** production-shaped workload, security tests, failure injection, restore drill and cost report.
- **Pilot:** capped users, explicit support/on-call, kill switches, alert thresholds and daily reconciliation.
- **Public launch:** all critical gates in external-gates.md and roadmap/acceptance.md evidenced; no unresolved critical/high security or financial correctness defects.

A benchmark passes only for the exact tested build and environment. Material changes to service versions, topology, database tier, media provider configuration or event semantics require targeted reruns.
