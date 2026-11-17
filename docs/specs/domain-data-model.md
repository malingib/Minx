# Domain and Data Model

## Ownership map
Identity owns accounts/auth-subject mapping and device sessions. Profiles owns handles and profile settings. Organisations owns legal/business entities, memberships and roles. Social owns follows/blocks/mutes. Content owns posts/comments/reactions/shares/saves. Media owns asset metadata and renditions. Feed owns ranking projections only. Live owns session and participant metadata. Messaging owns conversations/messages/receipts. Communities owns groups and memberships. Commerce owns shops/products/variants/inventory/carts/orders/fulfilment. Creator economy owns earning rules and attribution. Finance owns payment state, journals, payables, payouts and reconciliation. Trust & Safety owns reports/cases/evidence/enforcement/appeals. Notifications owns notification records and delivery attempts.

## Core entities
- accounts(id, auth_subject, status, created_at, updated_at)
- profiles(id, account_id, handle, display_name, profile_type, visibility, status)
- organisations(id, legal_name, display_name, organisation_type, country_code, status)
- organisation_memberships(id, organisation_id, account_id, role_id, status)
- profile_operators(profile_id, account_id, permission_set, status)
- posts, post_media, comments, reactions, shares, saves
- media_assets, media_renditions, media_uploads
- shops, shop_members, products, product_variants, inventory_items, inventory_reservations
- carts, cart_items, orders, order_items, order_status_events, fulfilments, shipments
- conversations, conversation_members, messages, message_attachments, delivery_receipts, read_receipts

## Rules
Use immutable UUID/ULID IDs. Handles and phone numbers are mutable attributes, never primary keys. Snapshot item title, variant, price, tax and fees at checkout so later catalogue changes cannot rewrite historical orders. Media binaries live in object/video storage; metadata remains in PostgreSQL. Message bodies/history remain durable in PostgreSQL; sockets are transient.

## Tenant isolation
Organisation-owned rows carry organisation_id. Every query enforces actor/resource/action/scope. RLS is defence in depth, not sole protection; privileged owners can bypass. Test cross-organisation IDOR, role escalation and admin/support impersonation. Sensitive admin access is time-limited, reason-coded and audited.

## Classification
Public; private; sensitive (identity/contact/evidence); financial; security-critical; analytical. Define purpose, retention, legal basis, access and deletion/export rules for every class. Statutory financial records are restricted and retained under documented legal basis rather than silently deleted.
