# API and Event Contracts

## API conventions
External API is HTTPS REST/JSON under /v1, documented with OpenAPI. OIDC tokens are validated for issuer, audience, expiry and scopes. Every request performs resource-level authorisation. Use opaque cursors for mutable lists, RFC3339 UTC timestamps and integer minor units for money. Mutations with duplicate side effects require idempotency keys. Return stable error codes and correlation IDs; never expose stack traces, secrets or raw database errors.

## Resource groups
/accounts, /profiles, /organisations, /social, /posts, /feed, /discovery, /live-sessions, /conversations, /messages, /communities, /shops, /products, /carts, /orders, /creator, /finance, /notifications and /reports.

## Checkout contract
Core validates price/stock/delivery and creates order AWAITING_PAYMENT with immutable snapshot. Finance creates an idempotent payment intent and submits to provider. Verified provider outcome is durably stored; finance posts journal and emits event. Core consumes event idempotently and updates order. Fulfilment starts only after authoritative confirmation.

## Event envelope
Each event contains event_id, event_type, schema_version, occurred_at, producer, aggregate_type/id, correlation_id, causation_id and payload. Initial catalogue: identity.account.created.v1; profile.published.v1; social.follow.created.v1; content.published.v1; media.asset.ready.v1; live.session.started/ended.v1; message.created.v1; commerce.order.created/paid/fulfilment_updated.v1; finance.payment.confirmed/failed/reversed.v1; finance.refund.confirmed.v1; finance.creator_payable.created.v1; finance.merchant_payable.created.v1; finance.payout.completed.v1; moderation.action.applied.v1.

## Delivery guarantees
Outbox row commits with aggregate transaction. Publisher retries until broker acknowledgement. Consumer inbox/deduplication state commits atomically with side effects. Acknowledge after durable processing. Poison events enter dead-letter stream with alert and replay procedure. Schemas are versioned; breaking changes create a new version. No exactly-once guarantee is assumed.

## WebSockets
Authenticated WSS supports message.new, message.receipt, presence.changed, conversation.updated, live.reaction, live.product_pin.updated and notification.created. Reconnect uses last durable cursor. Durable missed messages are replayed from source of truth; typing/presence may be dropped.
