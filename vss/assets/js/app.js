// Site behaviours shared across pages.
//
// Datastar handles fragment navigation, the header solid-on-scroll
// state and (later) forms and the plot map. This file covers the small
// animation helpers that would otherwise need per-element signal
// plumbing: animated counters, reveal-on-scroll and the value-numeral
// flourish. All are direct ports of the approved artifact scripts.

// animated counters — .count[data-target], run once on load, never
// replay on re-entry (locked in design.md)
const ease = (t) => 1 - Math.pow(1 - t, 3);
const runCount = (el) => {
  const target = +el.dataset.target;
  const t0 = performance.now();
  const dur = 1400;
  const tick = (now) => {
    if (!el.isConnected) return;
    const p = Math.min(1, (now - t0) / dur);
    el.textContent = Math.round(target * ease(p)).toLocaleString('en-IN');
    if (p < 1) requestAnimationFrame(tick);
  };
  requestAnimationFrame(tick);
};

// reveal on scroll — fail-safe: everything becomes visible after 2.5s
// so full-page captures / print never show blank sections
const rio = new IntersectionObserver(
  (es) => es.forEach((e) => {
    if (e.isIntersecting) {
      e.target.classList.add('in');
      rio.unobserve(e.target);
    }
  }),
  { threshold: .12 },
);
setTimeout(() => document.querySelectorAll('.reveal:not(.in)').forEach((el) => {
  rio.unobserve(el);
  el.classList.add('in');
}), 2500);

// value numerals: fill on direct interaction only — hover (CSS) or a
// brief flourish on click
const initNode = (root) => {
  if (root.nodeType !== 1) return;
  root.querySelectorAll('.count').forEach(runCount);
  root.querySelectorAll('.reveal').forEach((el) => rio.observe(el));
  root.querySelectorAll('.prop').forEach((el) => {
    if (el._litInit) return;
    el._litInit = true;
    el.addEventListener('click', () => {
      el.classList.add('lit');
      clearTimeout(el._litT);
      el._litT = setTimeout(() => el.classList.remove('lit'), 1200);
    });
  });
  if (root.matches('.count')) runCount(root);
  if (root.matches('.reveal')) rio.observe(root);
};
initNode(document.body);

// re-bind helpers after fragment navigation patches new content in.
// The morph REUSES matching DOM nodes (no addedNodes) but strips the
// .in class those nodes earned earlier — and rio unobserves after first
// reveal — so morphed-in-place reveals must be re-observed here.
new MutationObserver((muts) => {
  for (const m of muts) {
    m.addedNodes.forEach(initNode);
    if (m.type === 'attributes' && m.target.classList?.contains('reveal') &&
        !m.target.classList.contains('in')) {
      const el = m.target;
      // commit the hidden state instantly (the morph stripped .in while
      // the node was still on screen — without this the replay would
      // start from ~opacity 1 and show no rise)
      el.style.transition = 'none';
      void el.offsetWidth;
      el.style.transition = '';
      rio.observe(el);
    }
  }
}).observe(document.body, {
  childList: true,
  subtree: true,
  attributeFilter: ['class'],
});

// fragment navigation keeps the URL in sync via history.pushState
// (patched in by the server after each Datastar @get); back/forward
// re-render through a normal load for now
addEventListener('popstate', () => location.reload());
