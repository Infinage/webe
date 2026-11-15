---
name: VSS Associates
description: Proof-led plotted-development site for a Chennai land firm — navy, brass, and survey paper
colors:
  navy: "#0e2a47"
  navy-deep: "#081d33"
  gold: "#c9a24b"
  gold-soft: "#e8d9b4"
  brass-dark: "#9a7422"
  brass-ink: "#856517"
  paper: "#f6f4ef"
  card: "#ffffff"
  line: "#e3ded3"
  ink: "#1d232b"
  ink-soft: "#4a5563"
  ink-faint: "#8a94a3"
  green: "#1d6b45"
  green-bg: "#eef7f1"
  red: "#a3543d"
  red-bg: "#fbf1ee"
typography:
  display:
    fontFamily: "Manrope, ui-sans-serif, system-ui, sans-serif"
    fontSize: "clamp(40px, 5.4vw, 64px)"
    fontWeight: 800
    lineHeight: 1.06
    letterSpacing: "-0.02em"
  headline:
    fontFamily: "Manrope, ui-sans-serif, system-ui, sans-serif"
    fontSize: "clamp(28px, 3.4vw, 40px)"
    fontWeight: 800
    lineHeight: 1.12
    letterSpacing: "-0.015em"
  title:
    fontFamily: "Manrope, ui-sans-serif, system-ui, sans-serif"
    fontSize: "19px"
    fontWeight: 800
  body:
    fontFamily: "Manrope, ui-sans-serif, system-ui, sans-serif"
    fontSize: "16px"
    fontWeight: 400
    lineHeight: 1.6
  lead:
    fontFamily: "Manrope, ui-sans-serif, system-ui, sans-serif"
    fontSize: "18px"
    fontWeight: 400
    lineHeight: 1.65
  label:
    fontFamily: "Manrope, ui-sans-serif, system-ui, sans-serif"
    fontSize: "12px"
    fontWeight: 700
    letterSpacing: "0.16em"
  button:
    fontFamily: "Manrope, ui-sans-serif, system-ui, sans-serif"
    fontSize: "15px"
    fontWeight: 600
rounded:
  sm: "8px"
  md: "12px"
  pill: "999px"
spacing:
  section: "96px"
  section-tight: "64px"
  gutter: "28px"
components:
  button-gold:
    backgroundColor: "{colors.gold}"
    textColor: "{colors.navy-deep}"
    rounded: "{rounded.pill}"
    padding: "15px 26px"
    typography: "{typography.button}"
  button-navy:
    backgroundColor: "{colors.navy}"
    textColor: "#ffffff"
    rounded: "{rounded.pill}"
    padding: "15px 26px"
    typography: "{typography.button}"
  button-outline:
    backgroundColor: "transparent"
    textColor: "{colors.navy}"
    rounded: "{rounded.pill}"
    padding: "15px 26px"
    typography: "{typography.button}"
  badge-dtcp:
    backgroundColor: "{colors.green-bg}"
    textColor: "{colors.green}"
    rounded: "{rounded.pill}"
    padding: "6px 12px"
  card:
    backgroundColor: "{colors.card}"
    rounded: "{rounded.md}"
  input:
    backgroundColor: "{colors.card}"
    textColor: "{colors.ink}"
    rounded: "{rounded.sm}"
    padding: "13px 15px"
---

# Design System: VSS Associates

## Overview

**Creative North Star: "The Verified Ledger"**

Meridian treats the site as an official registry of the ground. The aesthetic currency is proof: DTCP approval stamps, spec ledgers with label-left/value-right rows, plot ledgers with status dots, and honest `—` placeholders where the client has not yet supplied a fact. Navy owns hierarchy; brass (gold) is spent like wax on a seal — only on actions and proof marks, never decoration. White instrument cards float on warm survey paper.

The system reads institutional, not corporate-glossy. Photography is real site ground — soil, lawns, roads, signboards — because buyers buy land. Density alternates deliberately: full-bleed image moments against dense typographic ledgers, one message per hero.

