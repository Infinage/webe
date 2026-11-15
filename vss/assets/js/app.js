// Site behaviours shared across pages.
//
// Datastar handles fragment navigation, the header solid-on-scroll
// state, hero scroll fades and the site-plan signals (selection,
// filters). This file covers the small imperative helpers that would
// otherwise need per-element signal plumbing: animated counters,
// reveal-on-scroll, the value-numeral flourish, the site-plan
// viewBox math (zoom/pan/fit) and the gallery carousel. All are direct
// ports of the approved artifact scripts.

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
  root.querySelectorAll('#plotmap').forEach(initMap);
  root.querySelectorAll('.galwrap').forEach(initGallery);
  if (root.matches('#plotmap')) initMap(root);
  if (root.matches('.galwrap')) initGallery(root);
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

// ---------- interactive site plan (03 project detail) ----------
// Zoom / pan / fit and focus-on-select: direct port of the artifact
// script. Selection itself runs on Datastar signals ($sel) — this is
// only the viewBox math, which stays imperative on purpose.
const initMap = (svg) => {
  if (!svg.matches('#plotmap') || svg._vssMap) return;
  svg._vssMap = true;
  const zoomctl = svg.closest('.mapwrap')?.querySelector('.zoomctl');
  const zin = document.getElementById('zin');
  const zout = document.getElementById('zout');
  const zfit = document.getElementById('zfit');
  if (!zoomctl || !zin || !zout || !zfit) return;

  let ztHide;
  const zoomLive = () => {
    zoomctl.classList.add('zoomctl--live');
    clearTimeout(ztHide);
    ztHide = setTimeout(() => zoomctl.classList.remove('zoomctl--live'), 2500);
  };

  let vb = { x: 0, y: 0, w: 1060, h: 430 };
  const setVB = () => svg.setAttribute('viewBox', `${vb.x} ${vb.y} ${vb.w} ${vb.h}`);
  const clampVB = () => {
    vb.x = Math.min(Math.max(vb.x, 0), 1060 - vb.w);
    vb.y = Math.min(Math.max(vb.y, 0), 430 - vb.h);
  };
  const vbZoom = (f) => {
    const nw = Math.min(Math.max(vb.w * f, 300), 1060);
    const nh = Math.min(Math.max(nw * 430 / 1060, 300 * 430 / 1060), 430);
    vb.w = nw; vb.h = nh; clampVB(); setVB(); zoomLive();
  };
  zin.onclick = () => vbZoom(.8);
  zout.onclick = () => vbZoom(1/.8);
  zfit.onclick = () => { vb = { x: 0, y: 0, w: 1060, h: 430 }; setVB(); zoomLive(); };
  ['wheel', 'pointerdown', 'pointermove'].forEach((ev) =>
    svg.addEventListener(ev, zoomLive, { passive: true }));
  zoomctl.addEventListener('pointerenter', () => {
    clearTimeout(ztHide);
    zoomctl.classList.add('zoomctl--live');
  });
  zoomctl.addEventListener('pointerleave', zoomLive);

  svg.addEventListener('wheel', (e) => {
    e.preventDefault();
    const nw = Math.min(Math.max(vb.w * (e.deltaY > 0 ? 1.12 : .9), 300), 1060);
    const nh = nw * 430 / 1060;
    const rect = svg.getBoundingClientRect();
    const mx = vb.x + (e.clientX - rect.left) / rect.width * vb.w;
    const my = vb.y + (e.clientY - rect.top) / rect.height * vb.h;
    vb.x = Math.min(Math.max(mx - (mx - vb.x) * nw / vb.w, 0), 1060 - nw);
    vb.y = Math.min(Math.max(my - (my - vb.y) * nh / vb.h, 0), 430 - nh);
    vb.w = nw; vb.h = nh; setVB(); zoomLive();
  }, { passive: false });

  // Tap-select runs on pointerup (artifact R10): capturing the pointer on
  // pointerdown retargets the browser click away from the plots, so a
  // capture is only taken once the pointer has actually travelled (drag).
  let drag = null, downPlot = null, downPt = null;
  svg.addEventListener('pointerdown', (e) => {
    drag = { x: e.clientX, y: e.clientY, vb: { ...vb }, captured: false };
    downPlot = e.target.closest?.('.plot') || null;
    downPt = { x: e.clientX, y: e.clientY };
  });
  svg.addEventListener('pointermove', (e) => {
    if (!drag) return;
    if (!drag.captured) {
      if (Math.hypot(e.clientX - drag.x, e.clientY - drag.y) < 6) return;
      drag.captured = true;
      svg.setPointerCapture(e.pointerId);
    }
    const rect = svg.getBoundingClientRect();
    const k = vb.w / rect.width;
    vb.x = Math.min(Math.max(drag.vb.x - (e.clientX - drag.x) * k, 0), 1060 - vb.w);
    vb.y = Math.min(Math.max(drag.vb.y - (e.clientY - drag.y) * k, 0), 430 - vb.h);
    setVB();
  });
  svg.addEventListener('pointerup', (e) => {
    // tap (< 6px travel): select the plot tapped, or deselect on an
    // empty-map tap — via the declarative click handlers on the plot /
    // <body> (layout.html), so signals stay in the Datastar attributes
    if (downPt && Math.hypot(e.clientX - downPt.x, e.clientY - downPt.y) < 6) {
      const target = downPlot || document.body;
      target.dispatchEvent(new MouseEvent('click', { bubbles: true, cancelable: true }));
    }
    downPlot = null; downPt = null; drag = null;
  });
  svg.addEventListener('pointercancel', () => { downPlot = null; downPt = null; drag = null; });

  // selecting a zoomed plot centres it (artifact scrollToPlot)
  window.vssFocusPlot = (n) => {
    if (vb.w >= 1060) return;
    const g = svg.querySelector(`g[data-n="${n}"]`);
    if (!g) return;
    const r = g.querySelector('rect');
    const cx = +r.getAttribute('x') + +r.getAttribute('width') / 2;
    const cy = +r.getAttribute('y') + +r.getAttribute('height') / 2;
    vb.x = Math.min(Math.max(cx - vb.w / 2, 0), 1060 - vb.w);
    vb.y = Math.min(Math.max(cy - vb.h / 2, 0), 430 - vb.h);
    setVB();
    zoomLive();
  };
};

// ---------- gallery strip (03) ----------
// one card per click, arrows disabled at either end — same vocabulary
// as the home-page hscroll nav
const initGallery = (wrap) => {
  if (!wrap.matches('.galwrap') || wrap._vssGal) return;
  wrap._vssGal = true;
  const gal = wrap.querySelector('.gal');
  const prev = wrap.querySelector('.galnav.prev');
  const next = wrap.querySelector('.galnav.next');
  if (!gal || !prev || !next) return;
  const step = () =>
    gal.querySelector('figure').getBoundingClientRect().width + 18;
  const sync = () => {
    prev.disabled = gal.scrollLeft <= 4;
    next.disabled = gal.scrollLeft >= gal.scrollWidth - gal.clientWidth - 4;
  };
  wrap._vssGalSync = sync;
  prev.onclick = () => gal.scrollBy({ left: -step(), behavior: 'smooth' });
  next.onclick = () => gal.scrollBy({ left: step(), behavior: 'smooth' });
  gal.addEventListener('scroll', sync, { passive: true });
  sync();
};
addEventListener('resize', () =>
  document.querySelectorAll('.galwrap').forEach((w) => w._vssGalSync && w._vssGalSync()));

initNode(document.body);
