# Security, Privacy and Trust & Safety

## Security baseline
Least privilege, defence in depth, secure defaults, explicit trust boundaries, auditable privileged actions and data minimisation. Use managed OIDC, short-lived access tokens, refresh/session rotation, MFA for staff and high-risk actions, device/session revocation and resource-level server authorisation.

## Secrets and transport
Store secrets in managed secret manager. Separate credentials and signing keys per environment/service. TLS 1.2+ (prefer 1.3). Managed encryption at rest, rotation, access logs and emergency revocation. No custom crypto or homemade E2EE.

## Threat model
Account takeover, credential stuffing, SIM-swap/social engineering, IDOR, organisation escalation, fake merchants, payment replay, callback spoofing, duplicate payout, inventory race, gift fraud, bot engagement, stream disruption, harassment, malicious uploads, insider abuse and data exfiltration.

## Controls
Schema/size validation, rate limiting, WAF/bot controls, idempotency, callback inbox/replay detection, provider reconciliation, malware/content scanning, signed media URLs, audit trail, dependency/secret/SAST/DAST scanning, backup restore tests and incident response.

## Mobile security
Use OWASP MASVS/MASTG as mobile verification baseline; assess backend separately against ASVS and API threat cases. Validate secure storage, authentication, transport, platform interaction, privacy and code integrity.

## Trust and safety
Report users/profiles/posts/videos/live/messages/products/shops; block, mute, restrict, hide and appeal. Live controls include mute/remove, disable chat, suspend monetisation and terminate session. Marketplace controls include delist, freeze shop and hold payout where legally/contractually justified. Cases include evidence, severity, owner, SLA, action history and appeal. Age-appropriate policies and restricted monetisation are required.

## E2EE
Not in launch baseline. Messages are encrypted in transit and at rest with strict access. Any future E2EE must use an established audited protocol and resolve multi-device, key recovery, group changes, user-submitted evidence and abuse handling.

## Kenya privacy and payments
Map controller/processor roles, purposes, legal bases, retention, rights, processor contracts and international transfers. Complete DPIA where required and verify ODPC registration obligations. Data residency alone does not prove compliance. Obtain Kenyan counsel on PSP role/scope, marketplace settlement, gifts, stored value, tax, KYC/KYB, AML, consumer refunds and payouts.
