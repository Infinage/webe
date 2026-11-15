# VSS Associates — Design System & Information Architecture

Status: **plan only — no implementation yet.**
Source: `site_content/` (extracted legacy content), `.lavish/design-directions.html` (approved Meridian direction, client-locked).

The client decision on record: **Direction A "Meridian"** — navy + gold, established-professional tone. Feedback incorporated: hero shows the ground (lawn → shore), portrait/vertical project cards, no dark canvas, no filter-bar UX.

---

## 1. Design tokens

### Color
| Token | Value | Use |
|---|---|---|
| `--navy` | `#0e2a47` | Headers, depth, primary text on dark surfaces |
| `--gold` | `#c9a24b` | CTAs, accents, proof marks — **sparing: actions + proof only** |
| `--paper` | `#f6f4ef` | Page background (warm paper) |
| `--card` | `#ffffff` | Cards, panels |
| `--line` | `#e3ded3` | Hairlines, card borders |
| `--ink` | `#1d232b` | Body text |
| `--ink-soft` | `#4a5563` | Secondary text |
| `--green` | `#1d6b45` (bg `#eef7f1`) | DTCP/approval "verified" states |
| `--red` | `#a3543d` (bg `#fbf1ee`) | "sold out / inactive" states |

Rules: gold never used for decoration; navy owns hierarchy; white cards float on warm paper; green reserved exclusively for verification signals.

### Typography
- **Single family: Manrope** (400/500/600/700/800). Precise, institutional.
- Scale (desktop reference): hero h1 ~64px/800 · section h2 ~40px/800 · card title 18px/700 · body 16px/400 · eyebrow/label 12px uppercase, letter-spacing 0.08–0.18em.
- Numerals: tabular-nums for all stats and prices.

### Shape, spacing, elevation
- Radii: 8px (small controls) / 12px (cards) / 999px (pills, badges).
- Section rhythm: 96px desktop / 56px mobile vertical padding; max content width ~1200px.
- Elevation: single soft shadow `0 2px 8px rgba(14,42,71,.09)` for cards; stronger shadow only for sticky/floating elements.
- 3px gold rule as top border on proof bands (the stat band signature).

### Imagery rules (client-mandated)
1. **Show the ground** — buyers buy land; favour photos with visible soil/lawn/roads over skies and canopies.
2. Real site photography only — photos with signboards are the strongest trust asset.
3. Hero: full-bleed, dark gradient overlay top+bottom for legibility (validated overlay: `.60/.24/.78`).
4. Placeholder discipline: anything unconfirmed renders as `—` with a `* = client to confirm` footnote; **no lorem, no fake numbers**.

---

## 2. Reusable components

Ordered by build sequence. All are content-driven (schema first, then markup).

| # | Component | Anatomy / notes | Used on |
|---|---|---|---|
| C1 | **Header / Nav** | Logo `VSS.` + links (Projects, About, Approvals, Contact) + gold pill CTA "Book a Site Visit". Transparent-over-hero variant (gradient scrim) for pages with a hero; solid paper variant for the rest. Tap-to-call phone in mobile variant. | all |
| C2 | **Hero** | Full-bleed image, h1 + one-line sub + single gold CTA. One message per page. | Home, Project detail, About |
| C3 | **Stat band** | 5 cells on white with gold top rule: `DTCP ✓ / — yrs / — + / — acres / 100%`. Numbers render as `—` until confirmed. | Home (also reusable as narrow proof strip on About) |
| C4 | **Project card (portrait)** | Tall photo → name → `— /sq.ft · — sq.ft plots` → green `✓ DTCP APPROVED`. Click → project page. Grid of 3 (desktop) / 1 (mobile). | Home, Projects index |
| C5 | **DTCP / status badge** | Pill, green family; only appears with a real approval number on detail pages. | cards, project pages, Approvals |
| C6 | **Section header** | Eyebrow (uppercase gold-brown) + h2 + optional sub. | everywhere |
| C7 | **Photo gallery** | Real-photos-first grid, lightbox; captions name the viewpoint. | Project detail, About |
| C8 | **Process timeline** | Visit → choose plot → agreement → registration → patta. Numbered steps, gold numerals. | Home (short), About or own section |
| C9 | **Testimonial** | Name + plot number + project (needs client supply; renders only with real content). | Home, Project detail |
| C10 | **Enquiry / site-visit form** | Name, phone, project (dropdown), optional message. Short on purpose. | Contact, modal/drawer from CTA |
| C11 | **Sticky site-visit CTA** | Persistent bar/button with phone `044-42087913` on scroll. | all, mobile-first |
| C12 | **Footer** | Address, landline, email, hours, project links, DTCP line, copyright. | all |
| C13 | **Verification strip** | Checklist row (clear title · patta support · DTCP number · bank loans). | Project detail |
| C14 | **FAQ accordion** | Optional, only with real client answers. | Contact/About |

