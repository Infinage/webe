//@ts-check
"use strict";

/**
  * Helper to toggle checkbox and apply text strikethrough.
  * @param {HTMLElement} ele
  */
function toggleElement(ele) {
  const textItem = ele.querySelector("span:not(:has(input))");
  const checked = textItem?.classList.toggle("strikethrough");

  const inp = ele.querySelector("input");
  if (inp) inp.checked = checked || false;

  ele.classList.toggle("selected");
}

document.querySelectorAll("li").forEach(ele => {
  ele.addEventListener("click", evt => {
    // Click events fire after drag, if dragged do not click
    const selection = window.getSelection();
    const highlightedTextNode = selection?.toString()? selection?.anchorNode: undefined;
    const highlightedLi = highlightedTextNode?.parentElement?.parentElement;
    if (highlightedLi === ele) return;

    if (evt.shiftKey && !ele.classList.contains("selected")) {
      let above = /** @type {Element | null} */ (ele);
      for (; above && !above.classList.contains("selected"); above = above.previousElementSibling);
      if (above) {
        above = above.nextElementSibling;
        for (; above && above !== ele; above = above.nextElementSibling) {
          if (above instanceof HTMLElement) toggleElement(above);
        }
      }
    }

    toggleElement(ele);
  })
});
