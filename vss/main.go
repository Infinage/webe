// VSS Associates website.
//
// Server-rendered Go templates styled with Tailwind, interactive via
// Datastar. Nav links issue Datastar @get requests; the server then
// responds with an SSE patch replacing <main> (plus a pushState so the
// URL stays in sync). Without the Datastar request header the same
// routes render the full page.
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/starfederation/datastar-go/datastar"
)

// Spec is one label/value row in a project spec ledger.
type Spec struct {
	K, V string
}

// Project is one card in the home showcase.
type Project struct {
	Name, Loc, Status, Image, Href string
	Lead                           bool
	Specs                          []Spec
}

// Stat is one number in the home stat band.
type Stat struct {
	Target int
	Unit   string
	Label  string
	Req    bool
}

// Prop is one numbered value proposition.
type Prop struct {
	Num         string
	Title, Href string
	Body        string
}

// Marker is one pin on the home geographic band.
type Marker struct {
	Left, Top string
	Name, Sub string
	Href      string
}

// Step is one buying-process step.
type Step struct {
	Num         string
	Title, Body string
}

// FAQ is one question in the home accordion. Answers are trusted,
// hand-written HTML (links, emphasis).
type FAQ struct {
	Idx, Q string
	A      template.HTML
	Open   bool
	Req    bool // starred — client confirmation required
}

// Page carries everything the layout and page templates need.
type Page struct {
	Title       string
	HeaderSolid bool // solid header (pages without a dark hero)
	Data        any
}

var pages = map[string]string{
	"home":        "home.html",
	"placeholder": "placeholder.html",
}

var tmpl map[string]*template.Template

// homeData is sample content mirroring the approved 01-home artifact.
// Real data modelling comes in a later iteration.
func homeData() map[string]any {
	return map[string]any{
		"Stats": []Stat{
			{Target: 10, Unit: "yrs", Label: "In business", Req: true},
			{Target: 600, Unit: "+", Label: "Families served", Req: true},
			{Target: 120, Unit: "+", Label: "Acres developed", Req: true},
			{Target: 4, Label: "Layouts delivered"},
			{Target: 100, Unit: "%", Label: "Clear titles", Req: true},
		},
		"Projects": []Project{
			{
				Name: "Landmark City", Loc: "Chennai North*",
				Status: "Viewing by appointment", Image: "/assets/img/card-landmark.jpg",
				Href: "/projects/landmark-city", Lead: true,
				Specs: []Spec{
					{K: "Area", V: "~9.8 acres*"},
					{K: "Plots", V: "24*"},
					{K: "Plot sizes", V: "600–2,400 sq.ft*"},
					{K: "Price", V: "₹1,850 / sq.ft*"},
				},
			},
			{
				Name: "Srivelan Nagar", Loc: "Chennai North*",
				Status: "Ready for visits", Image: "/assets/img/card-srivelan.jpg",
				Href: "/projects/srivelan-nagar",
				Specs: []Spec{
					{K: "Area", V: "~7.4 acres*"},
					{K: "Plots", V: "36*"},
					{K: "Plot sizes", V: "800–2,600 sq.ft*"},
					{K: "Price", V: "on request*"},
				},
			},
			{
				Name: "Chandrayan Nagar", Loc: "Thiruppalai Nadu*",
				Status: "Ready for visits", Image: "/assets/img/card-chandrayan.jpg",
				Href: "/projects/chandrayan-nagar",
				Specs: []Spec{
					{K: "Area", V: "~6.2 acres*"},
					{K: "Plots", V: "30*"},
					{K: "Plot sizes", V: "600–1,800 sq.ft*"},
					{K: "Price", V: "on request*"},
				},
			},
		},
		"Props": []Prop{
			{Num: "01", Title: "Approved before it's offered",
				Body: "Every layout is released only after DTCP approval — and the approval number for each project is published on its own page, so you can verify it yourself."},
			{Num: "02", Title: "Choose your plot before you visit",
				Href: "/projects/landmark-city",
				Body: "Every layout carries a live plot map — numbered plots, areas and status.\nShortlist your number on the site, then walk that shortlist in person."},
			{Num: "03", Title: "Inspect the ground, not the brochure",
				Body: "Guided site visits on working days. Walk the plot you're buying — stand on it, check the road frontage, see the neighbours — before you pay a rupee."},
			{Num: "04", Title: "Clear title to patta",
				Body: "Transparent paperwork at every step: agreement, registration and patta transfer, supported by the same Adambakkam office that sold you the plot."},
			{Num: "05", Title: "Built to hold value",
				Body: "Fenced boundaries, laid internal roads and layouts that are maintained after handover — the difference between land you own and land you'd want to own."},
		},
		"Markers": []Marker{
			{Left: "18%", Top: "34%", Name: "Landmark City",
				Sub: "Chennai North* · 3 plots remaining*", Href: "/projects/landmark-city"},
			{Left: "38%", Top: "62%", Name: "Srivelan Nagar",
				Sub: "Chennai North* · Now selling*", Href: "/projects/srivelan-nagar"},
			{Left: "63%", Top: "30%", Name: "Chandrayan Nagar",
				Sub: "Thiruppalai Nadu* · Limited plots*", Href: "/projects/chandrayan-nagar"},
			{Left: "82%", Top: "56%", Name: "OMR project",
				Sub: "OMR, Chennai · details being prepared", Href: "/projects"},
		},
		"Steps": []Step{
			{Num: "01", Title: "Site visit",
				Body: "Walk the layout on a working day — Mon–Sat, 8.00–18.00, by appointment."},
			{Num: "02", Title: "Choose & reserve",
				Body: "Pick your plot number on the live map and place it on hold while you decide."},
			{Num: "03", Title: "Agreement & verification",
				Body: "Title documents, sale agreement and bank-loan paperwork handled transparently."},
			{Num: "04", Title: "Registration & patta",
				Body: "We stand beside you at registration and follow through to patta in your name."},
		},
		"FAQs": []FAQ{
			{Idx: "01", Open: true, Req: true,
				Q: "What does DTCP approval mean for me as a buyer?",
				A: template.HTML(`Directorate of Town &amp; Country Planning approval confirms the layout's plan — roads, plot boundaries and common areas — meets state norms before a single plot is sold. We publish the approval number for each project on its page. <em>(Answer to be confirmed by client.)</em>`)},
			{Idx: "02", Req: true,
				Q: "Can nationalised banks fund these plots?",
				A: template.HTML(`DTCP-approved layouts with clear titles are generally eligible for plot loans from major banks. We'll share the exact lenders currently working with each project. <em>(Answer to be confirmed by client.)</em>`)},
			{Idx: "03",
				Q: "Will I get patta in my name?",
				A: template.HTML(`Yes — supporting the transfer of patta to your name after registration is part of the standard process we handle for every buyer.`)},
			{Idx: "04",
				Q: "Can I visit the site before booking?",
				A: template.HTML(`Always, and we insist on it. Site visits run Monday to Saturday, 8.00–18.00. Call <a href="tel:04442087913" style="text-decoration:underline">044 4208 7913</a> to book a guided walk of the actual plots.`)},
		},
	}
}

