// @ts-check
"use strict";

const observer = new IntersectionObserver(entries => {
  entries.forEach(entry => {
    const img = entry.target.nextElementSibling;
    if (!img || !(img instanceof HTMLImageElement)) 
      throw "Expected image";
    entry.isIntersecting?
      img.classList.add("visible"): 
      img.classList.remove("visible");
  })  
}, { threshold: 0.25 });

document.querySelectorAll(".animate-trigger").forEach(ele => observer.observe(ele));
