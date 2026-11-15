# Meridian → Tailwind v4 style guide

How to translate the approved artifacts' CSS (`.lavish/screens/meridian.css`)
into native Tailwind utilities in the Go templates. **Styling lives in the
templates**; `assets/css/input.css` holds only theme tokens, global resets,
animation, and a tiny `@utility` set.

## Where things live

| Layer | Contains | Rule of thumb |
|---|---|---|
| `@theme` (input.css) | colors, shadows, breakpoints, radii, fonts, `--animate-*` | every new token goes here, never in markup |
| `@layer base` (input.css) | body defaults, resets | global only |
| `@utility` (input.css) | multi-state / pseudo-element composites reused across pages | only if utilities genuinely can't express it (see list below) |
| templates | everything else | single source of styling |

Current `@utility` whitelist: `btn`, `btn-gold`, `btn-navy`, `btn-ghost-dark`,
`btn-outline`, `btn-text`, `btn-sm`, `brand-logo`, `reveal`, `marker`.
Do not extend this list for anything a utility chain can express.

## Token map

```text
color tokens            navy navy-deep gold gold-soft brass card paper line
                        line-strong warm ink ink-soft ink-faint
                        green green-bg red red-bg
usage                   bg-navy  text-brass  border-line  bg-green-bg …
opacity modifiers       text-white/82  bg-navy/92  border-white/12
shadows                 shadow-card  shadow-lift
radii (overridden)      rounded-sm=8px  rounded-md=12px  rounded-lg=16px
breakpoints             sm=640  md=900  lg=1020  xl=1280
                        mobile-first:  md: ≥900   max-md: <900
fonts                   font-sans (Manrope)  font-serif (Fraunces italic 600)
animation               animate-cuedot (hero scroll cue)
```

Note: Tailwind's default type scale differs from the artifact's px values —
use arbitrary values to match the artifact exactly (`text-[17.5px]`,
`py-[5.5px]`, `tracking-[.16em]`, `duration-[250ms]`).

## Core translation patterns

| meridian.css | Tailwind |
|---|---|
| `.container` | `mx-auto max-w-[1200px] px-7` |
| `.section` / `--tight` / `--white` | `py-24 max-md:py-16` / `py-16 max-md:py-16` / `+ bg-card` |
| `.eyebrow` (+`--light`) | `m-0 mb-[14px] text-xs font-bold uppercase tracking-[.16em] text-brass` (light: `text-gold-soft`) |
| `.h-section` | `m-0 mb-4 text-[clamp(28px,3.4vw,40px)] font-extrabold leading-[1.12] tracking-[-.015em] text-navy` |
| `.h-display` | `m-0 text-[clamp(40px,5.4vw,64px)] font-extrabold leading-[1.06] tracking-[-.02em] tabular-nums` |
| `.lead` | `max-w-[62ch] text-[18px] leading-[1.65] text-ink-soft` |
| `.muted` / `.req` | `text-ink-soft` / `font-bold text-brass` |
| `.badge-*` | `inline-flex items-center gap-1.5 rounded-full px-3 py-1.5 text-[11.5px] font-bold uppercase tracking-[.08em]` + color pair (dtcp `bg-green-bg text-green`, status `bg-paper text-ink-soft border border-line`, available/reserved/sold analogously, navy `bg-navy/8 text-navy`) |
| `.statband` | wrapper `relative z-[5] -mt-[74px] max-sm:-mt-[40px]`; band `rounded-[2px] border-t-[3px] border-gold bg-card shadow-card`; grid `mx-auto grid max-w-[1200px] grid-cols-5 gap-5 px-7 py-[34px] max-md:grid-cols-2` |
| `.stat .num/.unit/.lbl` | see home.html; `min-h-[44px]` keeps rows aligned |
| `.spec-row` | row `flex items-baseline justify-between gap-3 border-t border-line py-[5.5px] [&:last-child]:border-b`; label `text-[11px] font-bold uppercase tracking-[.12em] text-ink-soft`; value `whitespace-nowrap text-[13px] font-bold text-navy` |
| `.pcard` | `group bg-card flex flex-col overflow-hidden rounded-md shadow-card transition-[transform,box-shadow] duration-[250ms] hover:-translate-y-1 hover:shadow-lift` |
| `.pcard .ph` + photo zoom | anchor `relative block overflow-hidden after:absolute after:inset-0 after:bg-[linear-gradient(180deg,rgba(8,29,51,0)_55%,rgba(8,29,51,.42))] after:content-['']`; img `h-full w-full object-cover transition-transform duration-500 group-hover:scale-[1.045]` |
| `.sechead` | `mb-12 max-w-[640px]` (center: `+ mx-auto text-center`) |
| `.faq` | border `border-t border-line` on list, `group border-b border-line` per `details`; summary `flex cursor-pointer list-none items-center gap-[18px] px-1 py-[22px] text-[17px] font-bold text-navy [&::-webkit-details-marker]:hidden [&::after]:content-['+'] [&::after]:ml-auto [&::after]:text-[22px] [&::after]:text-gold [&::after]:transition-transform group-open:[&::after]:rotate-45` |
| `.cta-band` | section `relative overflow-hidden text-white`; bg `absolute inset-0 bg-[url(...)] bg-cover bg-center opacity-[.38] saturate-[.9]`; veil `absolute inset-0 bg-[linear-gradient(90deg,rgba(8,29,51,.88),rgba(8,29,51,.55))]` |
| `.field` (forms) | label `mb-[7px] block text-[12.5px] font-bold uppercase tracking-[.06em] text-ink-soft`; input `w-full rounded-sm border border-line bg-card px-[15px] py-[13px] text-[15.5px] font-medium outline-none transition-[border-color,box-shadow] focus:border-gold focus:shadow-[0_0_0_3px_rgba(201,162,75,.18)]` |
| `.plot--*` | SVG-specific — keep as `@utility` when 03-detail is converted |
| `.hscroll*` | see `docs` when 02-projects is converted (`snap-x snap-proximity` etc.) |

