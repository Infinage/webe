// @ts-check
"use strict";

const recognition = new window.SpeechRecognition();
recognition.interimResults = true;

let stopped = false;

const page = document.querySelector(".page");
if (!page) throw "Page not found";

recognition.addEventListener("result", 
  /** @param {SpeechRecognitionEvent} evt */ 
  evt => {
  const transcript = Array.from(evt.results)
    .map(result => result[0].transcript)
    .join('');

  let p = document.querySelector("div > .para:last-child");
  if (!p) throw "Missing para tag";
  p.textContent = transcript;

  // Add a 'para' break
  if (evt.results[0].isFinal) {
    if (transcript === "stop") {stopped = true; return};
    page.insertAdjacentHTML("beforeend", "<div class='para' contenteditable></div>");
  }
});

recognition.addEventListener("end", () => !stopped && recognition.start());

recognition.start();