**Key Characteristics:**
- Navy hierarchy, brass scarcity, warm paper canvas
- Single type family (Manrope) with tabular numerals on every figure
- Ledger grammar: uppercase micro-labels opposite tabular values, separated by hairlines
- Approval/verification as designed first-class elements (stamps, green badges), never footnotes
- Unconfirmed facts render as `—` / marked `*` sample data — visible honesty as craft

## Colors

A restrained maritime-institutional palette: one deep navy family for authority, one brass family for actions and proof, warm paper neutrals, and two status hues that exist only for verification states.

### Primary
- **Deep Harbor Navy** (`#0e2a47`): Headers, primary buttons, headline text, depth. The system's voice of authority; owns all structural color.
- **Harbor Midnight** (`#081d33`): Footer canvas, navy-button text on gold, deepest overlay layers in hero veils.

### Secondary
- **Antique Brass** (`#c9a24b`): CTAs, accents, proof marks. Appears on ≤10% of any screen — its rarity is the meaning.
- **Aged Brass** (`#9a7422`): Eyebrow labels, links-on-hover, unit suffixes, the `*` requirement mark on light backgrounds.
- **Parchment Brass** (`#e8d9b4`): Eyebrow/label color on navy surfaces; h1 accent words over photography.

### Neutral
- **Survey Paper** (`#f6f4ef`): Page background. Warm, not clinical white.
- **Card White** (`#ffffff`): Cards, panels, inputs, available plots.
- **Hairline Sand** (`#e3ded3`): All card borders, dividers, spec-row rules.
- **Ledger Ink** (`#1d232b`): Body text.
- **Ink Soft** (`#4a5563`): Secondary text, leads, spec labels.
- **Ink Faint** (`#8a94a3`): Hints and tertiary text only.

### Status (reserved)
- **Verified Green** (`#1d6b45`, bg `#eef7f1`): DTCP-approved badges and available states. Never decorative. Extended to the rotated DTCP stamp ink (critique decision 2026-11-13): the seal shares the badge hue — never terracotta, which means *sold*.
- **Sold Terracotta** (`#a3543d`, bg `#fbf1ee`): Sold/inactive states only.

### Named Rules
**The Gold-Sparseness Rule.** Brass appears only on actions (buttons, CTAs) and proof marks (top rule of the stat band, eyebrow labels). If gold makes something prettier but not more actionable or more provable, remove it.

**The Green-Means-Verified Rule.** `#1d6b45` is exclusively a verification signal — DTCP approval, plot availability. No green headings, no green decoration.

**The Honest-Placeholder Rule.** Unconfirmed facts render as `—` or marked `*` sample data with the `.sample-flag` ribbon. Never lorem, never invented numbers, never silent blanks.

## Typography

**Display Font:** Manrope (400/500/600/700/800) with system sans fallback — the *only* family in the system.
**Body Font:** Manrope.

**Character:** Precise and institutional. One family, weight does all the hierarchy work: 800 for display/headlines, 700 for labels, 400 for reading. Numerals are tabular (`font-variant-numeric: tabular-nums`) on every stat, price, and ledger value so columns align like a registry.

### Hierarchy
- **Display** (800, `clamp(40px, 5.4vw, 64px)`, 1.06, −0.02em): Hero h1 only. One message per hero; accent words in Parchment Brass over photography.
- **Headline** (800, `clamp(28px, 3.4vw, 40px)`, 1.12, −0.015em, navy): Section h2.
- **Title** (800, 19px, navy): Card titles.
- **Lead** (400, 18px, 1.65, ink-soft, max 62ch): Section intros.
- **Body** (400, 16px, 1.6): All reading text.
- **Label** (700, 12px, 0.16em, uppercase): Eyebrows (Aged Brass on light, Parchment Brass on navy); 11.5px/0.14em variant inside stat bands; 11px/0.12em for spec-row labels. Small brass text (eyebrows, `.req`, reserved badges, stat units, FAQ indices, link hovers) uses **Brass Ink `#856517`** — the ≥4.5:1 compliant cut of the brass family; the brighter brasses stay on fills, seals and hover accents.
- **Button** (600, 15px): Pill buttons.

