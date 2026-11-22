// @ts-check
"use strict";

const vids = Array.from(document.querySelectorAll(".videos li"));

/** @typedef Time 
  * @property {number} hours
  * @property {number} minutes
  * @property {number} seconds
*/

const aggregated = vids.reduce((/** @type {Time} */ agg, /** @type {Element} */ ele) => {
  const timeStr = /** @type HTMLElement */ (ele).dataset.time;
  if (!timeStr) return agg;

  const [min, sec] = timeStr.split(":").map(parseFloat);
  agg.seconds += sec;
  [agg.minutes, agg.seconds] = [agg.minutes + Math.floor(agg.seconds/60), agg.seconds % 60];

  agg.minutes += min;
  [agg.hours, agg.minutes] = [agg.hours + Math.floor(agg.minutes/60), agg.minutes % 60];

  return agg;

}, {hours: 0, minutes: 0, seconds: 0});

console.log(aggregated);
