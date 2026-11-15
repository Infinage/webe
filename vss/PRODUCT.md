# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

## Users

**Primary:** Plot buyers in and around Chennai — first-time home/plot builders comparing DTCP-approved layouts. They are evaluating legitimacy (is this layout really approved?), location suitability, and plot availability; their goal is to shortlist a project and book a site visit.

**Secondary (confirmed by designed admin screens 07–09):** VSS office staff, who log in to manage projects, plot layouts, plot statuses (available/reserved/sold), and enquiries.

## Product Purpose

Revamp of the VSS Associates website (vssassociates.com — currently a lorem-heavy Bootstrap 3 template). The new site presents VSS's DTCP-approved plotted developments as proof-led, honest marketing: real site photography, verified-approval signals, transparent plot data. Success = contact-form enquiries, phone calls, and site-visit bookings.

## Positioning

Trust through verification: a Chennai plotted-development firm whose site proves rather than claims — DTCP approval states, real signboard photography, and visible plot data — where neighbouring developer sites typically use stock imagery and unverifiable marketing copy. The tone on record: "verified already", "is just… honest".

## Operating Context

- Enquiries convert through phone (044-42087913), email (contact@vssassociates.com), and in-person site visits; office hours Mon–Sat 8.00–18.00, Sunday closed.
- VSS buys land, obtains DTCP approval, lays out plots, and sells them; buyers then build. The buying journey is: discover → verify approval → study plot layout → enquire → site visit → agreement → registration → patta.
- Design artifacts live in `.lavish/screens/` (9 standalone HTML mockups, client-reviewed iteration by iteration); the build stack is Go + html/template + Tailwind + Datastar, server-rendered (mandated in project AGENTS.md).

## Capabilities and Constraints

- **Placeholder discipline (hard rule):** no lorem, no invented facts, fake numbers, testimonials, or stock family photography. Anything unconfirmed renders as `—` or clearly-marked sample data with the `*` flag and `.sample-flag` ribbon.
- **Pricing policy (user decision):** price-on-request as the official policy; rough indicative estimates are shown where available (marked `*` until client-supplied).
- **Interactive plot layout:** plot-map with zoom/pan/fit, plot selection with detail panel, status legend as filter (available/reserved/sold). On real layouts this becomes the central browsing tool.
- **Privacy/consent:** the contact form requires an explicit consent checkbox (starts unchecked). The current mockups run privacy gates that work with JavaScript disabled — progressive enhancement is an operating expectation.
- **Approval numbers on signboards** (e.g. "DTCP APPROVED No. 26/2016") are visible in photos but NOT confirmed data — never printed as page text.
- **Undecided (open):** real surveyed plot layouts and per-plot data for Landmark City; working-hours/contact confirmations; founder quote and milestone years for About; price display per project before launch.

## Brand Commitments

- **Meridian design system is client-locked:** navy `#0e2a47` + gold `#c9a24b` (gold only for actions + proof, never decoration), warm paper background, Manrope type, green reserved exclusively for DTCP-verification states. Established-professional tone, "Direction A Meridian" chosen over a dark-canvas alternative.
- **Logo:** `assets/logo.png` (on paper) / `logo-white.png` (on navy/photo); the original company logo re-spaced version was explicitly approved after a redesign was rejected.
- **Imagery rules (client-mandated):** show the ground (soil/lawn/roads over skies); real site photography only; signboard photos are the strongest trust asset; the aerial render (`assets/aerial-hero.jpg`) is a concept render usable as brand mood only — never captioned as a specific VSS project.
- Reference sites (Anehana, Hosachiguru, Kendrick Guehr, G Square) are for borrowing ideas, never look.

## Evidence on Hand

- Real company identity and contact details (from legacy site, `site_content/README.md`).
- 24 genuine project photos with VSS signboards (`site_content/images/`, prepped variants in `assets/`) for 3 named projects: Landmark City, Srivelan Nagar, Chandrayan Nagar (plus "OMR" listed without gallery).
- 9 client-reviewed screen mockups in `.lavish/screens/` + `design-system.md` component/token record.
- **Absences future work must not fabricate:** testimonials, client approvals, statistics/years-in-business numbers, team/founder info, per-plot real data, real surveyed layouts, awards. The legacy About page was ~100% placeholder — everything from it was discarded.

## Product Principles

1. **Proof before claim.** Every trust statement must trace to a real asset or render as marked sample data; unverifiable claims are worse than omissions.
2. **The ground is the product.** Buyers buy land — imagery and page hierarchy lead with the actual ground, not atmosphere.
3. **Honesty as a design feature.** `—` and `*` placeholders are visible craft, not gaps to paper over; the interface itself signals "nothing hidden".
4. **Small studio, big credibility.** The site must let a modest Chennai firm look institutional without pretending to be one.
