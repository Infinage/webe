// @ts-check
"use strict";

const nav = document.querySelector("nav");
if (!(nav instanceof HTMLElement)) throw "Nav tag not found";

const hero = document.querySelector("header"); 
if (!(hero instanceof HTMLElement)) throw "Hero heading not found";

const obs = new IntersectionObserver(entries => {
  entries.forEach(entry => {
    if (!entry.isIntersecting) {
      nav.classList.add("sticky");
      document.body.style.paddingTop = `${nav.offsetHeight}px`;
    } else {
      nav.classList.remove("sticky");
      document.body.style.paddingTop = "";
    }
  }) 
});

obs.observe(hero);
