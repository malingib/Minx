# Minx Mobile Screen Map

**Status:** Information architecture baseline — screen-level visual designs pending  
**Primary clients:** Android and iOS  
**Navigation model:** Five bottom destinations plus contextual creation

## 1. Global app structure

- App launch and session restore.
- Authentication: welcome, sign up/sign in, phone/email verification, recovery, terms/privacy.
- Onboarding: identity and username, profile photo/bio, interests, follow suggestions, just-in-time permissions, personal/creator/business intent.
- Main shell: Home, Live, Market, Inbox, You and global Create action.

## 2. Main navigation destinations

| Destination | Primary purpose | Main child screens |
|---|---|---|
| Home | Discover and consume social content | For You/Following, Stories, post detail, comments, search, topic results, creator profile |
| Live | Discover and participate in live experiences | Live discovery, category, scheduled live, live room, battle room, replay |
| Market | Discover and purchase products | Market home, category, search/results, shop, product detail, cart, checkout, orders |
| Inbox | Communicate and receive activity | Conversation list, message requests, conversation, group conversation, notifications |
| You | Manage identity and access tools | Profile, edit profile, profile switcher, saved items, settings, Creator Studio, Business Centre, Wallet |

A prominent Create action is available from the shell but is not a sixth destination. It opens Post, Short video, Go live, Product and Story actions according to permissions and account state.

## 3. Authentication and onboarding

| ID | Screen | Required states / notes |
|---|---|---|
| AUTH-01 | Welcome | New user, returning user |
| AUTH-02 | Create account | Validation, duplicate account, network failure |
| AUTH-03 | Sign in | Invalid credentials, locked/restricted account |
| AUTH-04 | Verify phone/email | Code sent, invalid, expired, resend cooldown |
| AUTH-05 | Recover account | Request, verification, completion |
| ONB-01 | Choose identity | Personal profile, creator intent, business intent |
| ONB-02 | Set username | Availability, reserved name, invalid format |
| ONB-03 | Profile basics | Photo, display name, bio, skip |
| ONB-04 | Select interests | Search, multi-select, skip |
| ONB-05 | Follow suggestions | Follow/unfollow, refresh, empty |
| ONB-06 | Permission education | Camera, mic, photos, notifications; request at point of use |
| ONB-07 | Completion | Personalised next step; no forced purchase |

## 4. Home / Social

| ID | Screen | Key capabilities |
|---|---|---|
| SOC-01 | Home feed | For You/Following, mixed posts and short videos, refresh, pagination |
| SOC-02 | Stories tray | Followed creators, unread state, create story entry |
| SOC-03 | Full-screen media viewer | Video/photo, captions, mute, quality, report |
| SOC-04 | Post detail | Full caption, media, comments, creator identity |
| SOC-05 | Comments | Threaded replies, likes, report, block, sort |
| SOC-06 | Create post | Text, image, video, audience, optional location, disclosure |
| SOC-07 | Create story | Media, text overlay, audience, expiry preview |
| SOC-08 | Search | People, posts, topics, shops and products |
| SOC-09 | Search results | Typed result tabs, filters, empty and error states |
| SOC-10 | Topic / hashtag | Related content and follow topic |
| SOC-11 | Share sheet | In-app message, copy link, OS share |
| SOC-12 | Saved items | Posts, videos, products, collections |
| SOC-13 | Creator profile | Follow, message, content, live status, shop entry |
| SOC-14 | Followers / following | Lists, search, remove/block where permitted |

## 5. Live

| ID | Screen | Key capabilities |
|---|---|---|
| LIV-01 | Live discovery | Live now, recommended, followed, categories |
| LIV-02 | Live schedule | Upcoming sessions, reminders, creator schedule |
| LIV-03 | Pre-live setup | Title, cover, audience, moderation, product selection, connectivity check |
| LIV-04 | Live room — viewer | Video, chat, reactions, host profile, product tray, gifts, report |
| LIV-05 | Live room — host | Start/end, mute, camera, co-host invite, pin product, moderation |
| LIV-06 | Co-host lobby | Invite, accept, reject, permissions, reconnect |
| LIV-07 | Battle setup | Opponent/team, rules, duration, eligibility |
| LIV-08 | Battle room | Score, timer, vote/judging; gifts shown separately from score |
| LIV-09 | Gift picker | Item, price/value disclosure, confirmation, limits |
| LIV-10 | Live product panel | Product, variants, stock, seller, price snapshot |
| LIV-11 | Live purchase confirmation | Order/payment status without losing session context |
| LIV-12 | Live ended | Summary, replay, follow, order access |
| LIV-13 | Live moderation | Mute, remove, block, report, slow mode, banned terms |
| LIV-14 | Replay / recording | Playback, creator controls, report/takedown |

## 6. Market and commerce

