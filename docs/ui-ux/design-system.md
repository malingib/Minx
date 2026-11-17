# Minx UI/UX Design System

**Status:** Foundation specification — visual direction not yet selected  
**Primary surfaces:** Android and iOS (React Native + Expo)  
**Secondary surface:** Public marketing website (React + Vite); not a consumer web application

## 1. Design principles

1. **Creator-first, not creator-only.** Discovery and creation are prominent, while shopping, communication and community remain first-class.
2. **One platform, coherent modes.** Social, Live, Market, Connect, Wallet, Creator Studio, Communities and Business share navigation patterns without crowding one screen.
3. **Content before chrome.** Media and creator content receive visual priority. Persistent controls must not obscure video, captions, products or live interaction.
4. **Commerce without context loss.** A viewer can inspect a product, choose a variant and start checkout without losing the live session or feed position.
5. **Money must be unmistakable.** Distinguish available cash, pending earnings, reserved funds, refunds, promotional credits and non-cash virtual items. Never use one ambiguous balance.
6. **Progressive disclosure.** Expose common actions first; place advanced creator, merchant and financial controls in contextual subflows.
7. **Low-bandwidth by design.** Respect data saver, adaptive media quality, explicit autoplay preferences, resumable uploads and reduced-motion settings.
8. **Accessible by default.** Support screen readers, dynamic type, sufficient contrast, target sizing, captions and clear focus states.
9. **Trust is part of the interface.** Reporting, blocking, privacy, age-appropriate experiences, transaction status and appeals must be discoverable.
10. **Real states, not decorative mocks.** Every screen defines loading, empty, error, offline, permission-denied, success and partial-success states where applicable.

## 2. Experience architecture

### Primary navigation

Use a five-destination bottom navigation model:

- **Home** — personalised social feed and discovery entry.
- **Live** — live discovery, scheduled sessions and active rooms.
- **Market** — product discovery, shops, categories and orders.
- **Inbox** — direct messages, conversation requests and notifications entry.
- **You** — active profile, profile switcher, creator/merchant tools and settings.

A prominent, centrally placed **Create** action is available from the shell but is not a sixth destination. It opens a creation sheet with Post, Short video, Go live, Product and Story actions according to permissions and account state.

The active profile/mode is visible in the You area and relevant creation flows. Switching a profile changes the active identity and presentation only; server-side authorisation remains resource- and role-based.

### Contextual navigation

- Live room: back, minimise where safe, product tray, chat, gifts, participants and host controls.
- Product detail: return to originating feed/live/shop context.
- Checkout: secure, focused flow with persistent order summary and explicit payment state.
- Wallet: separate money and earnings areas; no misleading merged balance.
- Creator Studio and Business Centre: accessible from You and contextual prompts, not permanently added to the global tab bar.

## 3. Visual language — constraints before concept selection

The final palette, typography pairing, corner language, icon style and media treatment remain unselected pending visual exploration. Approved constraints:

- Distinctive and editorial, not a clone of TikTok, Instagram, Facebook or a generic commerce template.
- Avoid excessive gradients, glassmorphism, oversized rounded cards, decorative shadows and dashboard-like card grids.
- Prefer strong typographic hierarchy, disciplined spacing, intentional negative space and restrained surfaces.
- Use colour semantically for actions, status and financial meaning; never rely on colour alone.
- Use consistent icon weight and a coherent icon family. Do not use emoji as primary interface icons.
- Product imagery and creator media must be real content or deliberate placeholders, never random decorative stock.
- Evaluate dark and light themes as concepts. Support system theme preference even if one theme is the launch default.

### Visual concepts to compare

1. **Editorial Creator:** expressive typography, bold media framing, strong creator identity and restrained UI chrome.
2. **Immersive Live Commerce:** edge-to-edge media, contextual product overlays, clear purchase actions and compact live controls.
3. **Modern Social Utility:** balanced content density, calm surfaces, strong navigation clarity and accessible commerce patterns.

These are exploration directions, not approved final themes. Select one coherent direction after reviewing visual screens, not by combining unrelated elements.

## 4. Design tokens

Implement tokens as semantic variables, not hard-coded values in individual screens.

### Required token groups

- Colour: background, surface, surface-raised, text-primary, text-secondary, text-inverse, border, divider, accent, accent-foreground, success, warning, danger, info, focus and overlay.
- Typography: display, heading-1, heading-2, heading-3, body, body-small, label, caption and numeric.
- Spacing: 4-point base scale with named semantic steps.
- Radius: none, subtle, control, panel, sheet and pill.
- Elevation: restrained levels for menus, sheets and overlays; avoid shadow as the primary grouping mechanism.
- Motion: duration and easing tokens for instant, short, standard and emphasis transitions; obey reduced-motion.
- Sizing: minimum touch target 44×44 logical points as a design baseline; platform conventions may require larger.
- Layering: navigation, sticky controls, sheets, modal dialogs, toasts and live overlays.

### Semantic implementation rules

- Components consume semantic tokens only; no raw hex values or one-off spacing in feature screens.
- Light/dark themes map to the same semantic token names.
- Financial states use labels and icons in addition to semantic colour.
- Text scaling must not clip critical actions or financial amounts.
- Numeric money fields use tabular figures and explicit currency codes/symbols.

