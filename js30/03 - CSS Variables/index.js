//@ts-check
"use strict"

/** @type {HTMLElement | null} */ 
const root = document.querySelector(":root");
if (!root) throw "Missing root";

root.style.setProperty("--spacing", "10px");
root.style.setProperty("--blur", "10px");
root.style.setProperty("--color", "#ffc600");

/**
  * @param {string} varName
  * @param {string} unit
  * @returns {function(Event)}
  */
function modify(varName, unit="") {
  return function(evt) {
    if (!evt.target || !(evt.target instanceof HTMLInputElement)) return;
    root?.style.setProperty(varName, evt.target.value + unit);
  }
}

document.querySelectorAll("input").forEach(ele => {
  const varName = "--" + ele.getAttribute("id");
  const unit = ele.dataset.unit;

  ele.addEventListener("change", modify(varName, unit));
  if (varName !== "color") {
    ele.addEventListener("mousemove", modify(varName, unit));
  }
});
