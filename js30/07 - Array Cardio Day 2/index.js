// @ts-check
"use strict";

const people = [
  { name: 'Wes', year: 1988 },
  { name: 'Kait', year: 1986 },
  { name: 'Irv', year: 1970 },
  { name: 'Lux', year: 2015 }
];

const comments = [
  { text: 'Love this!', id: 523423 },
  { text: 'Super good', id: 823423 },
  { text: 'You are the best', id: 2039842 },
  { text: 'Ramen is my fav food ever', id: 123523 },
  { text: 'Nice Nice Nice!', id: 542328 }
];

const year = (new Date()).getFullYear();

// 1. Is at least one person 19 or older?
console.log({anyAdult: people.some(p => year - p.year >= 19)});

// 2. Is everyone 19 or older?
console.log({allAdult: people.every(p => year - p.year >= 19)});

// 3. Find the comment with the ID of 823423
console.log(comments.find(c => c.id === 823423)?.text);

// 4. Find the comment with this ID
const pos = comments.findIndex(c => c.id === 823423);
console.log({pos}); // clever way to quickly display var: value on console

// 5. delete the comment with the ID of 
console.log({len: comments.length});
pos != -1 && comments.splice(pos, 1);
console.log({len: comments.length});