## 5. Component inventory

### Foundations
- Typography, icon wrapper, spacing primitives, divider, surface, safe-area container.
- Avatar, profile identity, verified/official indicator, creator badge.
- Button, icon button, segmented control, tabs, chip, tooltip, badge.
- Text input, search field, textarea, amount input, phone input, OTP input.
- Bottom sheet, modal, confirmation dialog, action menu, toast, banner.
- Skeleton, empty state, error state, offline banner, retry action.

### Content and social
- Post composer, post header, media viewer, carousel, caption block.
- Reaction bar, comment composer, comment thread, share sheet, bookmark.
- Story ring, story viewer, profile header, follow button.
- Feed item, creator card, topic/category tile, discovery result.

### Live
- Live stage, viewer count, host/co-host tile, participant tray.
- Live chat message, pinned comment, gift picker, battle score panel.
- Product pin, product tray, live purchase confirmation.
- Connection quality indicator, reconnecting state, host moderation controls.

### Commerce
- Product card, variant selector, quantity stepper, stock indicator.
- Seller/shop header, rating/review summary, delivery estimate.
- Cart line, order summary, checkout step, payment method.
- Order timeline, fulfilment status, refund status, seller contact.

### Finance
- Available balance, pending amount, reserved amount, earnings summary.
- Ledger activity row, payout destination, payout status, fee breakdown.
- Transaction receipt, refund record, reconciliation status (staff-only).
- Every amount shows currency and a clear status; never imply pending earnings are withdrawable.

### Messaging and community
- Conversation row, message bubble, attachment, voice note, delivery/read receipt.
- Conversation request, member list, role badge, invite flow.
- Community feed, event card, poll, moderation action.

### Creator and business
- Analytics metric, chart, content performance row, monetisation status.
- Product manager row, inventory status, order queue, campaign card.
- Settings group, permission explanation, integration status.

## 6. Interaction and state standards

Every interactive component must define:
- Default, pressed, focused, disabled and loading states.
- Validation and error messaging adjacent to the affected control.
- Haptic feedback only for meaningful actions and where platform conventions support it.
- Optimistic updates only when safely reversible; financial actions must reflect authoritative server state.
- Double-submit prevention and idempotent action handling for purchases, payouts and publishing.
- Confirmation for destructive, irreversible or financially material actions.
- Clear retry behavior that does not create duplicate messages, orders, payments or payouts.

### Required global states

- First launch and onboarding.
- Signed out, session expired and account restricted.
- No network, slow network and reconnecting.
- Permission not granted (camera, microphone, photos, notifications, location).
- Empty feed, no search results, no messages, no orders and no earnings.
- Upload processing, upload failed and upload resumed.
- Payment pending, confirmed, failed, unknown and reversal/refund.
- Live session ended, host disconnected, viewer reconnecting and content unavailable.
- Account/profile switch and access denied.

## 7. Accessibility and localisation

- Target WCAG 2.2 AA for applicable UI content and web surfaces.
- Verify contrast for text, icons, controls, focus and status.
- Screen-reader labels and logical reading order for custom controls.
- Dynamic type / font scaling without loss of functionality.
- Captions and accessible alternatives for audio/video content.
- Reduced-motion behavior for transitions, autoplay and animated gifts.
- Touch targets at least 44×44 logical points as a baseline.
- Localise dates, numbers, currencies and phone formats by market.
- Avoid embedding important text in images.
- Initial market language strategy must be configurable; English and Kiswahili should be first-class launch candidates.

## 8. Performance and low-bandwidth UX

- Feed loads text and thumbnails before high-bitrate media.
- Autoplay is preference-controlled and respects data saver.
- Video quality adapts to network conditions.
- Uploads are resumable and show clear progress.
- Avoid unnecessary re-rendering of live chat and scrolling feeds.
- Preserve scroll position and draft state when navigating to product detail or checkout.
- Do not prefetch large media on metered connections without user consent.
- Use stable skeletons; avoid layout shifts when media loads.

## 9. Product instrumentation

Design events around user outcomes, not component clicks alone:
- onboarding_completed
- profile_switched
- feed_item_viewed
- creator_followed
- content_published
- live_joined, live_left, live_reconnected
- live_product_opened
- checkout_started, payment_status_viewed, order_confirmed
- message_sent
- creator_earnings_viewed
- payout_requested

Events must not include message bodies, sensitive financial credentials, full payment identifiers or unnecessary personal data. Define event schemas and consent/retention policy before production instrumentation.

## 10. Definition of design-system readiness

The design system is ready for implementation when:
- One visual direction is selected and documented.
- Colour, type, spacing, radius, elevation and motion tokens are mapped for light/dark themes.
- Navigation and profile/mode switching are validated across core journeys.
- Core components have anatomy, variants, state coverage and accessibility notes.
- High-risk finance, checkout, live and moderation flows have reviewed interaction specs.
- Representative Android and iOS screens are reviewed at small and large device sizes.
- Storybook (or equivalent component catalogue) demonstrates the shared components and states.
