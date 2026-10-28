import styles from "/assets/live-search.css" with { type: 'css' };

const response = await fetch("/assets/live-search.html");
const template = document.createElement("template");
template.innerHTML = await response.text();

class LiveSearch extends HTMLElement {
  constructor() {
    super();

    this.attachShadow({mode: 'open'});
    this.shadowRoot.appendChild(template.content.cloneNode(true));
    this.shadowRoot.adoptedStyleSheets = [styles];

    this.timeout = null;
    this.search = this.shadowRoot.querySelector("input[type='text']");
    this.resultList = this.shadowRoot.querySelector("ul");
    this.indicator = this.shadowRoot.querySelector("span");

    this.state = new Proxy ({
      query: "",
      results: [],
      isSearching: false,
    }, {
      set: (target, property, value) => {
        target[property] = value;  

        if (property === "isSearching") {
          this.indicator.hidden = !value;
          if (value) {
            this.resultList.hidden = true;
          }
        }

        if (property === 'results') {
          this.resultList.innerHTML = '';
          value.forEach(l => {
            const li = document.createElement("li");
            li.innerText = l;
            this.resultList.appendChild(li);
          });
          if (this.resultList.children.length > 0) {
            this.resultList.hidden = false;
          }
        }

        return true;
      }
    });
  }

  connectedCallback() {
    this.search.oninput = (event) => {
      clearTimeout(this.timeout);
      this.state.query = event.target.value;
      const paramString = new URLSearchParams({q: this.state.query}).toString();

      if (this.state.query === '') {
        this.state.results = [];
        return;
      }

      this.timeout = setTimeout(async () => {
        this.state.isSearching = true;
        try {
          const response = await fetch(`/api/search?${paramString}`);
          this.state.results = await response.json() || [];
        } catch (err) {
          console.error("Fetch failed: ", err);
        } finally {
          this.state.isSearching = false;
        }
      }, 500);
    }
  }

}

customElements.define("x-live-search", LiveSearch)
