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
  * Utility function that wraps highlighted sections with span.highlight
  * @param {string} text
  * @param {string} search
  * @returns {HTMLElement}
*/
function highlight(text, search) {
  const span = document.createElement("span");
  const searchCleaned = search.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  text = text.replaceAll(new RegExp(`(${searchCleaned})`, "ig"), 
    `<span class="highlight">$1</span>`);
  span.innerHTML = text;
  return span;
}

/**
  * Helper to return a list with given city.
  * @param {City} c
  * @param {string} search
  * @returns {DocumentFragment}
  */
function li(c, search) {
  const templateContent = /** @type {HTMLTemplateElement} */ (cityTemplate).content
  const template = document.importNode(templateContent, true);

  const cityName = template.querySelector(".city-name")
  const cityState = template.querySelector(".city-state")
  const cityPopulation = template.querySelector(".city-population")

  if (!cityName || !cityState || !cityPopulation) 
    throw "Missing required template tags";

  cityName.appendChild(highlight(c.city, search));
  cityState.appendChild(highlight(c.state, search));
  cityPopulation.textContent = Number(c.population).toLocaleString();

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
  resultList.append(...filtered.map((c) => li(c, search)));
});