### Named Rules
**The Tabular Ledger Rule.** Every figure on the page uses tabular numerals. Stats, prices, plot numbers, availability counts — all align vertically.

## Layout

Content lives in a 1200px container (`--maxw`) with 28px side padding. Sections breathe at 96px vertical (64px tight variant, 64px below 900px). The header is fixed: transparent over heroes with a navy gradient scrim, turning solid Survey Paper with a hairline shadow once scrolled. Heroes are full-bleed, min-height 92vh, content bottom-aligned, with the validated navy veil gradient (`.60/.24/.78`).

Signature spatial pattern: the horizontal ledger stack — cards in a 302px-column scroll-snap row with floating circular arrows (hidden below 640px, columns go 78vw). Project pages alternate asymmetric split rows (photo one side, typographic ledger the other, mirrored per row). On the homepage the showcase is deliberately asymmetric (critique redesign): a large lead layout card with a 2×2 spec ledger, two compact photo-row cards beside it, and the unlaunched project as a dashed-outline footnote strip — never a uniform card row. Footer is a 4-column navy-deep grid (1.6fr 1fr 1fr 1.2fr).

Breakpoints: 900px (nav links hide; a gold "Book a Visit" pill + circular tap-to-call icon take over the header bar, sections tighten, footer 2-col) and 640px (carousel arrows hidden, card columns 78vw, geo pins become labeled cards — pin glyph above, tip text as a static caption). The desktop header carries no gold CTA: the fixed bar would duplicate every in-page "Book" moment as the page scrolls (critique decision) — booking lives in the hero, the closing CTA band, the contact page, and the mobile header pill.

## Elevation & Depth

Layered, two-step. Surfaces are flat with hairlines at rest; the ambient shadow makes white cards float on paper without drama.

### Shadow Vocabulary
- **Ambient** (`0 2px 8px rgba(14,42,71,.09)`): Cards, stat band, resting surfaces.
- **Lift** (`0 14px 40px rgba(14,42,71,.16)`): Floating chrome only — carousel arrows, sticky elements, the sample-data ribbon, and hover lift (−4px translate).

### Named Rules
**The Two-Step Shadow Rule.** Never a third shadow strength. Depth is response or floating chrome, never decoration — shadows don't pile up inside nested cards.

## Shapes

Gentle institutional rounding: 8px for small controls and inputs, 12px for cards, 999px pills for buttons and badges. Hairlines (1px `#e3ded3`) carry most structure. Proof bands carry a 3px brass top rule — the stat band's signature. Card photos inside cards use small 10px radii with a navy bottom-gradient scrim for legibility. The DTCP stamp is a rotated (−9°) sticker treatment on imagery, like a physical seal pressed on a survey drawing. Project rows in the index carry a full 1px hairline border with 14px radius and may carry a static ±1° photo tilt for editorial character.

## Components

Buttons, cards, and inputs feel like official instruments — precise, tabular, quiet until interacted with.

