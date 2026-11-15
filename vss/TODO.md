# TODO — VSS website revamp

## Iteration status — visual design phase (this iteration)

**Completed**
- [x] Asset prep: `assets/` extracted from old site; `aerial-plot.png` resized/cropped to web variants (`hero-*, wide, portrait, square, card, thumb`)
- [x] Meridian design system baseline `.lavish/design-system.md` + shared `meridian.css` (client-approved; unchanged this iteration)
- [x] Brief for all 9 screens `.lavish/screens/BRIEF.md`
- [x] Screen 01 Homepage `.lavish/screens/01-home.html` — verified desktop + mobile
- [x] Screen 02 Projects index `.lavish/screens/02-projects.html` — verified desktop + mobile
- [x] Screen 03 Project detail `.lavish/screens/03-project-detail.html` — verified desktop + mobile (footer paint bug fixed during review R1)
- [x] Screen 04 Interactive plot layout `.lavish/screens/04-plot-layout.html` — verified desktop + mobile
- [x] Screen 05 About `.lavish/screens/05-about.html` — verified desktop + mobile
- [x] Screen 06 Contact / site visit `.lavish/screens/06-contact.html` — verified desktop + mobile
- [x] Screen 07 Admin login `.lavish/screens/07-admin-login.html` — verified
- [x] Screen 08 Project administration `.lavish/screens/08-admin-projects.html` — verified desktop
- [x] Screen 09 Plot/layout administration `.lavish/screens/09-admin-plots.html` — verified desktop (B6 select → details rail sync)
- [x] Per-screen design rationale `.lavish/screens/DESIGN-RATIONALE.md`
- [x] Lavish interactive review of 01-home.html — 5 feedback rounds applied (see revision log at page bottom):
      logo bug in scrolled header; logo swapped back to ORIGINAL company logo (re-spaced, no halo/box border — `assets/logo-spread.png` / `logo-white-spread.png`); kicker + h1 resizing; “apples first”
      → “approved first”; serif italic on intro; tagline verbs corrected (“verified already”, “is just… honest”); stats tooltips + FIXED SHOW, VERIFY option; schedule band copy tightened + “Sundays/marked
      holidays” + 3 items/row; 04/560 chips removed from home cards; spec rows (Area/Plots/Sizes/Price) on cards; early-flag carousel scroll-snap fix (twin left-arrows, OMR panel mis-snap, hover drift);
      sample-data chip + dev-server note (privacy gates verified incl. JS-off); tablet break-in at 1060px; value-band interactivity polish queued for a future pass
      — reviewer acknowledged final state (“Understood thanks”); no pending feedback
- [x] Lavish interactive review of 02-projects.html — in progress (rounds 1–2 applied, revision log at page bottom):
      R1: availability dot-strip (available/reserved/sold, tooltips + “N of M available*”) added to rows 01–03; ghostnum raised (bottom 96→72px) so it clears photo edge; img vertical fill fix (`height:calc(100% - 52px)`); nav active underline removed; dev-server note added
      R2: quiet row (04) given standard hairline border-bottom to match other cards; foot wrap fix (btn/avail wrap cleanly at 1280, `row-gap`); avail dots 6px/2px + 12.5px font; mobile: avail + spec `.v` wrap allowed at ≤900px, `.prow .media/.info/.specs div` min-width:0 — fixed pre-existing 537px grid blowout and clipped PLOTS/PRICE values at 390px (page horizontal overflow now 0); desktop re-verified rows 01–02
      R3: hero top whitespace under sticky header removed (`.phero` padding 172→112px desktop, 130→104px mobile); card 04 dark overlay rectangle removed — photo now framed identically to cards 01–03 (NOTE: ghostnum “04” reads faint over bright grass — reviewer to confirm if acceptable)
      R4: second whitespace pass — `.phero` padding 112→92px desktop, 104→84px mobile; card 04 media raised to shared 320px height, beige spec cells made white — card 04 now structurally identical to cards 01–03
      R5: R4's hero padding REVERTED to 112/104px on client request (keep the card-04 media unification); ghostnum gradient-fill trials (gold then glossy white) both REJECTED by client as too prominent — `.ghostnum` back to original transparent glyphs + 1.5px rgba(255,255,255,.75) stroke. Do NOT re-propose gradient fills for ghostnum
      R6 (ACCEPTED): client liked .btn-ghost-dark glass but wanted it ONLY inside the digits. mask-image:text unsupported in Chromium 157 → frost built as `.prow .media::before` (content attr(data-num), inset 26px, flex bottom-right, padding-right 18px/bottom 46px, background var(--frost) center/cover + background-clip:text + filter blur+brightness), each .media gets data-num + --frost url (same image as its img); hover scale(1.035) mirrored. Blur tuned 7→12→20px per client; 20px accepted. Session ended: "Okay good job lets start with reviewing the remaining pages"