| ID | Screen | Key capabilities |
|---|---|---|
| MKT-01 | Market home | Categories, featured shops, creator picks, live shopping |
| MKT-02 | Category | Product list, sort, filters |
| MKT-03 | Product search | Query, suggestions, recent searches |
| MKT-04 | Search results | Filters, sort, seller and delivery context |
| MKT-05 | Shopfront | Seller identity, catalogue, reviews, policies |
| MKT-06 | Product detail | Media, description, variants, price, stock, delivery, reviews |
| MKT-07 | Cart | Quantities, remove, seller grouping, totals |
| MKT-08 | Checkout — delivery | Address/contact, delivery method, order review |
| MKT-09 | Checkout — payment | Supported payment partner methods, fees and amount due |
| MKT-10 | Payment processing | Pending/unknown/failed/succeeded; safe status refresh |
| MKT-11 | Order confirmation | Order reference, seller, next steps |
| MKT-12 | My orders | Active, completed, cancelled/refunded |
| MKT-13 | Order detail | Timeline, fulfilment, tracking, contact seller, support |
| MKT-14 | Review product | Rating, text/media, eligibility and moderation |
| MKT-15 | Returns / refund request | Reason, evidence, policy, status |

## 7. Connect / messaging

| ID | Screen | Key capabilities |
|---|---|---|
| CON-01 | Inbox | Conversations, unread, search, pinned |
| CON-02 | Message requests | Accept, decline, block, report |
| CON-03 | Direct conversation | Text, media, voice note, product share, receipts |
| CON-04 | Group conversation | Members, roles, invite, leave |
| CON-05 | Conversation details | Participants, shared media, mute, block/report |
| CON-06 | Attachment picker | Camera, gallery, file; permission and upload states |
| CON-07 | Voice note recording | Permission, recording, cancel, upload, retry |
| CON-08 | Notification centre | Social, live, commerce, security and account activity |
| CON-09 | Call screen | Voice/video call, mute, camera, network/reconnect, end |
| CON-10 | Call history | Missed, received, outgoing |

## 8. Wallet and financial experience

Wallet is a sensitive financial surface and must not be visually conflated with virtual gifts or promotional credits.

| ID | Screen | Key capabilities |
|---|---|---|
| WAL-01 | Wallet overview | Separate customer payment methods and creator/merchant earnings |
| WAL-02 | Payment methods | Add/remove supported provider method; no raw credential storage |
| WAL-03 | Transaction activity | Filter by type, date, status and currency |
| WAL-04 | Transaction detail | Amount, currency, reference, status, fees, receipt |
| WAL-05 | Creator earnings | Gross, fees, pending, available, held, paid out |
| WAL-06 | Earnings detail | Source breakdown, adjustments, settlement status |
| WAL-07 | Payout setup | Eligibility, identity checks, destination verification |
| WAL-08 | Request payout | Amount, fees, destination, confirmation |
| WAL-09 | Payout status | Requested, reserved, submitted, completed, unknown, failed, on hold |
| WAL-10 | Refund detail | Original payment, reason, status, next step |
| WAL-11 | Virtual items / gifts | Non-cash item inventory, purchase disclosure and usage rules |
| WAL-12 | Financial help | Dispute, payment issue, payout support, evidence |

No customer-to-customer transfers, general stored-value wallet or cash-redeemable transferable coins are assumed for launch. The final legal/provider model determines which screens and actions can be enabled.

## 9. Creator Studio

| ID | Screen | Key capabilities |
|---|---|---|
| CRE-01 | Creator overview | Performance snapshot, tasks, monetisation eligibility |
| CRE-02 | Content library | Draft, published, scheduled, removed |
| CRE-03 | Content analytics | Reach, views, retention, engagement |
| CRE-04 | Audience insights | Aggregate audience and follower trends |
| CRE-05 | Monetisation hub | Available programmes and eligibility |
| CRE-06 | Memberships | Tiers, benefits, subscriber list and status |
| CRE-07 | Paid content | Offers, access rules, purchases |
| CRE-08 | Affiliate centre | Products, links, attribution, commissions |
| CRE-09 | Brand campaigns | Applications, deliverables, disclosure |
| CRE-10 | Creator settings | Public profile, audience, content and monetisation preferences |

## 10. Communities

| ID | Screen | Key capabilities |
|---|---|---|
| COM-01 | Community discovery | Recommended, trending, categories |
| COM-02 | Community detail | About, rules, posts, members, join request |
| COM-03 | Community feed | Posts, events, polls, pinned content |
| COM-04 | Create community | Name, description, visibility, rules |
| COM-05 | Member management | Roles, invitations, bans, approvals |
| COM-06 | Community settings | Moderation, permissions, privacy |
| COM-07 | Events | Create, RSVP, reminder, event live link |

## 11. Business Centre

