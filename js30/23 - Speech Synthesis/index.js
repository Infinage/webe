// @ts-check
"use strict";

/** @type SpeechSynthesisVoice[] */
let voices = [];

const voicesDropdown = document.querySelector("select[name='voice']");
if (!(voicesDropdown instanceof HTMLSelectElement)) throw "Select not initialized";

const speechRate = document.getElementById("speech-rate");
if (!(speechRate instanceof HTMLInputElement)) throw "Speech rate input not found";

const speechPitch = document.getElementById("speech-pitch");
if (!(speechPitch instanceof HTMLInputElement)) throw "Speech pitch input not found";

const textArea = document.querySelector("textarea");
if (!(textArea instanceof HTMLTextAreaElement)) throw "Speech input not found";

const stopBtn = document.querySelector("button[name='stop']")
if (!(stopBtn instanceof HTMLButtonElement)) throw "Stop button not found";

const speakBtn = document.querySelector("button[name='speak']")
if (!(speakBtn instanceof HTMLButtonElement)) throw "Speak button not found";

const msg = new SpeechSynthesisUtterance();
msg.text = textArea.value;

speechSynthesis.addEventListener("voiceschanged", () => {
  voices = speechSynthesis.getVoices().filter(voice => voice.lang.includes("en"));
  voices.forEach(voice => {
      const opt = `<option value="${voice.name}">${voice.name} (${voice.lang})</option>`;
      voicesDropdown.insertAdjacentHTML("beforeend", opt);
  });
});

voicesDropdown.addEventListener("change", evt => {
  if (!msg) throw "SpeechSynthesisUtterance is null";
  const voice = /** @type {HTMLSelectElement} */ (evt.target).value;
  msg.voice = voices.find(voiceObj => voiceObj.name == voice) || null;
  console.log(`Voice set to ${msg.voice?.name}`);
  if (speechSynthesis.speaking) {
    speechSynthesis.cancel();
    speechSynthesis.speak(msg);
  }
});

[speechRate, speechPitch, textArea].forEach(ele => ele.addEventListener("change", evt => {
  const value = /** @type {HTMLInputElement} */ (evt.target).value;

  // @ts-ignore 
  msg[ele.name] = value;
  console.log(`${ele.name} is now set to ${value}`)
  if (speechSynthesis.speaking) {
    speechSynthesis.cancel();
    speechSynthesis.speak(msg);
  }
}));

stopBtn.addEventListener("click", () => speechSynthesis.cancel());
speakBtn.addEventListener("click", () => {
  speechSynthesis.cancel();
  speechSynthesis.speak(msg);
});