- [x] Lavish interactive review of 03-project-detail.html — reopened for the 04→03 fold (R5+R6 applied), revision log at page bottom:
      R1: footer/rationale z-index fix — `.fixedhero` (position:fixed, z-index 0) painted OVER the static footer + rationale at page end (the long-standing "footer messed up" item); `.site-footer, .rationale { position:relative; z-index:2 }` added, verified page ends cleanly at scroll bottom. Mid-page 70vh "This ground is staked…" reveal band REMOVED per client request. `.gal` gallery (6 figures, horizontal snap scroller) given floating ‹ › buttons: `.galwrap .galnav` (46px circle, left -12px/right -12px, top 122px / 88px mobile), `galScroll(dir)` scrolls one card width + 18px gap — verified scrollLeft 0→438. Story overlap img `.m2` enlarged 46%→58% width (inline style)
      R2: galnav restyled to home-page carousel vocabulary (meridian.css .hscroll-nav look): white paper circle rgba(253,251,246,.92), navy glyph, 1px var(--line) border, hover gold border + scale 1.06 — verified
      R3: story image pair rebuilt as stacked polaroids: `.imgstack .m1` rotate(-2.2deg), `.m2` 50% width + 8px white mat rotate(2.4deg), hover on either → rotate(0) scale(1.09) + deep shadow (verified via native hover command); mobile `.imgstack` padding-bottom 40px, m2 52%/6px mat
      R4: `.gal figure` hover lift −6px + img scale 1.025 (verified); `.locrows` filled with dummy landmark names/distances (Minjur bus stand ~2 km*, Ponneri schools ~4 km*, Govt. hospital Ponneri ~6 km*, Ponneri railway station ~5 km*, Minjur market ~2 km*), all asterisked sample data
      R5: 04-plot-layout capabilities folded in at client request — zoomctl (+/−/Fit) in mapbar, wheel-zoom + pointer-drag pan on #plotmap, selectPlot auto-centres zoomed plot, panel gains Dimensions/corner/price-“—”/context CTA; "Open full plot layout →" link removed
      R6: 04-plot-layout.html DELETED (client confirmed); home card repointed to 03-project-detail.html#availability; status-filter chips added as .wchip row (superseded by R7)
      R7: filter row REMOVED per client (redundant with legend) — legend pills (.lg: All · 24/Available · 15/Reserved · 3/Sold · 6) now filter, gold active ring, dim .22; map column widened availgrid 1.7fr/1fr → 2.2fr/1fr. Large-layout plan (agreed, not built): if layouts grow much bigger, flip to full-width map + floating detail card; zoom/pan/fit covers 2–4× current plot count
      R8: zoomctl +/−/Fit moved out of mapbar (wrapped, crowded legend) → floating overlay bottom-right of plan, hidden by default, fades in on map interaction (wheel/pointerdown/pointermove) and out after 2.5s (.mapwrap .zoomctl--live); mapbar justify-content flex-start gap 12px (legend near title, not far right)
      R9: legend spacing re-balanced at client request — mapbar gap 30px + .legend margin-left:auto (right edge again, title left). ACCEPTED, 03 session closed.
      R10 (reopen session e7451578b9e03891): REAL-browser plot clicks were dead — root cause: svg.setPointerCapture on pointerdown retargeted click events away from plots. Selection moved to the svg pointerup handler (selectPlot on tap < 6px drag, deselectPlot on empty-map tap); document click-away deselect exempts .plot/#plotPanel/.zoomctl/#plotmap. Verified headless mouse clicks select A5/A11 and click-away reverts.
      R11–R13 (same session, accepted "Good done for now"): detail panel geometry reworked — #plotPanel is desktop-absolute right-column (width calc((100% - 36px)/3.2)), vertically centred (top 50% translateY(-50%)) against #availability .container in BOTH states; empty panel min-height 545px to match selected card height; duplicated fine-print under CTA removed. Mobile (≤980px) unchanged (in-flow).
    Post-session polish (user request, applied + verified):
      02-projects: .rows containerised — max-width var(--maxw) 1200px, margin auto, padding 0 28px (About-page grammar); cards no longer hug viewport edges; aerial strip + CTA band stay full-bleed. Logged on-page as R9.
      05-about: reveal-on-scroll added, same grammar as 03 (.rv/.rv-d1/.rv-d2/.rv-d3 classes + IntersectionObserver IIFE; note 05 already had a second `io` inside a scoped IIFE for counters — new observer is its own IIFE, no clash). Applied to split (imgwrap + text), vals (head + 4 vitems staggered), growth (head + .tl), proof (.pf-photos + .pf-strip), ab-cta .inner. Verified: 12/12 .rv fire on scroll; hero/mid-scroll screenshots look right.
    06-contact: consent checkbox #f-ok now starts UNCHECKED (user request — visitor must tick it).
    02-projects round 2 (user request, applied + verified): cards rounded + tilted + reveal — .rows gap 0→28px with padding-bottom 44px; .prow border-bottom → full 1px hairline border + border-radius 14px + overflow hidden; .media img border-radius 2→10px + static tilt rotate(1deg) odd rows / rotate(-1deg) even rows (hover keeps tilt + scale 1.035, frost ::before digits stay unrotated); .rv reveal on all 4 rows (even rows rv-d1 stagger) + IO script appended after header script. Verified via eval (rot matrices ±0.0174524, radius 14px, card 1 .in op 1) + screenshot. Logged on-page as R10. LESSON: a failed multi-edit call is atomic — edits[0..1] of the failed batch were silently not applied; always re-apply all edits from a failed batch.
    02-projects round 3 (user reference screenshot, applied + verified): gallery deck per project — .media now holds .deck (data-photos JSON per row) with .deck-l4 white solid, .deck-l3 tan #d8a578, img.deck-l2 second gallery photo, .deck-front white mat (10px padding, img radius 8) + .deck-next circular → button; deck keeps odd/even ±1° tilt; JS cycles front/back photos on arrow click. Row 04 gets the same deck (white card included, per user). Photo sets: 01 landmark/roads/build, 02 srivelan/site-2/visit, 03 chandrayan/site-1/stakes, 04 aerial-close/rows/band (all sample images — CLIENT INPUT REQUIRED for real per-project galleries). Verified: arrow cycles card-landmark→proof-roads; row 04 deck renders. Logged on-page as R11.
    02-projects round 4 REVERTED then superseded by round 5: client asked to swap scroll-reveal for a hover fan-out + make the deck a link (arrow removed); implemented, then client said "No please revert doesn't look very good" (m02013) — full revert to R11 (scroll reveal restored, arrow + cycling restored). Round 5 (accepted "Okay this is nice"): arrow removed entirely ("gallery is just for show; project page shows the images anyway") and the deck is now a link to 03-project-detail.html (verified by real click navigation). Verified: 4 a.deck anchors, 0 arrows, 0 cycle scripts. LESSON: hover pseudo-states can't be triggered by synthetic JS events — use agent-browser native `hover <sel>` (still didn't register :hover chain in eval; verify hover visuals via screenshot during native hover).

