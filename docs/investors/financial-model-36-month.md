# Minx — 36-Month Financial Model (Investor Planning Version)

**Model date:** 18 November 2026  
**Currency:** Kenyan shillings (KES)  
**Period:** 36 commercial operating months following a 6-month build and 3-month controlled beta.

> This is an assumption-led planning model, not audited financial information, a provider quotation, a valuation, an investment offer or a guarantee. The monthly forecast is a linear intra-year ramp of annual average operating assumptions. Actual monthly cash needs depend on hiring dates, customer collection timing, settlement delays, seasonality, taxes and signed vendor terms.

## 1. Reconciled base case

| KES millions | Year 1 | Year 2 | Year 3 |
|---|---:|---:|---:|
| Average MAU | 10,000 | 35,000 | 100,000 |
| Marketplace GMV | 50.400 | 252.000 | 972.000 |
| Marketplace commission | 2.520 | 12.600 | 48.600 |
| Creator monetisation share | 0.384 | 2.520 | 11.520 |
| Merchant subscriptions | 2.700 | 14.400 | 54.000 |
| Advertising | 0.600 | 4.200 | 21.600 |
| **Total platform revenue** | **6.204** | **33.720** | **135.720** |
| Payment processing | 0.907 | 4.536 | 17.496 |
| Media + storage | 1.330 | 6.170 | 22.600 |
| **Total operating costs** | **18.437** | **38.906** | **89.896** |
| **Operating surplus/(deficit)** | **(12.233)** | **(5.186)** | **45.824** |

### Revenue formulas
- Marketplace GMV = average MAU × buyer conversion × monthly GMV per buyer × 12.
- Marketplace commission = GMV × 5%.
- Creator monetisation = MAU × creator payer conversion × monthly creator spend × 12 × Minx share (20%).
- Merchant subscriptions = average paying merchants × monthly fee × 12.
- Advertising = average MAU × advertising ARPU per month × 12.

### Base-case operating assumptions

| Assumption | Year 1 | Year 2 | Year 3 |
|---|---:|---:|---:|
| Average MAU | 10,000 | 35,000 | 100,000 |
| Buyer conversion | 12% | 15% | 18% |
| Monthly GMV per buyer | KES 3,500 | KES 4,000 | KES 4,500 |
| Marketplace take rate | 5% | 5% | 5% |
| Creator payer conversion | 4% | 6% | 8% |
| Monthly creator spend per payer | KES 400 | KES 500 | KES 600 |
| Minx creator share | 20% | 20% | 20% |
| Average paying merchants | 150 | 600 | 1,800 |
| Monthly merchant subscription | KES 1,500 | KES 2,000 | KES 2,500 |
| Ad ARPU / MAU / month | KES 5 | KES 10 | KES 18 |
| Viewing hours / MAU / month | 1.0 | 1.5 | 2.0 |
| Payment processing assumption | 1.8% of marketplace GMV | 1.8% | 1.8% |
| Media delivery assumption | KES 9 / viewer-hour | KES 9 | KES 9 |

Annual operating-cost assumptions (KES): staff 9M / 15M / 24M; marketing 3M / 6M / 12M; core infrastructure 1.8M / 3M / 6M; support and moderation 1.2M / 2.4M / 4.8M; G&A/legal/compliance 1.2M / 1.8M / 3M. Storage allowances are KES 250k / 500k / 1M per year. Payment and media cost assumptions must be replaced with contracted rates and actual usage.

## 2. Monthly base-case cash flow

The detailed 36-month CSV is in [monthly-base-case.csv](monthly-base-case.csv). MAU and paying merchants ramp linearly from 50% to 150% of each year's average across that year, preserving the annual average. Revenue and operating costs are modelled monthly; the cumulative operating cash flow identifies the maximum operating deficit before build/launch costs and reserve.

- Peak cumulative operating deficit in the model: **KES 18.541M**.
- 20% working-capital reserve on that deficit: **KES 3.708M**.
- This is a cash-flow proxy, not a substitute for a dated 36-month collection, settlement and supplier-payment schedule.