## Pattern cookbook

- **Pseudo-elements** → `before:`/`after:` variants; always pair with
  `before:content-['']` when the original had `content: ''`.
- **Ancestor state** → `[.lit_&]:…`, `[.site-header--solid_&]:…` (JS/Datastar
  toggles the class on an ancestor).
- **Self state** → `[&.site-header--solid]:bg-paper` (class on the same element).
- **`details[open]`** → `group` on `<details>`, `group-open:` on children.
- **Hover zoom on child** → `group` on card, `group-hover:scale-[1.045]` on img.
- **Overriding a `@utility btn` metric** → important suffix: `px-[18px]!`,
  `mt-[18px]!` (shorthand-vs-longhand cascade is not order-guaranteed).
- **prefers-reduced-motion** → `motion-reduce:` variant (baked into `.reveal`).
- **JS/Datastar hooks to preserve in markup**: `.count[data-target]`,
  `.reveal` (+JS adds `.in`), `.prop` (+JS adds `.lit`), `.arr`,
  `site-header--glass/--solid` (+`data-class:site-header--solid`).
- **Dynamic Go values** (marker left/top) stay as inline `style` — acceptable
  exception, they are data not design.

## Header/footer quick map

- glass header: `fixed inset-x-0 top-0 z-50 text-white transition-[background,box-shadow] duration-300` + `[&.site-header--solid]:bg-paper [&.site-header--solid]:text-navy [&.site-header--solid]:shadow-[0_1px_0_var(--color-line)]`
- bar scrim (fades when solid): on the bar div, `before:` gradient utilities + `[.site-header--solid_&]:before:hidden`
- nav link underline: `after:right-full → hover:after:right-0` transition
- footer headings: `mb-4 mt-0 text-xs uppercase tracking-[.16em] text-gold-soft`
- footer link: `inline-block py-0.5 text-white/82 hover:text-gold-soft` (staff `+ opacity-75`)

## Process for converting a new artifact page

1. Copy the artifact's markup into a page template; keep Datastar attrs.
2. Replace every meridian class with the utilities above (grep this file
   first — most are already mapped).
3. Anything repeated across pages that needs pseudo-elements/states → propose
   a new `@utility` in review, don't invent silently.
4. Rebuild CSS (`~/bin/tailwindcss -i assets/css/input.css -o assets/css/index.css --minify`),
   `go build`, visually verify 1440 + 390 against the artifact.