## Admin screens review (07 → 08 → 09) — ALL ACCEPTED — DEMO-ONLY, fine-tune later when needed
    07-admin-login.html session 64477ed676c69b72 — ACCEPTED ("Looks good, lets move on to next one", session ended by user). R1: brand logo enlarged 30→48px (height 48/width 194) per feedback on .login-brand span; screenshot-verified.
    08-admin-projects.html session c7b21f07304a3098 — ACCEPTED ("Looks good for now", session ended by user). R1: adminbar logo 18px/73px → 40px/161px (public-page size) + nav links 13→15px; same changes applied to 09-admin-plots. Window disconnected mid-review; reopened with --reopen at user request, then accepted.
    09-admin-plots.html session 9748c6391e99dcf8 — REOPENED per user request (same lavish session id), 2 rounds, ACCEPTED ("Looks great", session ended by user):
      R2 (mobile responsiveness, client: "Layout map and plot details page are not mobile responsive"): root cause — grid items' min-width:auto let the 24-plot map (min-content 916px) stretch .work to 942px at a 390px viewport. Fix: desktop grid `minmax(0,1fr) 340px` + `@media ≤960px` single column with `.work > * { min-width:0 }`; map card scrolls internally. Verified docW 390 @390px (was 942), plot details stacks cleanly, desktop unchanged. LESSON: local static server + browser may serve STALE CSS after edits — cache-bust with `?v=2` query param when verifying.
      R3 (client: "Can we have the plot layout zoomable similar to how we have it in the project details page? Add zoom in / zoom out fit to screen buttons float"): ported 03-project-detail's zoom system to 09 verbatim (adapted to 1060×470 canvas) — floating .zoomctl (+/−/Fit) bottom-right inside .mcanvas, hidden by default, fades in on wheel/pointer interaction + keyboard focus (.zoomctl--live, 2.5s fade-out); wheel zoom anchored at cursor (300–1060 viewBox clamp); pointer-drag pan; tap-select moved from click handler to svg pointerup (<6px threshold) because setPointerCapture retargets clicks (same 03 R10 lesson); select() now auto-centres a zoomed plot (edge-clamped); mapfoot hint "Scroll to zoom · drag to pan · tap a plot to edit"; removed min-width:880px + overflow-x:auto (fit view now always fits the card; zoom for labels); touch-action:pan-y (no mobile scroll trap). Verified via synthetic events: zin×2 → 611×271, zfit → 0 0 1060 470, zout clamps at fit, pan changes vb.x, tap B10 selects + fills panel, A9 recentre clamps at canvas edge (vb.x 212), keyboard ArrowLeft B7→B6, docW 390 @390px + zoomctl visible. All 9 screens reviewed + accepted.
    All 9 screens now reviewed + accepted in their Lavish sessions.

## 05-about.html review (session 595afaeb06039c27, opened next per user)
      R1: four feedback items — (1) .pf-photos figure hover lift −6px + img scale 1.025 (same vocabulary as 03 R4); (2) stat strip numerals use 01-home's .count[data-target] pattern, runs ONCE on page load (6/600/120/100 verified); (3) CTA band rebuilt with shared meridian cta-band (aerial-rows.jpg + navy veil + container inner, gold + ghost-dark buttons) replacing broken .ab-cta/.cta-inner (missing bg, bad margins); (4) hero got a photo treatment: aerial-close.jpg under a warm paper veil (rgba(247,243,236) gradient) keeping navy headline legible. All verified via screenshots. 05 has NO on-page revision log section (unlike 01–03) — log rounds here only.
      R2: veil lightened (hero photo more visible: .68/.42/.88 gradient) — user wanted background image more opaque.
      R3: growth/proof section paddings trimmed to 64px (whitespace).
      R4: footer rebuilt to canonical 4-column meridian structure (brandline white logo, Projects/Company/Visit us, two-part legal row) — now identical to 01–03.
      R5: eyebrow legibility over photo — final: near-black navy #050f1c, weight 800, 13px, light text-halo; CTA band top padding 110→48px (whitespace above Final step).
      R6: header Book a Site Visit switched btn-navy → btn-gold on 05 (matches 01–03/06 gold headers); eyebrow "more bold, darker" per user.
      R7: proof band upgraded: 5-photo mosaic (3-col grid, lead spans 2 rows) with floating white circular arrow buttons; 2 extra pool images added (proof-launch.jpg, proof-event.jpg); arrows functional — all 5 tiles rotate through 6-photo pool with staggered slide animation (55ms/tile, exit-left/enter-right on ›, mirrored on ‹, 700ms busy-lock). User: "Lets move the left most image as well. So it will basically rotate the images around" — lead photo rotates too. ACCEPTED ("Looks good, lets review the next one now"), session closed.
      Next screens: 06-contact, 07-admin-login, 08-admin-projects, 09-admin-plots

## 06-contact.html review (session 23774260f3179580 — RESUMED today)
      R1 (4 prompts, all applied + verified, then user ended session): (1) footer rebuilt to canonical 4-column meridian structure copied from 05 (brandline white logo + description, Projects/Company/Visit us columns, two-part legal row) — replaced old 3-col footer with foot-logo/f-line/foot-note markup; (2) removed Srivelan layout-entrance map card (.sidecard--maps img was explicitly flagged; whole card removed, CSS block deleted) — right column now exactly two halves; (3) #book grid restructured: .ct-body align-items stretch, form card fills full column height; .ct-side grid-template-rows 1fr 1fr with .sidecard height:100% — left = form full height, right top half = Reach us directly (navy), right bottom half = How a visit works (white). Verified via screenshots (756px column height, halves ~367px each).
      User: "Make these changes and park this for now. Will continue tomorow" — session ended; do NOT reopen unless user asks. More 06 review expected tomorrow; then 07-admin-login, 08-admin-projects, 09-admin-plots.

NEXT: 04→03 fold confirmed & applied; 04-plot-layout.html DELETED; re-review 03 session reopened — continue polling. Then 07/08/09

