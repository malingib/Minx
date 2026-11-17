# Financial Architecture Specification

## Boundary and launch rules
Finance is a separately deployed Go service with its own managed PostgreSQL cluster, network policy, credentials, migration role, backup policy and audit. Core services cannot write to the finance database.

At launch, Minx does not issue general-purpose stored value, enable customer-to-customer transfers or sell transferable/cash-redeemable virtual coins. Collections and payouts use appropriately authorised providers. Creator and merchant earnings are contractual payables, not customer deposits.

## Subdomains
Payment intents and attempts; provider adapters; double-entry ledger; merchant payables; creator payables; refunds; payout requests/attempts; reconciliation; financial holds/risk; audit.

## Ledger invariants
- Debits equal credits for each journal and currency.
- Amounts are integer minor units; floating point is prohibited.
- Posted journals are immutable. Corrections use reversal and replacement journals.
- Every journal references a source transaction, actor/system, legal entity, idempotency key and correlation ID.
- Cross-currency movement requires explicit FX journal and rate evidence.
- Available payable is computed from posted journals and holds; no client can mutate balances.
- Balance snapshots are rebuildable projections, not the authoritative record.

## Core records
legal_entities, financial_accounts, account_owners, payment_intents, payment_attempts, provider_transactions, provider_callbacks, journal_entries, journal_lines, merchant_payables, creator_payables, payout_requests, payout_attempts, refunds, reconciliation_runs, reconciliation_items, financial_holds, idempotency_records, outbox_events, audit_records.

## Payment state machine
CREATED → SUBMITTING → PENDING → SUCCEEDED or FAILED. Operational states include UNKNOWN, EXPIRED, REQUIRES_REVIEW and REVERSED. Timeout is not failure. Contradictory provider results create an exception/reversal workflow; they must not silently overwrite posted history.

## Idempotency
Persist actor, operation, key, canonical payload hash, result resource, state and response. Same key and same payload returns the original result. Same key with a different payload returns conflict. Concurrent requests produce one money movement. Provider references have independent unique constraints.

## Provider adapter
Initial Kenya rail: M-Pesa through an approved Safaricom/Daraja or licensed PSP arrangement. Adapter capabilities include STK Push, supported collection flows, authorised payouts, status query, supported reversal and settlement ingestion. Capabilities must be confirmed against the actual provider contract. Callback verification is documented per provider; persist bounded evidence securely, deduplicate and match expected amount, currency and reference before posting.

## Payout lifecycle
REQUESTED → VALIDATING → RESERVED → SUBMITTED → COMPLETED, with REJECTED, FAILED, UNKNOWN, ON_HOLD and REVERSED branches. Reserve funds atomically. On timeout query provider and reconcile; never blindly resubmit.

## Reconciliation
Three-way match finance records, provider transaction/settlement data and commercial order/creator records. Exceptions include missing callback, success without journal, duplicate reference, amount/currency mismatch, fee mismatch, settlement gap, payout mismatch and refund mismatch. Every exception has owner, severity, evidence and approved resolution.

## Release blockers
No unbalanced journal, duplicate collection/payout, lost confirmed payment, cross-tenant access, silent journal mutation, payout from unavailable funds or hidden reconciliation mismatch. Kenyan legal review is mandatory before production money movement, including PSP scope, marketplace settlement, KYC/KYB, AML, gifts, tax and consumer obligations.
