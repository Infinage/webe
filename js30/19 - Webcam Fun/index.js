// @ts-check
"use strict";

async function startCamera() {
  const device = await navigator.mediaDevices.getUserMedia({video: true, audio: false});
  const video = document.querySelector("video");
  if (!video) throw "Expected video element";

  video.srcObject = device;
  await video.play();
}

async function stopCamera() {
  const video = document.querySelector("video");
  if (!video) throw "Expected video element";
  video.pause();

  const stream = video.srcObject;
  if (!stream || !(stream instanceof MediaStream)) return;
  stream.getTracks().forEach(track => track.stop());
  video.srcObject = null;
}

async function takePhoto() {
  const canvas = document.querySelector("canvas");
  if (!canvas) throw "Canvas not found";

  const blob = canvas.toDataURL("data/webp");
  const downloadLink = document.createElement("a");
  downloadLink.href = blob;
  downloadLink.setAttribute("download", "handsome");
  downloadLink.innerHTML = `<img src="${blob}" alt="Handsome Man" />`
  document.querySelector(".strip")?.appendChild(downloadLink);
}

/**
 * Helper to create a visual effect of splitting the 
 * different channel colors to either side.
 *
 * @param {ImageData} img 
 * @returns {ImageData}
 */
function rgbSplitEffect(img) {
  for (let i = 0; i < img.data.length; i+=4) {
    img.data[i - 150] = img.data[i + 0]; // RED
    img.data[i + 100] = img.data[i + 1]; // GREEN
    img.data[i - 150] = img.data[i + 2]; // BLUE
  }
  return img;
}

window.addEventListener("load", async () => {
  const video = document.querySelector("video");
  const canvas = document.querySelector("canvas");
  if (!video || !canvas) throw "Video / Canvas element(s) not found";

  await startCamera();
  [canvas.width, canvas.height] = [video.videoWidth, video.videoHeight];

  const ctx = canvas.getContext("2d");
  if (!ctx) throw "Canvas not supported";

  setInterval(() => {
    ctx.drawImage(video, 0, 0);
    let img = ctx.getImageData(0, 0, canvas.width, canvas.height);
    ctx.globalAlpha = 0.1;
    ctx.putImageData(rgbSplitEffect(img), 0, 0);
  }, 15);

  const btn = document.querySelector("main form button");
  if (!btn || !(btn instanceof HTMLButtonElement)) 
    throw "Expected to find button";
  btn.disabled = false;
  btn.addEventListener("click", takePhoto);
});
