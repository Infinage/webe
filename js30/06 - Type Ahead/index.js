// @ts-check
"use strict";

// Await doesn't work unless the js file is a 
// module, hence adding export.
export {};
 
/** Denotes a city
  * @typedef {Object} City
  * @property {string} city
  * @property {string} growth_from_2000_to_2013 
  * @property {number} latitude
  * @property {number} longitude
  * @property {string} population
  * @property {string} rank
  * @property {string} state
*/

/** @type City[] */
const cities = Array.from(await (await fetch('cities.json')).json());

const resultList = document.getElementById("search-result");
if (!resultList) throw "Search result not initialized";

const cityTemplate = document.getElementById("city-template");
if (!cityTemplate) throw "City template not initialized"; 

/**
  * Helper to return a list with given city.
  * @param {City} c
  * @returns {HTMLElement}
  */
function li(c) {
  const template = /** @type {HTMLElement} */ (cityTemplate).cloneNode();
  if (!template || !(template instanceof HTMLElement)) throw "Clone fail";
  template.innerHTML = c.city;
  return template;
}

document.getElementById("search-input")?.addEventListener("keyup", evt => {
  if (!evt.target || !(evt.target instanceof HTMLInputElement)) return;
  const search = evt.target.value.toLowerCase();
  resultList.innerHTML = "";

  if (!search) {
    const li1 = document.createElement("li");
    const li2 = document.createElement("li");
    li1.innerText = "Fiter For A City";
    li2.innerText = "Or A State";
    resultList.append(li1, li2);
    return;
  }

  const filtered = cities.filter(c => {
    return c.city.toLowerCase().includes(search) || 
      c.state.toLowerCase().includes(search);
  });
  resultList.append(...filtered.map(li));
});
