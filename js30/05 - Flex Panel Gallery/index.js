// @ts-check
"use strict";

document.querySelectorAll(".card").forEach(ele => {
  ele.addEventListener("click", () => ele.classList.toggle("open"));

  ele.addEventListener("transitionend", (evt) => {
    if (/** @type TransitionEvent */ (evt).propertyName === "flex-grow")
      ele.classList.toggle("active"); 
  });
});