// loadTemplates parses the shared layout/partials once, then clones it
// per page so every page can define its own "main" template.
func loadTemplates() error {
	tmpl = make(map[string]*template.Template)
	base, err := template.ParseFS(os.DirFS("assets"),
		"html/layout.html", "html/partials.html")
	if err != nil {
		return fmt.Errorf("parse shared templates: %w", err)
	}
	for name, file := range pages {
		t, err := base.Clone()
		if err != nil {
			return fmt.Errorf("clone base for %s: %w", name, err)
		}
		if _, err := t.ParseFS(os.DirFS("assets"), "html/"+file); err != nil {
			return fmt.Errorf("parse %s: %w", file, err)
		}
		tmpl[name] = t
	}
	return nil
}

// isFragment reports whether the request came from a Datastar action
// (and therefore expects an SSE patch instead of a full page).
func isFragment(r *http.Request) bool {
	return r.Header.Get("Datastar-Request") == "true"
}

// renderFull writes the complete page for non-Datastar requests.
func renderFull(w http.ResponseWriter, name string, p Page) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl[name].ExecuteTemplate(w, "layout", p); err != nil {
		log.Printf("render %s: %v", name, err)
		http.Error(w, "template error", http.StatusInternalServerError)
	}
}

// renderFragment patches <main> over SSE and keeps the URL in sync.
func renderFragment(w http.ResponseWriter, r *http.Request,
	name string, p Page,
) {
	var content bytes.Buffer
	if err := tmpl[name].ExecuteTemplate(&content, "main", p); err != nil {
		log.Printf("render fragment %s: %v", name, err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
	elem := "<main id=\"main\">" + content.String() + "</main>"
	sse := datastar.NewSSE(w, r)
	if err := sse.PatchElements(elem); err != nil {
		log.Printf("patch elements: %v", err)
		return
	}
	if err := sse.ExecuteScript(
		fmt.Sprintf("history.pushState(null, '', %q); document.title = %q",
			r.URL.Path, p.Title),
	); err != nil {
		log.Printf("pushState script: %v", err)
	}
}

// render serves a page both ways: full HTML normally, an SSE patch of
// <main> for Datastar fragment navigation.
func render(w http.ResponseWriter, r *http.Request, name string, p Page) {
	if isFragment(r) {
		renderFragment(w, r, name, p)
		return
	}
	renderFull(w, name, p)
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, r, "home", Page{
		Title: "VSS Associates — DTCP-approved plots in Chennai",
		Data:  homeData(),
	})
}

// placeholder serves stub pages for screens that arrive in later
// iterations; admin pages stay static placeholders for now.
func placeholder(title string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "placeholder", Page{
			Title:       title + " — VSS Associates",
			HeaderSolid: true,
			Data:        map[string]any{"Title": title},
		})
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	if isFragment(r) {
		renderFragment(w, r, "placeholder", Page{
			Title: "Page not found — VSS Associates", HeaderSolid: true,
			Data: map[string]any{"Title": "Page not found"},
		})
		return
	}
	renderFull(w, "placeholder", Page{
		Title: "Page not found — VSS Associates", HeaderSolid: true,
		Data: map[string]any{"Title": "Page not found"},
	})
}

func main() {
	if err := loadTemplates(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.HandleFunc("GET /{$}", home)
	mux.Handle("GET /projects", placeholder("Projects"))
	mux.Handle("GET /projects/{slug}", placeholder("Project"))
	mux.Handle("GET /about", placeholder("About"))
	mux.Handle("GET /contact", placeholder("Contact"))
	mux.Handle("GET /admin/login", placeholder("Staff login"))
	mux.HandleFunc("/", notFound)

	addr := "127.0.0.1:" + port()
	log.Printf("listening on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

// port returns the listen port (PORT env, default 8080).
func port() string {
	p := os.Getenv("PORT")
	if p == "" {
		return "8080"
	}
	return strings.TrimPrefix(p, ":")
}
