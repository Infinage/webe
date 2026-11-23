// @ts-check
"use strict";

const span = document.createElement("span");
span.classList.add("highlight");

const main = document.querySelector("main");
main?.appendChild(span);

document.querySelectorAll("a").forEach(ele => {
  ele.addEventListener("mouseenter", evt => {
    if (!(evt.target instanceof HTMLElement)) 
      throw "Expected an HTMLElement";

    const {offsetHeight: h, offsetWidth: w} = evt.target;
    [span.style.width, span.style.height] = [`${w}px`, `${h}px`];
    span.style.translate = `${evt.target.offsetLeft}px ${evt.target.offsetTop}px`;
  });
});
