# Minx Finance Service

This is an independently deployable runtime and module boundary. It owns its own database credentials, migrations and financial data. It must never import Core persistence packages or share the Core database.

## Initial boundary
- Payment orchestration and provider adapters
- Double-entry ledger and immutable posted journals
- Merchant and creator contractual payables
- Refunds, payouts, reconciliation, risk and audit

## Non-negotiable controls
- Integer minor units; no floating-point money
- Currency on every monetary amount and journal
- Balanced debit/credit entries per journal and currency
- Immutable posted journals; corrections through reversals
- Idempotent API requests, callbacks, consumers and payouts
- Provider timeouts remain UNKNOWN until queried/reconciled
- No customer stored-value wallet or peer-to-peer transfers at launch

The service is intentionally a compiling boundary only at this stage. No payment provider or production database integration is represented as complete.
