// @ts-check
"use strict";

/** @type {string[]} */
const keys = [];
const secret = "deesa";

document.addEventListener("keydown", evt => {
  if (keys.push(evt.key) > 5) keys.shift();
  if (keys.join('') === secret) {
    console.log("DING DING!")
    cornify_add();
  }
});
