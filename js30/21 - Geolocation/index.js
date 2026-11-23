// @ts-check
"use strict";

const arrow = document.querySelector(".arrow");
if (!arrow || !(arrow instanceof SVGElement)) throw "Compass arrow not found";

const speedValue = document.querySelector(".speed > .speed-value");
if (!speedValue || !(speedValue instanceof HTMLElement)) throw "Speed value span not found";

navigator.geolocation.watchPosition(data => {
  arrow.style.rotate = `${data.coords.heading}deg`;
  speedValue.innerText = String(data.coords.speed || 0);
  console.info(data.coords.toJSON());
}, err => console.log(err));
