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
const cityTemplate = document.getElementById("city-template");

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

/**
  * Returns a function to filter cities with specified debounce
  * @param {number} delay
  * @returns {(evt: Event) => void}
*/
function filterCitiesFn(delay) {
  /** @type {Number} */ let timeout; 

  return (evt) => {
    clearTimeout(timeout);
    timeout = setTimeout(() => {
      if (!evt.target || !(evt.target instanceof HTMLInputElement) || !resultList) 
        return;

      const search = evt.target.value.toLowerCase();

      // Clear current list
      resultList.innerHTML = "";

      // Default if no input
      if (!search) {
        const li1 = document.createElement("li");
        const li2 = document.createElement("li");
        li1.innerText = "Fiter For A City";
        li2.innerText = "Or A State";
        resultList.append(li1, li2);
        return;
      }

      // Filtered list if something is present
      const filtered = cities.filter(c => {
        return c.city.toLowerCase().includes(search) || 
          c.state.toLowerCase().includes(search);
      });
      resultList.append(...filtered.map((c) => li(c, search)));
    }, delay);
  }
};

document.getElementById("search-input")?.addEventListener("keyup", filterCitiesFn(200));