## 02 projects review (session b136c68d1766e343 — PARKED/ACCEPTED by user: "Looks good enough, lets park this for now")
      Opened next: 03-project-detail.html (remaining open items: footer render verification, re-review after 04→03 fold).
      Hero unification round: 05/06/02 heroes restyled to approved 05 pattern (padding 180/110, navy 68px h1, eyebrow halo). Then user feedback on 02: hero photo must cover the entire text — 02 veil now stays light (≤.5 alpha) through 82% of hero height, washing to paper only in the card-overlap seam; first project card overlap −150→−100px so it never rises above the hero text. Then user: hero height must match across all 3 pages, About's size wins — .ct-hero/.phero got min-height 678px @1280 + flex column justify-flex-end (mobile min-height:0); measured 05=678, 06=678, 02=678. Then user: text START must also match About exactly — removed flex-end (text starts at 180px top padding, not bottom-aligned) and narrowed hero text column to About's 1100px (02 .phero .grid, 06 .ct-hero .wrap). Verified via screenshots on 02+06 (eyebrow/h1 start at identical x/y as 05). R9 (02 redesign per client brief): hero split removed — single-column About/Contact grammar (eyebrow → headline → one short sentence), no gold rule/no DTCP block; sample-data note only in floating pill; rows kept (already editorial full-bleed bands). Shared invariants identified: same aerial-close.jpg hero + paper veil at 678px, 1100px column starting at 180px padding, transparent→solid header with gold CTA, paper/navy/gold tokens, hairline-bordered white cards overlapping the hero seam.

## 04 → 03 fold (client request, applied)
- Client: "I dont see a need for this page - lets move all of these complexities into prev site 03 itself. The zoom in / out and fit. Also add the extra card details to show up on click there. This page may be redundant"
- Applied to 03-project-detail.html (R5): zoomctl (+/−/Fit) in mapbar; wheel-zoom + pointer-drag pan on #plotmap; selectPlot auto-centres zoomed plot; panel gains Dimensions (dims() fn from 04), corner-plot flag (A1/A12/B1/B12), price "—" + context CTA + note for sold/reserved; removed "Open full plot layout →" link (04 reference)
- OPEN: retire 04-plot-layout.html? CONFIRMED — deleted; 01-home.html:328 card repointed to 03-project-detail.html#availability; status-filter chips carried into 03 as light-theme .wchip row (verified: Reserved dims 21 plots, All restores, screenshot clean/not cramped)
- Verified: zoom viewBox 1060→678, A11 panel (35×50 ft*, ₹33.3 lakh*), A9 sold state (price —, Ask about remaining plots), Fit restores 0 0 1060 430; screenshot checked

**Blockers — client input required (do not publish without)**
- [ ] Verified statistics (years, plots sold, layouts delivered, families) with dates/basis
- [ ] Landmark distances for Landmark City (currently "— km\*" placeholders)
- [ ] Approval number + scanned approval letter (DTCP\* format awaited)
- [ ] Real project names & facts for Project 4/5/6
- [ ] Actual surveyed plot layout (image and/or plot polygons) for Landmark City — screens 04/09 use sample geometry
- [ ] Per-plot real data: numbers, areas, prices, statuses, facings
- [ ] Working hours + enquiry phone/email/address confirmation (screen 06)
- [ ] Founder quote / milestone years for About timeline
- [ ] Naming confirmations (e.g. Chandrayan signboards read "Anna Avenue, Thiruppalai Nadu")
- [ ] Price display policy per project (number vs price-on-request)

