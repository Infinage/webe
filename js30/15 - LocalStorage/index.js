// @ts-check
"use strict";

const newItemTemplate = document.getElementById("new-item");
if (!newItemTemplate || !(newItemTemplate instanceof HTMLTemplateElement)) 
  throw "Template tag not found";

const addItemBtn = document.querySelector(".add-items button");
if (!addItemBtn) throw "Add item button missing";

const entries = document.querySelector(".entries");
if (!entries) throw "Entries list not found";

window.addEventListener("load", () => {
  entries.innerHTML = localStorage.getItem("local-tapas") || ""; 
});

addItemBtn.addEventListener("click", () => {
  const newItem = newItemTemplate.content.cloneNode(true);
  if (!(newItem instanceof DocumentFragment)) 
    throw "Expected document fragment";

  const itemNameField = document.querySelector(".add-items input");
  if (!itemNameField || !(itemNameField instanceof HTMLInputElement))
    throw "Expected input element here";

  // Ignore empty inputs
  if (!itemNameField.value) return;

  const itemCount = document.querySelectorAll(".entries li").length;
  newItem.querySelector("input")?.setAttribute("id", `item${itemCount}`);
  newItem.querySelector("label")?.setAttribute("for", `item${itemCount}`);

  const label = newItem.querySelector("label");
  if (!label) throw "Label not found in template";
  label.innerText = itemNameField.value;
  itemNameField.value = "";
  entries.appendChild(newItem);

  // Persist to storage, not as performant but should be okay
  localStorage.setItem("local-tapas", entries?.innerHTML || ""); 
});

// Modify the input checkbox with 'checked' attribute before storing
entries.addEventListener("click", evt => {
  if (!(evt.target instanceof HTMLInputElement)) return;
  if (evt.target.checked) { evt.target.setAttribute("checked", ""); } 
  else { evt.target.removeAttribute("checked"); }
  localStorage.setItem("local-tapas", entries?.innerHTML || ""); 
});