### Buttons
- **Shape:** Full pill (999px), 15px/26px padding, 600 weight.
- **Primary (gold):** Antique Brass fill, Harbor Midnight text, soft brass glow shadow. Hover: lightens to `#d6b15e`, lifts 1px. (Hover keeps navy text — white-on-gold is 2.4:1; the buying-process circle hover on 01 uses the same navy-on-gold fill.)
- **Label (canonical):** every booking CTA reads **“Book a Visit”** — header pills, CTA bands, and the contact form button alike (user decision, supersedes BRIEF.md's “Book a Site Visit” label; in-body prose may still say “site visit”).
- **Navy:** Deep Harbor Navy fill, white text. For actions on light surfaces where gold is already spent.
- **Ghost (dark):** 8% white fill, 40% white border, 6px backdrop blur — over photography only.
- **Outline:** Transparent, 1.5px navy border; hover inverts to navy fill/white text.
- **Text:** No fill, navy, weight 700; arrow glyph slides 4px right on hover, text warms to Aged Brass.

### Badges / Chips
- **Style:** Pill, 11.5px 700 uppercase with 0.08em tracking, 6px/12px padding.
- **DTCP badge:** Verified Green on green bg with tick. Status badges: available = green family; reserved = brass family (`#fdf6e7`/Brass Ink); sold = terracotta family; neutral = paper bg + hairline border.

### Cards / Containers
- **Corner Style:** 12px radius, overflow hidden.
- **Background:** Card White on Survey Paper.
- **Shadow Strategy:** Ambient at rest; Lift + 4px rise on hover; inner photo scales 1.045.
- **Border:** Hairline Sand where cards sit edge-to-edge.
- **Internal Padding:** 16–22px; ledger footer rows separated by hairline top borders.

### Spec Ledger Rows
- **Style:** Label (11px uppercase 0.12em, ink-soft) left, tabular value (13px 700 navy) right, hairline top border per row. The system's data voice — used on cards, admin tables, and plot dossiers. On showcase cards the status row below the ledger carries **no divider — whitespace only**; the ledger's own bottom rule closes the grid.

### Inputs / Fields
- **Style:** Card White fill, hairline border, 8px radius, 13px/15px padding, 500 weight.
- **Focus:** Brass border + 3px brass glow ring (`rgba(201,162,75,.18)`).
- **Labels:** 12.5px 700 uppercase 0.06em above the field; hints in Ink Faint.

### Plot Map (signature component)
SVG site plan on a cream road spine. Plot states: available = white/navy hairline, reserved = `#f3e7c9`/brass, sold = `#14314f` filled with 75% white label. Hover: 2.4px brass stroke + soft glow; selected: 3px Harbor Midnight stroke + deep shadow. Selection opens a dossier panel (spec-ledger rows: number, area, dimensions, facing, price, status). Zoom/pan/fit controls float bottom-right and fade in on interaction.

### Step & Index Numerals
- **Buying-process circles (`.step .num`):** 54px gold-ring circles; hover anywhere on the step segment (not just the circle) fills the circle Antique Brass with white text.
- **Value-list numerals (`.prop .n`):** 64px outlined tabular glyphs on paper (transparent fill, `#cbb47e` stroke). Hollow at rest — no scroll-triggered fill. They fill Antique Brass (stroke and interior) while the row is hovered, and briefly (~1.2s flourish) when the row is clicked. This is the one approved gold fill for numerals; it applies to flat page copy only — see the Do/Don't note about ghost numerals on imagery.

### Sample-Data Ribbon
Fixed bottom-right navy pill (`rgba(14,42,71,.92)`) listing what awaits client supply — the Honest-Placeholder Rule made visible.

## Do's and Don'ts

### Do:
- **Do** use tabular numerals on every stat, price, and count.
- **Do** reserve green for DTCP verification and availability only.
- **Do** mark every unconfirmed figure with the `*` mark and `.req` class, with the sample-flag ribbon on the page.
- **Do** show the ground — soil, lawns, roads, signboards — in photography.
- **Do** keep the validated hero veil (`.60/.24/.78` navy gradient) on photographic heroes.
- **Do** use the 3px brass top rule to mark proof bands.

### Don't:
- **Don't** use brass/gold as decoration — actions and proof only (client-locked).
- **Don't** print signboard approval numbers (e.g. "No. 26/2016") as page text — they are visible in photos but unconfirmed.
- **Don't** use real VSS site photos as full-bleed heroes (low-res camera shots); the only hero-grade image is the aerial concept render, never captioned as a specific project.
- **Don't** invent facts, testimonials, statistics, prices, or approvals — render `—` instead.
- **Don't** gradient-fill ghost numerals on project imagery — client explicitly rejected gold and glossy-white fills; keep transparent glyphs with the 1.5px white stroke.
- **Don't** copy reference sites (Anehana, Hosachiguru, Kendrick Guehr, G Square) — borrow ideas, never look.