## Critique of 01-home (2026-11-13, baseline score 24/32)
- [x] P0: mobile nav — ≤900px header now shows logo + gold "Book a Visit" pill + circular tap-to-call icon (`.nav-visit`/`.nav-call` in meridian.css + 01-home header markup); desktop header verified unchanged at 1440px
- [x] P0: mobile geo pins — ≤640px pins are labeled cards (pin glyph above, tip text as static caption); all 4 marker links now have accessible names (WCAG 1.1.1 fixed); verified via 390px screenshots
- [x] Decisions recorded in DESIGN.md: DTCP stamp ink → Verified Green; uniform carousel classified as drift (contradicts BRIEF asymmetric mandate); desktop nav carries no gold CTA
- [x] P1: showcase redesign — asymmetric (Landmark City large lead with 2×2 spec ledger, Srivelan/Chandrayan compact photo-row cards, OMR dashed footnote strip with badge + Enquire); verified 1440/900/390, no overflow, no new detector findings (`.sh-grid`/`.sh-lead`/`.sh-side`/`.sh-soon` in 01-home.html)
- [x] P1 (partial): DTCP stamp recolored to Verified Green (meridian.css `svg.stamp`)
- [x] User-reported: duplicate "Book a Site Visit" — the fixed header's gold CTA duplicated every in-page Book moment while scrolling; removed from the nav on all 5 screens (desktop). Booking now: hero CTA, closing CTA band, contact page, mobile header pill. DESIGN.md updated.
- [x] P1: verification misfires (remaining) — hero now captioned "Concept render — site photography to follow" under the CTA row (visible on all widths)
- [x] P2: ledger micro-typography system pass — spec labels 11px (meridian `.specs span` + DESIGN.md ramp), brass small-text → new `--brass-ink` token `#856517` (4.94:1 paper / 5.04:1 badge / 5.42:1 card; applied to eyebrows, `.req`, reserved badges, stat units, FAQ indices, link hovers, step numerals; 03 galnav/gallery too), `.btn-text` tap target ~46px via padding/margin compensation, stat unit floor `max(.55em, 12px)`
- [x] User-reported: sh-card status row divider removed — `.sh-grid .foot` border-top 0, whitespace only (verified 1440/390)
- [x] Cleanup: `.teaser .plot` dead JS + `.teaser-grid` media ref removed from 01-home
- [x] User request: value-list numerals fill gold on scroll-in/hover; buying-process circles fill gold on segment hover (01-home, DESIGN.md "Step & Index Numerals")
- [x] `.sh-loc` contrast: new token `--ink-subtle` `#626c7a` (3.95:1 on white — passes large/sub-label use; meridian.css + `.sh-loc`)
- [x] Footer muted text: the flagged item was 01's `Staff login` at `opacity:.45` (4.40:1). All 5 public-screen footers now use shared `.site-footer a.staff { opacity:.75 }` (≥4.5:1, matches 02's earlier .75 decision) — 01 .45→.75, 03/05/06 .5→class, 02 inline→class. Verified computed opacity .75 @?v=13
- [x] Gold-button hover 2.4:1: real culprit was `.step:hover .num` white-on-gold on 01 (white on #c9a24b = 2.40:1; `.btn-gold` itself is navy-on-gold 7.09/8.36 — fine). Hover text now `var(--navy-deep)` (7.09:1), matches button vocabulary. Verified in CSSOM @v=13
- [ ] Cleanup: `.hscroll`/`.hcard` CSS in meridian.css still used by 03-project-detail.html — verify before removing
- [ ] Minor (remaining): `#b08c3f` off-token brass, staff-login IA, `aria-label` on `.markers` (generic div — move to role/heading), footer h2→h4 skip, counter IntersectionObserver, project imgs lack width/height (CLS contained by fixed `.ph` frame)
- [x] Minor batch (this pass): `.teaser .plot` dead JS + `@keyframes cue` removed; skip link (`.skip-link` in meridian.css + markup after `<body>`) + `<main id="main">` on 01–06; `.hero-phone` tel tap target padded to ~44px + decorative phone SVGs `aria-hidden` (01); user-reported CTA-band side padding on mobile fixed (`.cta-band .inner` `88px 0` → `88px 28px` — affected 01/02/03/05)
- Corroboration (Assessment B static scan, `del_mweew8ac_q8zc`): duplicate IDs 0, all imgs have alt, all links named at desktop — matches the delivered critique
- Apply the same mobile header markup (`.nav-visit` + `.nav-call`) to screens 02–06 (CSS in shared meridian.css already hides them by default — markup is the only per-page change)
- Full report: `.impeccable/critique/2026-11-13T03-30-25Z__lavish-screens-01-home-html.md`

## Critique of 02-projects (2026-11-13, baseline score 21/32)
- [x] Critique run + full report persisted: `.impeccable/critique/2026-11-13T04-39-11Z__lavish-screens-02-projects-html.md` (dual-agent + parent DOM/visual verification + parent CLI detector)
- [x] P0: CTA band `.veil` div added (02 after `.bg`); `btn-ghost-dark` restyled shared → solid white chip + navy text (meridian.css:91) — translucent ghost composited 3.41:1 over the .55 veil end; used by 02/03/05, verified 1440px + 390px, band matches 03/05 grammar
- [x] P0: no-JS fail-safes in 02 — `<noscript>` block (rows `opacity:1`; header solid paper + navy text + red logo) + `@media (prefers-reduced-motion: reduce)` for `.rv`; verified with Chromium `scriptEnabled=false` (4 rows opacity 1, header rgb(246,244,239)/navy, brand-red.png)
- [x] P1: contradictions — row 01 badge "3 plots remaining*" vs dot strip "15 of 24 available*" (02:152 vs :162); footer "All layouts DTCP approved" vs row 04 "In process" (02:311 vs :253)
- [x] P1: quiet row deck links to 03-project-detail.html (Landmark) — wrong destination for unlaunched project (02:238)
- [x] P1: mobile spec-cell wrap at ≤560px — flex-row label/value squeeze caused mid-word label wraps (EXTEN/T*) + 4-line values at 390px; cells now stack label over value (`@media (max-width:560px) .prow .specs > div { flex-direction:column; align-items:flex-start; gap:3px }`, 02:109–112). Earlier "not reproducible" verdict was superseded — original probe missed flex-item top offsets. Verified 390px (labels one line, values wrap at word boundaries, no overflow)
- [x] P1: staff-login link 3.89:1 (02:299, opacity .5) — detector-corroborated
- [x] P1: availability breakdown hover-only (title attr only) + no dot legend on page
- [x] P1: heading skip h2→h4 (footer)
- [ ] P2: scarcity pill color drift (green/neutral/brass), asterisks not `.req` brass, inconsistent location starring, quiet row not quiet, no nav is-active, no plot-map accelerator, ghostnum aria-hidden, img width/height, prefers-reduced-motion, frost hover snap, .v nowrap @1000px, footer h4 12px, sample-flag overlap, rationale drift (two R9s, stale paddings), dead code

- 02 fix round (all P0/P1s from critique applied in one pass): contradictions, quiet-row link, staff-login contrast, availability legend/aria, footer h4, mobile spec stack. Verified: desktop 1440 + mobile 390 screenshots, detector clean, no-JS fail-safe unaffected. No new rationale entries needed (all fixes documented in code comments)
- 02 Lavish live-feedback rounds 6–9 (R20–R23 in file rationale list): md deck containment + rounded mat, white-border prints, hover zoom scoped to the deck itself (row hover = shadow only, tap-zoom dropped), mobile deck breakout with first-card top kept inside, deck radius 16px, md column media 240→300px. Verified sw=iw at 390/768/1440, hover rules in CSSOM

## Critique of 03-project-detail (2026-11-14, full report persisted)
- [x] Report: `.impeccable/critique/2026-11-14T02-48-14Z__lavish-screens-03-project-detail-html.md` (isolated design-review sub-agent → CLI detector 14 findings → browser verification 1440/390/scripts-off). Applied as on-page R14.
- [x] P0 no-JS dead page: all 11 `.rv` opacity 0 + glass header + empty svg without JS → `html.js` gate + inline class script + `<noscript>` header/map fallback; verified via injected noscript CSS (paper header rgb(246,244,239), navy links) and scripts-stripped copy (0 hidden .rv). NOTE: sed-stripped file:// copies don't activate `<noscript>` (browser still has scripting on) and relative `meridian.css?v=4` breaks outside .lavish/screens — absolute-rewrite needed for no-JS probes.
- [x] P1 self-contradicting sample data: "3 plots remaining*" → "15 of 24 plots available*", "600–2,400 sq.ft*" → "1,200 – 2,400 sq.ft*" (factsband aligned to PLOTS array/legend)
- [x] P1 desktop header CTA: `.nav-visit` now shows ≥901px (meridian.css, was mobile-only ≤900px); verified on 03 + 02 at 1440px
- [x] P1 map AT/keyboard: svg role="img"→"group"; zoomctl `:focus-within` visible; Esc deselects (panel empty state verified); filter dims set aria-disabled + tabindex=-1 (Enter on dimmed ignored)
- [x] P2: touch-action none→pan-y (mobile scroll trap); zin clamp min viewBox 300 (was unbounded); sample-flag dismiss ✕ added (verified removed); footer h4→h3 (heading skip); duplicate </ol> + "R31"→"R13" comment; mapbar title 15→17px (flat type step); mobile factsband gap 14px
- [x] Gallery arrows end-of-scroll disabled state: 01-home `.hscroll-nav` already had it; ported to 03's `.galnav` (`.galnav[disabled]` opacity .35 + cursor default; JS sync on scroll/resize/init, inline onclick removed). Verified: prev disabled at 0, next disabled at max 1556, both enabled mid-scroll
- [x] "Book a Visit" vs "Book a Site Visit" naming — user chose "Book a Visit" everywhere: header pills (unchanged), 01/02/03/05 CTA bands, 06 h2 + submit button + title. Body prose may still say "site visit". DESIGN.md updated (supersedes BRIEF.md label)
- [ ] Build-stage (remaining): panel fine-print 3.9:1 on navy; CTA links carry plot/project context (query param); `.fixedhero .inner` 1280 vs container 1200; "ask us for the approval copy" interim CTA for DTCP number
- Rejected detector findings (deliberate/client-locked): cream-palette (brand surface), hero-eyebrow-chip (house pattern), dark-glow (subtle lift shadows), em-dash advisory

## 05-about hero photo (2026-11-13, user-supplied)
- [x] Hero photo → client-supplied concept render (`assets/about-hero.jpg`, 3168×1200 JPEG q82 from user's `~/Downloads/Gemini_Generated_Image_…2vsd.png`): bottom ~12% cropped to drop the render-tool watermark; mirrored horizontally (`-flop`) so the family + team member sit right-of-center clear of the left-aligned headline; `.ab-hero .bg` + `aria-label` updated (05:13/:161). Verified 1440px (faces clear of headline) + 390px (mobile header pill/call over photo, text legible under veil)
- [ ] Note: image is AI-generated while other heroes are photo/aerial — if the honesty grammar should caption it (like 01's "Concept render" tag), decide wording with client

## Critique of 05-about leftover pass (2026-11-14, canon fixes from user-provided founder story)
- [x] CANON DECISION: location is **Chennai**, not Madurai — user confirmed founder worked in Chennai and VSS layouts are Chennai-area. Fixed in 05 (lede + footer), 06-contact (form options + footer) and 08-admin-projects (project meta). Site-wide grep: zero Madurai references remain.
- [x] Founder story (user-provided, from receptionist/employee interviews): founder worked inside another real estate firm, saw unclear documentation leaving buyers worried/confused, left to found VSS, grew it bottom-up plot by plot; runs office like a family, modest/accessible, flexibility + freedom rare in real estate. Written into 05 split narrative + timeline beats (Origins*/Founded*/Today/Next*); pull-quote still awaits client wording.
- [x] Strip + timeline canon: 6→4 "Layouts delivered" (matches home stat), timeline Today h3 "4 layouts delivered".
- [x] Pending-dot fix: 5px dashed border fragmented into broken arcs at 20px dot size → solid gold ring at 40% opacity (color-mix). Verified desktop + mobile.
- [x] Verified: 1280px + 390px full-page screenshots (scroll-reveal forced for capture); Chennai in lede/footer, story text, strip=4, clean pending dots.
- [x] 05 critique P1 — undefined CSS tokens: --gold-deep (#856517, = --brass-ink), --line-strong (#cfc8b8), --warm (#efeadd) defined in meridian.css :root; used by 05/06/07/08/09. Verified computed values on 05 (proof band bg = warm sand, req text = brass). NOTE: 06/07/08/09 admin/contact screens got these tokens implicitly — not visually re-verified there.
- [x] 05 critique P1 — .sample-flag ribbon + .rationale section added to 05 (mirrors 01/02/03 pattern in meridian.css): rationale = borrowed/adapted/improved/client-input + R1–R5 revision log; flag = "story beats, years & proof numbers marked * await client supply". Verified: pill fixed bottom-right desktop, in-flow after rationale on mobile (≤700px), both render.
- [x] Gallery placeholder tiles: explicit SAMPLE flag covered by the ribbon (per-tile flag judged unnecessary — tiles already carry asterisked captions like "project name to be confirmed*")
- [ ] Timeline yr labels "Origins*"/"Founded*" are era labels awaiting real years — replace when client provides founding year/name

## Critique of 06–09 leftover screens (2026-11-14, dual-agent impeccable critique)
Reports: `.lavish/screens/critique/agent-A-report.md` (design direction, source-level) + `agent-B-report.md` (interactions/a11y, source-level) + parent visual verification screenshots (`06-A-desktop-verify.png`, `09-verify-desktop.png`, `09-verify-mobile.png`). Both sub-agents were sandboxed (no browser/screenshots); visual claims below verified by parent. Scores: 06 · 7.5/6 · 07 · 7/5 · 08 · 8/5 · 09 · 7/6 (A/B).
CORRECTION: Agent B's "canon violation" on 06 contact facts is REFUTED — PRODUCT.md:25 canon IS 044-42087913 + contact@vssassociates.com; all screens consistent (annotated in agent-B-report.md).
- [x] Fact-checks: `assets/meridian.js` does NOT exist but was referenced by 05+06 (404, dead tag) — REMOVED from both in the 06 fix round; 01-home links `meridian.css?v=10` vs all others `?v=4`; 01-home has ZERO Fraunces imports so its hero `.accent` (01:211 "with confidence.") renders plain — same flat-hero defect as 06
- [x] 06 P1: hero flat — FIXED: Fraunces @import + `.ct-hero h1 em` rule; h1 → "Walk the plot before <em>you buy it.</em>" — verified gold italic serif renders desktop + mobile, matches 05 grammar
- [x] 06 P1: dead submit — FIXED: `novalidate` form, JS validation (name non-empty, phone ≥10 digits, consent ticked) with inline `.ferr` messages + `aria-invalid` + focus first invalid; success state swaps h2 → "Request received" + personalized panel (role=status, form hidden, fs-intro hidden); skip-link target restored (`<main class="ct-body" id="main">`, stray `</main>` removed). Verified: empty submit → 3 errors + focus on f-name; filled → success (Ravi / 98410 12345); desktop + mobile
- [x] 06 P1: project-name drift — FIXED: select now Landmark City / Srivelan Nagar / Chandrayan Nagar / Not sure yet / Other — matches footer canon
- [x] 06 P2 — FIXED: required + aria-required on name/phone/consent; hours → "Mon – Sat, 8.00 – 18.00" (matches footer); dead meridian.js tag removed (also from 05). Left as-is (intentional): undefined `eyebrow--gold`/`side-half` classes are harmless dead markup, removed opportunistically later
- [x] 07 P1+P2 (fix round, verified desktop 1440 + mobile 390): brand-white.png logo on navy (red variant gone); real auth flow — validation (.ferr inline errors + aria-invalid + red borders + focus-first-invalid), loading state ("Signing in…" disabled) then navigates to 08-admin-projects.html (silent "Continue →" bypass removed); "Forgot password?" is a real button toggling an honest demo note (role=status); show-password toggle (aria-pressed + label swap); tagline/backsite lifted rgba .4/.45 → .66 (legible, screenshot-verified)
- [x] 07 empty-submit → both errors + focus in username (screenshot); valid submit → navigation confirmed (get url = 08-admin-projects.html); show-password verified via DOM (type=text, pressed=true, label="Hide password")
- [x] 08 P1 (fix round, verified): mobile adminbar — nav wraps to own row ≤960px instead of display:none; project list items are real `<button class="pitem">`; 12 image buttons labeled (aria-label Move left/right/Remove) + wired; publish guard connected (confirm() on placeholder fields → demo toast); "Add new project" is a button
- [x] 08 P2 partial: dead .strow CSS removed; demo-toast convention noted for remaining inert controls (Enquiries, Save) — demo-only mockups, fine-tune later
- [x] 09 P1 (fix round, verified): canvas geometry gap fixed (GAP 5→4, maxRight 960 — no road-strip overlap); sold = navy #14314f / reserved #f3e7c9 (matches public palette); mobile map overflow-x scroll @ min-width 880px; inspector fields re-render per selection (no stale A5 data); status changes persist to PLOTS array; plots are keyboard buttons (tabindex, role=button, aria-label, arrow-key nav), svg role=group + arrow-key hint
- [x] 09 P2 partial: stepper says "site-plan.jpg placeholder*" (consistent with mapfoot "survey drawing awaited*"); remaining dead controls carry the demo-toast convention — demo-only, fine-tune later
- [x] Cross-screen (done): plot status palette unified (09 sold=navy/reserved=#f3e7c9 matches public); 01 Fraunces @import + `.hero h1 .accent` italic-serif rule — "with confidence." renders gold italic Fraunces, matches 05/06 grammar (screenshot-verified); all 8 screens now link `meridian.css?v=11` (was 1×v10 + 7×v4); 08 admin list "Srivelan" → "Srivelan Nagar" (canon)
- [x] Radius scale + admin chrome consolidation (done + verified): meridian.css tokens now `--r-sm:8px` / `--r-md:12px` / `--r-lg:16px`; all literal radii mapped (14/18/16px → --r-lg, 8/9/10/12px → --r-sm/--r-md; pills 999px, circles 50%, micro 2–6px kept); admin chrome (.adminbar + .bar + .brand-logo + .label + nav + responsive wrap @960px) lifted into meridian.css, duplicates stripped from 08+09 (08 keeps .who, 09 keeps .crumb); version param → ?v=12 on all 8 screens
- [x] Visual check after consolidation: 08 desktop (full render incl. shared chrome + list + editor), 09 desktop, 09 mobile 390 (adminbar wraps logo/label/who row 1, nav row 2), 02/05/06 desktop (radii + Fraunces accents intact). NOTE: single blank screenshot of 08 was a render-timing artifact — re-shot with 700ms wait, fine
- [x] Cross-screen (deferred polish, done): dead `eyebrow--gold` removed (06 + a second one on 05) and `side-half` removed (06); footer staff-link contrast + gold hover fixed (see 01 section); shared `.site-footer a.staff` rule in meridian.css; css version param bumped ?v=12→?v=13 on all 8 screens (browser had cached the old v=12 response)

## Implementation iteration 1 — Go + Tailwind + Datastar stack, 01-home (2026-11-15)

**Completed**
- [x] Stack wired: `main.go` (net/http + html/template + datastar-go v1.2.2), Tailwind v4 CLI watch (`assets/css/input.css` → `assets/css/index.css`), Air autoreload; design tokens + meridian components ported into the Tailwind build
- [x] Template architecture: `assets/html/layout.html` + `partials.html` (header/footer/stamp) + per-page `main` fragments cloned per page
- [x] Fragment navigation via Datastar SSE: nav links `@get` the server, server patches `<main>` and pushes URL (works in headed browser — user verified clicks)
- [x] 01-home converted to the artifact's verbatim markup/classes; all content copied literally from the artifact (an earlier draft had fabricated address/footer text — caught and replaced)
- [x] `assets/js/app.js`: stat-band counters + scroll reveals ported from artifact; sample-data flag, skip-link kept
- [x] Header scroll-solid fixed — expressions must use `$` signal prefix (`$_scrolled = ...`) in `data-init`/`data-on:scroll__window__passive`
- [x] desktop 1440 verified vs artifact (hero, statband, showcase, geo band, FAQ, CTA, footer)
- [ ] User re-verify after Tailwind migration + reformat (whitespace reflow in h1/br, stat spans)

**Deliberate deviation from locked artifact**
- [x] Geo-band "Explore by location" button: white ghost → gold (user request this session; artifact meridian.css `.btn-ghost-dark` is white/navy)

**Remaining before calling home fully done**
- [x] Mobile 390px pass verified via agent-browser (user upgraded it past the version gate): no horizontal overflow (scrollWidth 390), statband 2-col grid, hero, footer, solid header on scroll
- [x] Functional QA via agent-browser: zero console errors; fragment nav proven no-reload (JS marker survives click) with URL + title sync; scroll-solid header (`site-header--solid` at y>60); FAQ details toggle; counters land exactly on 10/600/120/4/100
- [x] Fixed: `document.title` now updates on fragment navigation (`renderFragment` pushes title in the same `ExecuteScript` as pushState)
- [x] djlint clean with `--configuration ~/.djlintrc` (golang profile): 3 remaining H021 are deliberate inline styles (artifact bg-image divs ×2, dynamic marker left/top); dead scaffold index.html removed; meta description added; templates reformatted per config (single_attribute_per_line, sorted attrs)

**Notes for next iterations**
- Official docs still point at bundle `@v1.0.5`; our local bundle is byte-identical — do NOT "upgrade" blindly; current colon syntax (`data-on:click`, `data-init`) verified compatible with it
- agent-browser quirks learned: batch `eval --stdin` silently drops payloads (use standalone eval calls); refs go stale after any navigation — snapshot fresh before clicking
- 02-projects → 03-detail → 05 → 06 → 07–09 conversion order agreed; admin screens remain static placeholders until data modelling

## Implementation iteration 2 — Tailwind-native migration (2026-11-15)

- [x] input.css reduced 893 → ~215 lines: only @theme tokens (+ cuedot keyframes), @layer base resets, and a minimal @utility set (btn family, brand-logo, reveal, marker)
- [x] All component CSS moved into templates as native Tailwind utilities (layout, partials, home, placeholder); footer inline styles converted too
- [x] Radius tokens added to @theme (rounded-sm 8 / md 12 / lg 16 = meridian --r-*)
- [x] docs/STYLE.md — meridian→Tailwind migration guide for remaining pages
- [x] User visual re-verification (headed browser) after the migration — spacing/type must be pixel-identical
- [ ] djlint reformat note: home.html grew ~2× from single_attribute_per_line; class strings are untouched

## Implementation iteration 3 — 02-projects + post-verification fixes (2026-11-15)

- [x] 02-projects implemented (hero + 4 data-driven rows + fragment nav), verified desktop/mobile vs artifact, committed `4535e32`…`0d570f1`
- [x] agent-browser functional QA: zero console errors; fragment nav proven no-reload; both pages render
- [x] sm call-button fix: base `a{color:navy}` beat inherited header color — `max-md:text-current` on the call link (white over hero glass, navy on solid header)
- [x] sample-data note: ≤700px renders as 2nd bullet under the DTCP line in footer legal row; ≥701px keeps the fixed dark pill (matches artifact)
- [x] lg home grid fix: `.sh-lead{grid-row:span 2}` was missing — `row-span-2 max-lg:row-span-1` on the lead card; verified 594px lead vs 2×287px side cards
- [x] fragment-nav scroll-to-top: `window.scrollTo({top:0,behavior:'instant'})` appended to the server's post-patch ExecuteScript (main.go)
- [x] NOTE: air only watches go/html — the tailwindcss --watch process must run alongside (`make dev`); after class edits with no rebuild, `row-span-2` silently missing from index.css cost a debug round
- [x] sm marker hover (client request m01173): max-sm caption hover no longer translates left — feedback via size instead (`b` 14→15.5px, `span` 12→12.5px, transform:none); verified via computed styles + screenshot on 390px viewport
- [x] reveal replay after fragment nav: morphed `.reveal` nodes (attribute-mutation branch in app.js) are re-observed with hidden state committed instantly (transition:none + reflow) so the 0→1 rise replays from opacity 0 — verified opacity ramps 0→0.02→0.33→0.74→0.92→1.00 over ~700ms; app.js cache-bust → v=4
- [x] `/projects/` trailing slash now 308-redirects to `/projects` (was 404); deep links unaffected; verified via curl (air rebuilt on main.go change)
- [x] `.djlintrc` added (profile golang, ignore H021 for intentional per-element CSS vars / reveal delays, max_line_length 200) — djlint clean on all 5 templates; committed `90eec32`, pushed
- [x] /projects/ trailing-slash 404 fixed in main.go (redirect)

## Next iteration — derive architecture from approved designs

- [ ] Present the 9 artifacts + rationale to client for approval
- [ ] Component structure: shared Go templates (header/hero/ledger/gallery/plot-map/gallery/marker) derived from `meridian.css` classes
- [ ] Data model: Project, Plot, PlotStatus, Enquiry, AdminUser + plot layout storage (polygon vs grid decision)
- [ ] API boundaries: /api/projects, /api/plots (public read), /admin CRUD, enquiry submit, layout publish
- [ ] Datastar interactions: plot map (signals for selection/filter/zoom), counters, header solid-on-scroll, form submit
- [ ] Backend + DB schema once screen designs approved

## Follow-on visual work (optional, after client review)
- [ ] Editorial/brand imagery generation to supplement proof photos (hero-grade, no factual claims)
- [ ] Admin: enquiries inbox screen (nav entry exists in 08/09 mockups)
- [ ] 404 / legal pages
