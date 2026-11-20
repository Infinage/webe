// @ts-check
"use strict"

const canvas = document.querySelector("canvas");
if (!canvas) throw "Canvas not initalized";

const canvasCtx = canvas.getContext("2d");
if (!canvasCtx) throw "Canvas not supported";

// Constants for the canvas api
let [ox, oy] = [0, 0];
let clicked = false;
let direction = true;
let hue = 0;

window.addEventListener("load", initCanvas);
window.addEventListener("resize", initCanvas);

window.addEventListener("mouseup", () => clicked = false);
window.addEventListener("mouseleave", () => clicked = false);

window.addEventListener("mousedown", (evt) => {
  clicked = true;
  [ox, oy] = [evt.x, evt.y];
  dot(evt.x, evt.y);
});

window.addEventListener("mousemove", evt => {
  if (!clicked) return;
  dot(evt.x, evt.y);
});

/**
  * Helper to set the canvas height, width and draw styles.
  */
function initCanvas() {
  if (!canvas || !canvasCtx) return;
  canvas.height = window.innerHeight;
  canvas.width = window.innerWidth;
  canvasCtx.lineJoin = 'round';
  canvasCtx.lineCap = 'round';
}

/**
  * Draw on the canvas given input coordinates
  * @param {number} x
  * @param {number} y
  */
function dot(x, y) {
  if (!canvasCtx) return;
  canvasCtx.strokeStyle = `hsl(${hue}, 100%, 50%)`;
  canvasCtx.beginPath();
  canvasCtx.moveTo(ox, oy);
  canvasCtx.lineTo(x, y);
  canvasCtx.stroke(); 
  [ox, oy] = [x, y];

  hue++;
  if (hue >= 360) hue = 0;

  if (direction) {
    canvasCtx.lineWidth++;
  } else {
    canvasCtx.lineWidth--;
  }

  if (canvasCtx.lineWidth >= 100 || canvasCtx.lineWidth <= 1)
    direction = !direction;
}
