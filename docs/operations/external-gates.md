# External Gates and Assumptions

The engineering baseline is approved. The following require evidence and cannot be marked complete by architecture documentation alone.

| Gate | Evidence required | Blocks |
|---|---|---|
| AWS regional availability | Verify ECS/Fargate, RDS PostgreSQL 18, cache, secrets, telemetry, backup in af-south-1 | Production infrastructure |
| Data transfer | Data map, controller/processor roles, ODPC assessment, contracts and safeguards | Production personal-data processing |
| PSP | Current CBK authorisation/scope, collection, marketplace settlement, refunds, payouts, SLA and signed contract | Live money movement |
| Marketplace legal model | Kenyan counsel opinion on platform role, funds flow, merchant settlement and creator payable | Marketplace checkout |
| Gifts/coins | Legal, accounting, tax and consumer treatment | Gifts/redeemable credits/cash-out |
| KYC/KYB/AML | Partner responsibilities, verification vendor, thresholds and escalation | Creator/merchant payouts |
| Media | LiveKit/Stream end-to-end benchmark, regional latency, quality, recording and cost | Public live launch |
| NATS | Region, replication, retention, backup, SLA and export path | Production event infrastructure |
| OIDC | Region, OTP delivery, MFA, recovery, pricing, export and exit path | Production authentication |
| Disaster recovery | Approved secondary region and transfer safeguards | Production DR |
| Mobile SDKs | Expo/RN, LiveKit, media, push, secure storage and camera compatibility matrix | App store release |

## Procurement rules
Verify actual processing locations and subprocessors, not just global availability. Obtain DPAs and breach notification commitments. Confirm export and exit path before critical state is stored. Document callback authenticity per provider. Never infer payment authorisation from a vendor relationship.