| ID | Screen | Key capabilities |
|---|---|---|
| BUS-01 | Business overview | Orders, sales, products, payouts, tasks |
| BUS-02 | Organisation setup | Legal/trading identity, owner, staff |
| BUS-03 | Verification | KYB requirements, evidence, review status |
| BUS-04 | Shop profile | Branding, policies, contact, delivery regions |
| BUS-05 | Product catalogue | Create/edit/archive, variants, media |
| BUS-06 | Inventory | Stock levels, adjustments, low-stock alerts |
| BUS-07 | Order management | New, accepted, processing, shipped, delivered, cancelled |
| BUS-08 | Order detail | Customer-safe information, fulfilment, communication |
| BUS-09 | Promotions | Discounts, bundles, campaign windows |
| BUS-10 | Business analytics | GMV, orders, conversion, repeat buyers |
| BUS-11 | Staff and roles | Invite, permissions, revoke |
| BUS-12 | Business settings | Payout destination, integrations, policies |

## 12. You, identity and settings

| ID | Screen | Key capabilities |
|---|---|---|
| YOU-01 | My profile | Active identity, content, shop/creator entry |
| YOU-02 | Edit profile | Photo, name, username, bio, links |
| YOU-03 | Profile switcher | Personal, creator and managed business identities |
| YOU-04 | Account settings | Login methods, phone/email, sessions |
| YOU-05 | Privacy settings | Audience, discovery, activity and messaging |
| YOU-06 | Security | Passkeys/2FA where supported, sessions, alerts |
| YOU-07 | Notifications settings | Category and channel preferences |
| YOU-08 | Data and permissions | Export, deletion request, app permissions |
| YOU-09 | Accessibility | Text size, captions, reduced motion, contrast |
| YOU-10 | Data saver and playback | Autoplay, media quality, download preferences |
| YOU-11 | Blocked / restricted accounts | Manage blocks and restrictions |
| YOU-12 | Help and support | Help, report issue, appeal, contact |
| YOU-13 | Legal and policies | Terms, privacy, community and commerce policies |

## 13. Cross-cutting trust and safety

Trust and safety actions are embedded in content, messaging, live, commerce and profile screens. Dedicated experiences include:

- Report content/account/message/live/product/order.
- Block, mute, restrict and hide.
- Report confirmation and case status.
- Age gate and age-appropriate product/feature restrictions.
- Creator/merchant identity verification status.
- Moderation decision notice and appeal.
- Account warning, suspension and reinstatement.
- Safety centre and community standards.
- Evidence upload with privacy notice and retention policy.

Moderator/admin tooling is a separate internal web application and is not part of consumer mobile navigation.

## 14. Global journeys to prototype first

1. **New user to first meaningful follow:** welcome → verification → profile → interests → follow suggestions → feed.
2. **Viewer to live shopping purchase:** feed/live discovery → live room → pinned product → variant → checkout → payment status → order confirmation → return to live.
3. **Creator to first live:** creator setup → permissions → live setup → preflight → go live → pin product → end session → analytics.
4. **Creator to earnings payout:** monetisation eligibility → earnings detail → payout setup → request → status → receipt/support.
5. **Merchant to first fulfilment:** business setup → verification → shop → product/stock → order received → fulfilment → settlement.
6. **Social message to commerce:** creator profile → message → shared product → product detail → checkout.
7. **Safety intervention:** content/live/message → report → confirmation → case status → decision/appeal.
8. **Identity switching:** You → switch profile → active profile confirmation → role-appropriate content/actions.

## 15. Prioritisation

### P0 — Pilot-critical
- Authentication and onboarding.
- Home feed, profiles, follow, create post/video.
- Live discovery, host/viewer live room, moderation essentials.
- Product discovery, shop/product detail, cart and checkout.
- Payment state and order confirmation.
- Orders and fulfilment basics.
- Inbox and direct messaging essentials.
- Creator earnings and payout status.
- Trust/safety reporting, blocking and account restrictions.
- Settings, privacy and data saver.

### P1 — Growth and monetisation expansion
- Stories, richer comments and sharing.
- Battles, gifts and leaderboards.
- Creator Studio analytics and monetisation tools.
- Memberships and paid content.
- Communities and events.
- Merchant analytics, promotions and staff roles.
- Voice/video calls.
- Reviews, returns and refunds self-service.

### P2 — Ecosystem expansion
- Affiliate marketplace and brand campaigns.
- Advanced creator tools and editing.
- Regional payment methods and multi-currency expansion.
- Advanced recommendations and discovery.
- Business integrations and richer commerce operations.

## 16. Screen design acceptance checklist

Each screen specification must include:
- Purpose, primary user and entry points.
- Main action and secondary actions.
- Layout hierarchy and component inventory.
- Data required and ownership domain.
- Loading, empty, error, offline, restricted and success states.
- Accessibility labels, focus order and text scaling.
- Analytics events and privacy classification.
- Navigation/back behavior and state preservation.
- Small-screen and large-screen behavior.
- Android and iOS platform-specific behavior.

The screen map is a complete planning inventory, not a commitment to ship every listed screen in the pilot. Detailed screen flows and visual prototypes must validate P0 journeys before implementation.
