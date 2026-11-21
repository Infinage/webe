// @ts-check
"use strict";

let volInpClicked = false;
let playbackInpClicked = false;
let progressBarClicked = false;

window.addEventListener("load", () => {
  const video = document.querySelector("video");
  if (!video) throw "Video tag not initialized";

  // Two way play / pause controls and state
  const playPause = document.querySelector("button[name='play-pause']");
  if (!playPause) throw "Play button not initialized";
  video.addEventListener("play", () => playPause.textContent = "❚❚");
  video.addEventListener("pause", () => playPause.textContent = "▶");
  playPause.addEventListener("click", () => video.paused? video.play(): video.pause());

  // Volume controls
  const volume = document.querySelector("input[name='volume']");
  if (!volume || !(volume instanceof HTMLInputElement))
    throw "Volume control not initialized";
  volume.addEventListener("change", () => video.volume = Number(volume.value));
  volume.addEventListener("mousedown", () => volInpClicked = true);
  volume.addEventListener("mouseup", () => volInpClicked = false);
  volume.addEventListener("mouseleave", () => volInpClicked = false);
  volume.addEventListener("mousemove", () => {
    if (!volInpClicked) return;
    video.volume = Number(volume.value);
  });

  // Playback controls
  const playback = document.querySelector("input[name='playback']");
  if (!playback || !(playback instanceof HTMLInputElement))
    throw "Playback control not initialized";
  playback.addEventListener("change", () => video.playbackRate = Number(playback.value));
  playback.addEventListener("mousedown", () => playbackInpClicked = true);
  playback.addEventListener("mouseup", () => playbackInpClicked = false);
  playback.addEventListener("mouseleave", () => playbackInpClicked = false);
  playback.addEventListener("mousemove", () => {
    if (!playbackInpClicked) return;
    video.playbackRate = Number(playback.value);
  });

  // Seek forward / backward controls
  const seekForward = document.querySelector("button[name='seek-forward']");
  const seekBackward = document.querySelector("button[name='seek-backward']");
  if (!seekForward || !seekBackward) throw "Seekn forward / backward buttons not found";
  seekForward.addEventListener("click", () => video.currentTime += 25);
  seekBackward.addEventListener("click", () => video.currentTime -= 10);

  // Update progres bar automatically
  const barPosition = document.querySelector(".bar-position");
  if (!barPosition || !(barPosition instanceof HTMLElement)) 
    throw "Progress bar not initalised";
  video.addEventListener("timeupdate", () => {
    barPosition.style.width = `${(video.currentTime / video.duration) * 100}%`;
  })

  // Click, drag on progress bar
  const progressBar = document.querySelector(".progress-bar");
  if (!progressBar) throw "Progress bar not initialized";
  progressBar.addEventListener("mouseup", () => progressBarClicked = false);
  progressBar.addEventListener("mouseleave", () => progressBarClicked = false);
  progressBar.addEventListener("mousedown", evt => {
    if (!(evt instanceof MouseEvent)) return;      
    progressBarClicked = true;
    const pct = evt.offsetX / progressBar.clientWidth;
    video.currentTime = video.duration * pct;
  });
  progressBar.addEventListener("mousemove", evt => {
    if (!progressBarClicked || !(evt instanceof MouseEvent)) return;
    const pct = evt.offsetX / progressBar.clientWidth;
    video.currentTime = video.duration * pct;
  });

  // Double click to go full screen
  const customPlayer = document.querySelector(".custom-player")
  if (!customPlayer || !(customPlayer instanceof HTMLElement)) 
    throw "Custom player not initialized";
  video.addEventListener("click", () => video.paused? video.play(): video.pause());
  video.addEventListener("dblclick", () => {
    if (document.fullscreenElement) {
      document.exitFullscreen();
      return;
    }
    customPlayer.requestFullscreen();
  });
})

