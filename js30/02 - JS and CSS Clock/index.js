//@ts-check
"use strict"

const hH = document.getElementById("hour");
const mH = document.getElementById("minute");
const sH = document.getElementById("second");

/**
  * Tick moves the clock hands around, should be called at 1s interval.
  * @returns
  */
function tick() {
  if (!hH || !mH || !sH) throw "Missing essential markup";

  const d = new Date();
  const [h, m, s] = [d.getHours(), d.getMinutes(), d.getSeconds()];

  // 60s => 360 deg; 1s => 6 deg
  sH.style.transform = `rotate(${90 + (s * 6)}deg)`;

  // 60m (3600s) => 360 deg; 1s => 1 / 10 deg
  const ms = (m * 60) + s;
  mH.style.transform = `rotate(${90 + (ms * 0.1)}deg)`;

  // 12h (12 * 60 * 60) => 360 deg; 1s => 360 / (12 * 60 * 60) deg
  const h12 = h % 12;
  const hs = (h12 * 60 * 60) + ms;
  hH.style.transform = `rotate(${90 + (hs * 360 / (12 * 60 * 60))}deg)`;
}

setInterval(tick, 1000);
