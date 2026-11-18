//@ts-check
"use strict";

window.addEventListener("keydown", ({key}) => {
  if (!["a", "s", "d", "f", "g", "h", "j", "k", "l"].includes(key)) return;

  const ele = document.querySelector(`[data-key=${key}]`)
  const audio = ele?.querySelector("audio")
  if (!ele || !audio) return;

  // Highlight the corresponding drum
  ele.classList.add("active");

  // Play the corresponding sound from start
  // playing an already started audio does nothing
  audio.currentTime = 0;
  audio.play();
});


/**
  * @param {Event} evt
  */
function removeTransition(evt) {
  // Satisfy linter type checks
  if (!(evt instanceof TransitionEvent)) return;
  if (!(evt.target instanceof HTMLElement)) return;

  if (evt.propertyName !== "transform") return;
  evt.target?.classList.remove("active");
}

const keys = Array.from(document.querySelectorAll(".drum"));
keys.forEach((ele) => ele.addEventListener("transitionend", removeTransition))
