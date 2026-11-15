# Website Revamp — Agent Rules

## Core principle

This is an iterative website redesign.

Do not attempt to complete the entire website in one iteration.

Before making significant implementation changes:
1. Understand the existing design decision.
2. Identify the smallest useful increment.
3. Implement it.
4. Run the available checks.
5. Inspect the result visually.
6. Report what changed and what should be done next.

## Design

Do not invent:
- company facts
- project facts
- prices
- approvals
- testimonials
- statistics
- awards
- amenities
- locations
- legal claims

When required information is unavailable, use an explicit placeholder such as:

`CLIENT INPUT REQUIRED`

Do not silently fabricate marketing copy.

## Architecture

Prefer reusable data-driven components over project-specific templates.

A new project should be addable through project data rather than requiring a new page template.

Do not introduce a framework, dependency, server, process or architectural pattern unless it is necessary for the current iteration.

Frontend should remain primarily server-rendered using Go + Tailwind + Datastar.

## Iteration discipline

Work on one clearly defined objective at a time.

Do not redesign previously approved sections without explaining why.

Do not refactor unrelated code while implementing a feature.

Do not implement future requirements speculatively.

## Visual verification

After meaningful UI changes:
- run the application
- inspect the rendered page
- check desktop and mobile
- verify images, spacing, typography and responsive behaviour

Do not assume the implementation looks correct merely because it compiles.

## Todo tracking

Maintain `TODO.md`.

Every iteration should:
- mark completed items
- add newly discovered work
- separate blockers from normal TODOs
- keep the next iteration small and concrete

At the end of each iteration report:

1. Completed
2. Verified
3. Remaining
4. Recommended next step

Do not proceed into a large unrelated task automatically.
