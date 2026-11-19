//@ts-check
"use strict";

const inventors = [
  { first: 'Albert', last: 'Einstein', year: 1879, passed: 1955 },
  { first: 'Isaac', last: 'Newton', year: 1643, passed: 1727 },
  { first: 'Galileo', last: 'Galilei', year: 1564, passed: 1642 },
  { first: 'Marie', last: 'Curie', year: 1867, passed: 1934 },
  { first: 'Johannes', last: 'Kepler', year: 1571, passed: 1630 },
  { first: 'Nicolaus', last: 'Copernicus', year: 1473, passed: 1543 },
  { first: 'Max', last: 'Planck', year: 1858, passed: 1947 },
  { first: 'Katherine', last: 'Blodgett', year: 1898, passed: 1979 },
  { first: 'Ada', last: 'Lovelace', year: 1815, passed: 1852 },
  { first: 'Sarah E.', last: 'Goode', year: 1855, passed: 1905 },
  { first: 'Lise', last: 'Meitner', year: 1878, passed: 1968 },
  { first: 'Hanna', last: 'Hammarström', year: 1829, passed: 1909 }
]

const people = [
  'Bernhard, Sandra', 'Bethea, Erin', 'Becker, Carl', 'Bentsen, Lloyd', 'Beckett, Samuel', 'Blake, William', 'Berger, Ric', 'Beddoes, Mick', 'Beethoven, Ludwig',
  'Belloc, Hilaire', 'Begin, Menachem', 'Bellow, Saul', 'Benchley, Robert', 'Blair, Robert', 'Benenson, Peter', 'Benjamin, Walter', 'Berlin, Irving',
  'Benn, Tony', 'Benson, Leana', 'Bent, Silas', 'Berle, Milton', 'Berry, Halle', 'Biko, Steve', 'Beck, Glenn', 'Bergman, Ingmar', 'Black, Elk', 'Berio, Luciano',
  'Berne, Eric', 'Berra, Yogi', 'Berry, Wendell', 'Bevan, Aneurin', 'Ben-Gurion, David', 'Bevel, Ken', 'Biden, Joseph', 'Bennington, Chester', 'Bierce, Ambrose',
  'Billings, Josh', 'Birrell, Augustine', 'Blair, Tony', 'Beecher, Henry', 'Biondo, Frank'
]

// 1. Filter the list of inventors for those who were born in the 1500's (16 century)
//console.log(inventors.filter(person => Math.floor(person.year / 100) === 15.));
console.table(inventors.filter(person => String(person.year).startsWith("15")));

// 2. Array of inventors first and last names
console.log(inventors.map(({first, last}) => { return `${first} ${last}`; }));

// 3. Sort the inventors by birthdate (oldest to youngest)
console.table(inventors.sort((p1, p2) => p1.year - p2.year));

/**
  * A person type
  * @typedef {Object} Person
  * @property {string} first
  * @property {string} last
  * @property {number} year
  * @property {number} passed
*/

/**
  * Reusable helper to compute the years lived
  * @param {Person} p
*/
const yearsLived = p => p.passed - p.year;

// 4. How many years did all the invetors live
console.log(inventors.reduce((total, p) => total + yearsLived(p), 0));

// 5. Sort the inventors by years lived
console.table(inventors.sort((p1, p2) => yearsLived(p1) - yearsLived(p2)));

// 6. create a list of Boulevards in Paris that contain 'de' anywhere in the name
//document.querySelectorAll(".mw-category-group li > a").forEach(ele => { 
//    const name = /** @type HTMLElement */ (ele).innerText;
//  if (name.includes("de")) console.log(name); 
//})

// 7. Sort the people alphabetically by last name
console.log(people.sort((p1, p2) => {
  const [lname1, lname2] = [p1.split(", ")[1], p2.split(", ")[1]];
  return lname1.localeCompare(lname2);
}));


// 8. Sum up the instances of each of these
const data = ['car', 'car', 'truck', 'truck', 'bike', 'walk', 
  'car', 'van', 'bike', 'walk', 'car', 'van', 'car', 'truck' ];
const dataCounter = data.reduce((counter, vehicle) => { 
  const count = counter.get(vehicle);
  counter.set(vehicle, count? count + 1: 1);
  return counter; 
}, new Map());
console.table(dataCounter);
