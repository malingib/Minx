# Media and Realtime Specification

## Provider routing
- LiveKit: interactive calls, co-host rooms, battles, speaker stages and low-latency creator interaction.
- Cloudflare Stream: broadcast delivery, VOD, recordings and adaptive playback.
- Go Media Orchestrator: session lifecycle, provider selection, role/permission policy, short-lived tokens, recording policy, quotas, webhooks and takedowns.
- R2: auxiliary non-sensitive images, thumbnails and attachments; sensitive evidence goes to approved regional storage.

## Session classes
INTERACTIVE, BROADCAST, HYBRID, CALL and BATTLE. Interactive sessions use LiveKit. Broadcast sessions use Stream. Hybrid sessions publish an interactive stage and relay it to the broadcast audience through a supported egress/ingest path. The orchestrator owns routing.

Initial operational trigger for considering broadcast fanout: 500 passive concurrent viewers. This is a configurable starting threshold, not a capacity claim; tune from load and cost tests.

## Lifecycle
SCHEDULED → CREATED → LIVE → ENDING → ENDED; exceptional states: FAILED, SUSPENDED, TERMINATED, REVIEW. Webhooks are validated according to provider capability, deduplicated and reconciled against authoritative session state.

## Permissions
Short-lived, room-scoped participant tokens. Server controls room, participant, role, track and publish permissions. No provider secrets in clients. Host transfer/co-host elevation is authorised server-side. Moderation termination blocks further token issuance and removes active participants.

## Live shopping
Product pin is owned by Minx, not the media provider. Pin events include session, seller, product/variant, price snapshot/version and event sequence. Checkout uses ordinary commerce and finance state machines; the app preserves playback context. A client success screen never establishes payment success.

## Battles
Battle ruleset and scoring window are versioned. Competition points and financial gifts are separate aggregates. Apply rate limits, anti-bot checks, replay protection and suspicious engagement detection. Admin pause/end actions are audited.

## Messaging and presence
Messages are durable in PostgreSQL; WebSockets are transport only. Conversation cursors support replay after reconnect. Delivery/read receipts are durable and deduplicated. Typing and presence are ephemeral Redis state with TTL. Attachments use signed URLs and scanning. E2EE is not in launch scope.

## Low bandwidth
Adaptive bitrate/simulcast, explicit autoplay and data saver, thumbnail-first feed, resumable uploads, bounded retries, network-aware prefetch, low-resolution/audio fallback, offline drafts and state preservation across network transitions.

## Benchmarks
Test 1, 10, 100, 500, 1,000 and 10,000 viewers where service plan permits. Measure join success, glass-to-glass latency, startup, rebuffer, packet loss, reconnect, device CPU/battery, data consumption and cost per viewer-minute. Test host loss, network change, late join, recording and moderation termination. Initial targets for suitable networks: interactive latency below 500 ms, broadcast below 2 s, playback start p95 below 2.5 s, rebuffer below 1% of viewing time. These are test targets, not guarantees.
