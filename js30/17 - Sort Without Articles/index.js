// @ts-check
"use strict";

const bands = [
  'The Plot in You', 'The Devil Wears Prada', 'Pierce the Veil', 'Norma Jean', 
  'The Bled', 'Say Anything', 'The Midway State', 'We Came as Romans', 'Counterparts', 
  'Oh, Sleeper', 'A Skylit Drive', 'Anywhere But Here', 'An Old Dog'
];

/** Utility to trim 'A', 'An', 'The' from string
  * @param {string} str
*/
function trimArticles(str) {
  const searchStr = new RegExp(/^(the|an|a)\b\s/, "i");
  return str.replace(searchStr, "");
}

const sortedBands = bands.sort((b1, b2) => {
  return trimArticles(b1).localeCompare(trimArticles(b2));
});


const ul = document.createElement("ul");
ul.innerHTML = sortedBands.map(band => `<li>${band}</li>`).join('');
document.body.appendChild(ul);
