# Mobile and Web Specification

## Mobile stack
React Native + Expo, New Architecture, Android and iOS. Use development builds and native modules as required. Pin SDK versions and test native compatibility in CI.

## Feature modules
identity, profiles, home-feed, discovery, content, creator, live, messaging, calls, market, checkout, orders, earnings, communities, business, notifications, settings and trust-safety. Shared packages provide design system, API client, OIDC/session manager, secure storage, analytics, feature flags, localisation, media utilities and offline queue.

## Native capabilities
Camera/microphone, media picker, native video playback, background/resumable upload, push, secure storage, deep links, share sheet, audio session and app lifecycle.

## Low-bandwidth design
Data saver; autoplay controls; adaptive quality; thumbnail-first feed; bounded cache; offline drafts; retryable message outbox; resumable uploads; limited background refresh; network-aware prefetch; preserve checkout and message state through connectivity changes.

## Accessibility/localisation
Screen readers, dynamic text, contrast, captions, reduced motion, localisation-ready strings, time zones and currency-aware display. Currency is explicit in domain data; do not hard-code KES into calculations.

## Public website
React + Vite for marketing, company, pricing/explainer, support, developer documentation and app-download routes only. No consumer feed, inbox, checkout or merchant console. Prerender/static-generate SEO-critical pages and maintain canonical URLs.

## Internal console
Separate authenticated web application for trust/safety, support, operations and finance review. MFA, strong RBAC, reason-coded sensitive actions, time-limited access and immutable audit. No generic impersonation without break-glass workflow.