## 3. Build, beta and commercial launch

### Build and controlled beta — KES 2.45M

| Item | KES |
|---|---:|
| Go contractor | 300,000 |
| React Native developer | 250,000 |
| UX/UI | 100,000 |
| QA/security | 150,000 |
| DevOps | 100,000 |
| Legal/payment integration | 100,000 |
| Founder/lead developer remuneration (paid) | 400,000 |
| **Professional fees** | **1,400,000** |
| Development infrastructure/tools (months 1–6) | 150,000 |
| Controlled beta (3 months) | 650,000 |
| Central contingency reserve | 250,000 |
| **Total** | **2,450,000** |

The beta allocation includes AWS compute and separate Core/Finance databases, Cloudflare Stream, LiveKit, monitoring/security, QA/support, creator/merchant onboarding, beta operations, payment testing/reconciliation and release administration. AWS KES 180k and other provider budgets are planning allowances, not verified quotes.

### One-off commercial launch allowance — KES 5.00M

| Item | KES |
|---|---:|
| Public launch marketing/acquisition | 2,500,000 |
| Creator and merchant activation | 1,000,000 |
| Launch support/moderation/operations setup | 750,000 |
| Legal/compliance/app release/launch contingency | 750,000 |
| **Total** | **5,000,000** |

## 4. Indicative funding bridge

| Component | KES |
|---|---:|
| Build and controlled beta | 2,450,000 |
| Commercial launch allowance | 5,000,000 |
| Peak cumulative operating deficit | 18,540,715 |
| 20% working-capital reserve | 3,708,143 |
| **Indicative planning envelope** | **29,698,858** |

**Indicative envelope: KES 29.70M.** This replaces the earlier KES 28.35M estimate, which understated the reconciled operating deficits. Do not present this as a firm raise amount until the monthly model is reviewed against a dated hiring plan, actual provider quotations, tax and payment settlement assumptions.

## 5. Scenario comparison

| Scenario | Y1 revenue (M) | Y1 operating result (M) | Y2 revenue (M) | Y2 operating result (M) | Y3 revenue (M) | Y3 operating result (M) |
|---|---:|---:|---:|---:|---:|---:|
| Conservative | 1.821 | (11.773) | 8.227 | (13.913) | 27.779 | (11.001) |
| Base | 6.204 | (12.233) | 33.720 | (5.186) | 135.720 | 45.824 |
| Growth | 12.915 | (11.293) | 74.370 | 15.308 | 312.264 | 155.840 |

Conservative and Growth cases are alternative planning sensitivities, not probability-weighted forecasts. See [scenario-assumptions.md](scenario-assumptions.md).

## 6. Sensitivities and validation work

Before investor circulation, validate at least:
1. MAU ramp, retention, buyer conversion, creator payer conversion and paying merchant retention.
2. Marketplace take rate, refunds, cancellations, fulfilment failures and payment-provider fees.
3. Creator monetisation share, provider fees, chargebacks, refunds and payout costs.
4. Advertising fill rate, advertiser demand, brand-safety controls and realistic monetisation start date.
5. Viewer-hours, bitrate mix, recording retention, CDN delivery, LiveKit usage and overage pricing.
6. Hiring start dates, fully loaded compensation, moderation staffing and customer support coverage.
7. Actual collection/settlement delays, VAT/tax, app-store fees, FX, legal and regulatory costs.
8. Downside case: launch delayed by 3–6 months, MAU 50% below plan, media unit cost +50%, and settlement delays.

## 7. Non-negotiable financial architecture assumptions

No customer stored-value wallet, C2C transfers or transferable/cash-redeemable coins at launch. Creator and merchant balances are contractual payables, paid through appropriately licensed partners. Finance maintains an isolated database, double-entry ledger, immutable posted journals, idempotent payment/callback/payout handling and reconciliation. Legal counsel must confirm PSP, marketplace, gifts, KYC/KYB, AML and tax scope before launch.