No component without confirmed or clearly-marked-placeholder content gets built.

---

## 3. Pages & information architecture

```
Home
├─ Projects (index of all layouts)
│   ├─ Landmark City
│   ├─ Srivelan Nagar
│   ├─ Chandrayan Nagar
│   └─ OMR project (currently a 404 — rebuild; confirm it is active)
├─ About (story, approvals, team)
├─ Approvals (its own nav item — DTCP numbers, patta policy)
└─ Contact (map, phone, hours, site-visit booking)
+ 404
```

**Conversion path:** Home → Projects → Project detail → Book site visit; About & Contact one tap away everywhere; sticky CTA + tap-to-call on all pages.

**Page anatomy (shared skeleton):**
Hero → proof → project showcase → process → testimonials → one clear CTA.

### Per-page content map
- **Home:** hero (headline + sub + CTA) · stat band (C3) · 3 portrait project cards (C4) · short process (C8) · testimonials (C9, if available) · closing CTA.
- **Projects index:** grid of all project cards + one-line pitch per project + status badges.
- **Project detail:** hero (project photo) · facts row (location, extent, plot sizes, rate/sq.ft, approval no.) · verification strip (C13) · gallery (C7) · location/access · enquiry form (C10).
- **About:** story & years · team · how we maintain layouts · process timeline (C8).
- **Approvals:** what DTCP approval means · per-project approval numbers · patta/clear-title policy · how a buyer verifies.
- **Contact:** map · tap-to-call `044-42087913` · `contact@vssassociates.com` · hours · enquiry form (C10).
- **Pricing page: dropped** (fake template content in legacy site).

### Mobile-first note
Legacy site is not responsive at all — every layout above gets a single-column mobile pass: portrait cards stack, stat band becomes 2×2+1, sticky CTA (C11) replaces header CTA.

---

## 4. Information to capture — and the order to capture it

Principle: **capture global facts first** (they unblock every page), then home-blocking content, then per-project data, then nice-to-haves. Every item below currently renders as `—` in the mockup.

**Tier 1 — Global facts (block everything; needed before any page is real)**
1. Years in business (stat band, About, trust copy).
2. Families served / customers count.
3. Total acres developed.
4. Official address, working hours, any second phone/WhatsApp.
5. Confirm the 4 projects list + which are active (is OMR alive?).

**Tier 2 — Per-project data (blocks Home cards + all project pages)**
6. Per project: full name, location/area, total extent (acres), number of plots, plot sizes (sq.ft), price/sq.ft (or "on request"), **DTCP approval number**, current availability status, 5–10 real photos each (signboards preferred), one-line pitch.

**Tier 3 — Home-blocking extras**
7. Hero choice (reference coastal photo vs Unsplash alternatives — form in the artifact; current default: client reference photo).
8. 1–3 testimonials with real names + plot numbers (or drop the section).
9. Process steps confirmation (visit → agreement → registration → patta — exact wording of each step).

**Tier 4 — Supporting pages**
10. About story (founding year, family-firm narrative, maintenance practice) + team names/roles/photos.
11. Approvals page content: the DTCP numbers per project (subset of #6) + any patta/title documents they can legally publish.
12. FAQ items (optional).

**Capture order rationale:** Tier 1 fills the stat band and every footer/header; Tier 2 is the largest volume and gates the core sales pages; Tier 3 is small but shapes the homepage's emotional layer; Tier 4 can trail without blocking launch.

**Rule for gaps:** anything missing at build time renders as an explicit `—` placeholder with the footnote "client to confirm" — never invented, never lorem.

---

## 5. Deliberately out of scope
- Dark editorial theme (Heritage direction) — dropped by client decision.
- Filter-bar/search UX (Plotfinder concept) — too little inventory today; revisit if project count grows.
- Pricing page — legacy content was template filler.
- Blog/news — no sustained content source; reconsider only if client commits to writing.
