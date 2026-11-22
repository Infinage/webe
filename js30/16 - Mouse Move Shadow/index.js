// @ts-check
"use strict";

document.addEventListener("mousemove", evt => {
  const root = document.querySelector(":root");
  if (!root || !(root instanceof HTMLElement)) throw "Root element not defined";

  const [mx, my] = [window.innerWidth/2, window.innerHeight/2];
  const [dx, dy] = [mx-evt.x, my-evt.y]

  root.style.setProperty("--pink-offset-x", `${-dx * 0.5}px`)
  root.style.setProperty("--pink-offset-y", `${-dy * 0.5}px`)

  root.style.setProperty("--cyan-offset-x", `${dx * 0.5}px`)
  root.style.setProperty("--cyan-offset-y", `${dy * 0.5}px`)

  root.style.setProperty("--blue-offset-x", `${-dy * 0.5}px`)
  root.style.setProperty("--blue-offset-y", `${dx * 0.5}px`)

  root.style.setProperty("--green-offset-x", `${dy * 0.5}px`)
  root.style.setProperty("--green-offset-y", `${-dx * 0.5}px`)
});
