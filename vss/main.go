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
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
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

// PSpec is one fact cell in a projects-index row.
type PSpec struct {
	K, V string
	Req  bool // starred — client confirmation required
}

// PRow is one editorial project row on the projects index.
type PRow struct {
	Num                string // two-digit ghost numeral
	Loc                string
	LocReq             bool
	Name               string
	DTCP               bool
	Status             string // second badge text, without its trailing star
	StatusReq          bool   // star as <em class=req> (client confirmation)
	StatusStar         bool   // star as plain text
	Quiet              bool   // row being prepared: no link, no availability strip
	Muted              string
	Specs              []PSpec
	Avail              int // availability strip counts; all zero = no strip
	Resv, Sold         int
	FrontImg, FrontAlt string
	L2Img              string
	Href               string // "" = decorative deck (quiet row)
	First, Even        bool   // computed row position (overlap, deck fan)
}

// Total is the full plot count behind the availability strip.
func (p PRow) Total() int { return p.Avail + p.Resv + p.Sold }

// Dots renders the availability strip in artifact order: available
// ("on"), reserved ("rsv"), sold ("off").
func (p PRow) Dots() []string {
	dots := make([]string, 0, p.Avail+p.Resv+p.Sold)
	for range p.Avail {
		dots = append(dots, "on")
	}
	for range p.Resv {
		dots = append(dots, "rsv")
	}
	for range p.Sold {
		dots = append(dots, "off")
	}
	return dots
}

// Plot is one plot on the interactive site plan. Every fact is
// sample data (asterisked in the UI) pending client supply.
type Plot struct {
	N       string // plot id, e.g. "A1"
	St      string // available | reserved | sold
	Status  string // panel status line, pre-computed
	Dims    string // panel dimensions line, pre-computed
	Area    int    // sq.ft
	AreaFmt string // "1,200" — en-IN grouped, for SVG + panel
	Facing  string
	Corner  bool
	Price   string // panel price line, pre-computed
	Cta     string // panel CTA label, pre-computed
	Note    string // panel fine-print, pre-computed
	X, Y    int    // SVG rect origin (layout geometry, not serialised)
}

// signalsMap is the JSON payload bound into Datastar signals; the
// site-plan template reads $plots.<id> for every panel field.
func (p Plot) signalsMap() map[string]any {
	return map[string]any{
		"st": p.St, "status": p.Status, "dims": p.Dims,
		"area": p.AreaFmt + " sq.ft", "facing": p.Facing,
		"corner": p.Corner, "price": p.Price, "cta": p.Cta, "note": p.Note,
	}
}

// GItem is one gallery figure.
type GItem struct {
	Img, Alt, Num, Title, Cap string
}

// LRow is one landmark distance row.
type LRow struct {
	Name, Sub, Dist string
}

// Legend is one filter pill on the site plan.
type Legend struct {
	Key, Label string
	Swatch     string // "" for All; av | rs | so
}

// projectDetailData is sample content mirroring the approved
// 03-project-detail artifact (Landmark City).
func projectDetailData() map[string]any {
	seeds := []struct {
		N    string
		St   string
		Area int
	}{
		{"A1", "available", 1200}, {"A2", "available", 1200},
		{"A3", "available", 1200}, {"A4", "available", 1200},
		{"A5", "available", 1500}, {"A6", "available", 1500},
		{"A7", "available", 1500}, {"A8", "available", 1500},
		{"A9", "sold", 1800}, {"A10", "reserved", 1800},
		{"A11", "available", 1800}, {"A12", "sold", 1800},
		{"B1", "sold", 2400}, {"B2", "available", 2400},
		{"B3", "reserved", 1500}, {"B4", "sold", 1500},
		{"B5", "available", 1500}, {"B6", "available", 1800},
		{"B7", "available", 1500}, {"B8", "available", 1500},
		{"B9", "available", 1500}, {"B10", "reserved", 1500},
		{"B11", "sold", 1500}, {"B12", "sold", 1500},
	}

	const pricePerSft = 1850
	plots := make([]Plot, 0, len(seeds))
	signals := map[string]any{}
	counts := map[string]int{"all": len(seeds)}
	for i, s := range seeds {
		w := int(math.Round(math.Sqrt(float64(s.Area)/1.5)/5) * 5)
		d := int(math.Round(float64(s.Area)/float64(w)/5) * 5)
		areaFmt := enIN(s.Area)
		p := Plot{
			N: s.N, St: s.St, Area: s.Area, AreaFmt: areaFmt,
			Dims:   fmt.Sprintf("%d × %d ft*", w, d),
			Facing: "East road",
			Corner: s.N == "A1" || s.N == "A12" || s.N == "B1" || s.N == "B12",
			// two rows of 12, top row faces the west road
			X: 100 + (i%12)*73,
			Y: 258,
		}
		if s.N[0] == 'A' {
			p.Facing = "West road"
			p.Y = 62
		}
		switch s.St {
		case "available":
			p.Status = "Available"
			p.Price = fmt.Sprintf("₹%.1f lakh*", float64(s.Area)*pricePerSft/100000)
			p.Cta = "Reserve for a site visit"
			p.Note = "Prices exclude registration & statutory charges.*"
			counts["available"]++
		case "reserved":
			p.Status = "Reserved — not bookable"
			p.Price = "—"
			p.Cta = "Ask about remaining plots"
			p.Note = "Reserved — under advance, not bookable."
			counts["reserved"]++
		default:
			p.Status = "Sold — not bookable"
			p.Price = "—"
			p.Cta = "Ask about remaining plots"
			p.Note = "Sold — shown for orientation only."
			counts["sold"]++
		}
		plots = append(plots, p)
		signals[s.N] = p.signalsMap()
	}
	plotsJSON, err := json.Marshal(signals)
	if err != nil {
		log.Printf("plots signals: %v", err)
	}

	legend := []Legend{
		{Key: "all", Label: fmt.Sprintf("All · %d", counts["all"])},
		{Key: "available", Label: fmt.Sprintf("Available · %d", counts["available"]), Swatch: "av"},
		{Key: "reserved", Label: fmt.Sprintf("Reserved · %d", counts["reserved"]), Swatch: "rs"},
		{Key: "sold", Label: fmt.Sprintf("Sold · %d", counts["sold"]), Swatch: "so"},
	}

	gallery := []GItem{
		{Img: "card-landmark.jpg", Alt: "Entrance signboard of Landmark City", Num: "01",
			Title: "Entrance signboard", Cap: "— approval number displayed at the gate"},
		{Img: "proof-roads.jpg", Alt: "Internal roads laid between plots", Num: "02",
			Title: "Site progress", Cap: "— internal roads & plot kerbs"},
		{Img: "proof-stakes.jpg", Alt: "Plot edges and road kerbs marked out", Num: "03",
			Title: "Plot edges", Cap: "— marked out for site visits"},
		{Img: "proof-visit.jpg", Alt: "Site visit group walking the layout", Num: "04",
			Title: "Site visits", Cap: "— walk the ground before you buy"},
		{Img: "proof-launch.jpg", Alt: "Launch event canopy", Num: "05",
			Title: "Launch day", Cap: "— booking event under canopy"},
		{Img: "proof-build.jpg", Alt: "Boundary construction work", Num: "06",
			Title: "Work in progress", Cap: "— boundary development"},
	}

	locRows := []LRow{
		{Name: "Nearest bus stop", Sub: "Minjur bus stand*", Dist: "~2 km*"},
		{Name: "Schools & colleges", Sub: "Ponneri school cluster*", Dist: "~4 km*"},
		{Name: "Hospital", Sub: "Govt. hospital, Ponneri*", Dist: "~6 km*"},
		{Name: "Railway / metro station", Sub: "Ponneri railway station*", Dist: "~5 km*"},
		{Name: "Market & daily needs", Sub: "Minjur market*", Dist: "~2 km*"},
	}

	return map[string]any{
		"Name":       "Landmark City",
		"Loc":        "Chennai North* · DTCP approved",
		"Sub":        "24 plots around a central spine road — extent ~9.8 acres*. A concept render stands in for project photography until client assets are supplied.",
		"AvailBadge": fmt.Sprintf("%d of %d plots available*", counts["available"], counts["all"]),
		"Plots":      plots,
		"PlotsJSON":  string(plotsJSON),
		"Legend":     legend,
		"Gallery":    gallery,
		"LocRows":    locRows,
	}
}

// enIN formats an integer with western comma grouping — matches
// toLocaleString('en-IN') for the sample range (≤ 9,999).
func enIN(n int) string {
	s := strconv.Itoa(n)
	if len(s) <= 3 {
		return s
	}
	var b strings.Builder
	for i, r := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteRune(r)
	}
	return b.String()
}

// Page carries everything the layout and page templates need.
type Page struct {
	Title       string
	HeaderSolid bool // solid header (pages without a dark hero)
	Data        any

	// per-artifact wordings (footer compliance line + floating
	// sample-data chip differ between the approved screens)
	DtcpNote    string
	SampleFlag  string
	FlagDismiss bool // 03-detail's chip carries a dismiss button
}

var pages = map[string]string{
	"home":           "home.html",
	"projects":       "projects.html",
	"project-detail": "project-detail.html",
	"about":          "about.html",
	"contact":        "contact.html",
	"placeholder":    "placeholder.html",
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
	base, err := template.New("base").Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"mul": func(a, b int) int { return a * b },
	}).ParseFS(os.DirFS("assets"),
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
		fmt.Sprintf("history.pushState(null, '', %q); document.title = %q; "+
			"window.scrollTo({top: 0, behavior: 'instant'})",
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

// projectsData is sample content mirroring the approved 02-projects
// artifact. Real data modelling comes in a later iteration.
func projectsData() []PRow {
	return []PRow{
		{
			Num: "01", Loc: "Chennai North", LocReq: true,
			Name: "Landmark City", DTCP: true,
			Status: "15 available", StatusReq: true,
			Specs: []PSpec{
				{K: "Extent", V: "~9.8 acres", Req: true},
				{K: "Plots", V: "24", Req: true},
				{K: "Plot size", V: "600–2,400 sq.ft", Req: true},
				{K: "Price", V: "₹1,850 / sq.ft", Req: true},
			},
			Avail: 15, Resv: 3, Sold: 6,
			FrontImg: "card-landmark.jpg",
			FrontAlt: "Landmark City — DTCP approved signboard at the layout entrance",
			L2Img:    "proof-roads.jpg",
			Href:     "/projects/landmark-city",
		},
		{
			Num: "02", Loc: "Chennai North", LocReq: true,
			Name: "Srivelan Nagar", DTCP: true,
			Status: "Now selling", StatusStar: true,
			Specs: []PSpec{
				{K: "Extent", V: "~7.4 acres", Req: true},
				{K: "Plots", V: "36", Req: true},
				{K: "Plot size", V: "800–2,600 sq.ft", Req: true},
				{K: "Price", V: "On request", Req: true},
			},
			Avail: 28, Resv: 4, Sold: 4,
			FrontImg: "card-srivelan.jpg",
			FrontAlt: "Srivelan Nagar — DTCP approved entrance signboard",
			L2Img:    "proof-site-2.jpg",
			Href:     "/projects/srivelan-nagar",
		},
		{
			Num: "03", Loc: "Thiruppalai Nadu", LocReq: true,
			Name: "Chandrayan Nagar", DTCP: true,
			Status: "Limited plots", StatusReq: true,
			Specs: []PSpec{
				{K: "Extent", V: "~6.2 acres", Req: true},
				{K: "Plots", V: "30", Req: true},
				{K: "Plot size", V: "600–1,800 sq.ft", Req: true},
				{K: "Price", V: "On request", Req: true},
			},
			Avail: 19, Resv: 4, Sold: 7,
			FrontImg: "card-chandrayan.jpg",
			FrontAlt: "Chandrayan Nagar — DTCP approved entrance signboard",
			L2Img:    "proof-site-1.jpg",
			Href:     "/projects/chandrayan-nagar",
		},
		{
			Num: "04", Loc: "OMR, Chennai",
			Name:  "Fourth layout — in preparation",
			Quiet: true,
			Muted: "A new VSS layout on the OMR corridor is being prepared for " +
				"release. Facts and approval details will be published here once " +
				"verified — nothing is offered until then.",
			Specs: []PSpec{
				{K: "Extent", V: "—"},
				{K: "Plots", V: "—"},
				{K: "Approval", V: "In process"},
				{K: "Price", V: "On request"},
			},
			FrontImg: "aerial-close.jpg",
			FrontAlt: "Aerial concept render — new layout under preparation",
			L2Img:    "aerial-rows.jpg",
		},
	}
}

// aboutData is sample content mirroring the approved 05-about
// artifact. Client-supplied facts (founding year, founder's name,
// pull-quote wording, pipeline) are marked * until they arrive.
func aboutData() map[string]any {
	req := func(s string) string {
		return `<span class="font-semibold text-brass">` + s + `</span>`
	}
	return map[string]any{
		"Values": []map[string]string{
			{"No": "01", "Title": "Approvals before sales",
				"Body": "Nothing is offered until the layout papers are in order. Approval documents are shared openly with every buyer."},
			{"No": "02", "Title": "The plan is the ground truth",
				"Body": "Plot sizes, road widths and boundaries on the plan match the ground. If a survey disagrees, we fix the ground."},
			{"No": "03", "Title": "Honest pricing",
				"Body": "One clear price per plot, stated upfront. No pressure escalation, no hidden charges introduced after booking."},
			{"No": "04", "Title": "Handover with papers",
				"Body": "Registration support, patta transfer guidance and a defined handover checklist — the plot is yours in fact, not just in promise."},
		},
		"Timeline": []struct {
			Yr, Title string
			Body      template.HTML
			Pending   bool
		}{
			{"Origins*", "The lesson",
				template.HTML("Working inside another real estate firm, our founder saw documentation that left buyers worried and confused. " + req("Founding year awaited from client.*")), true},
			{"Founded*", "VSS Associates",
				template.HTML("He left to build a firm where the papers are clear and the ground is honest — and grew it from the bottom up, plot by plot. " + req("Year and founder's name awaited from client.*")), true},
			{"Today", "4 layouts delivered",
				template.HTML("Chandrayan and Srivelan are live on this site with a further " + req("4 layouts to be listed — client input required.*")), false},
			{"Next", "Upcoming layouts*",
				template.HTML(req("Pipeline projects and target years awaited from client.*")), true},
		},
		"Tiles": []struct {
			Img, Alt, Cap string
			Lead, Right   bool
		}{
			{"proof-build.jpg", "Completed compound wall at a VSS layout",
				"Compound wall at handover · project name to be confirmed*", true, false},
			{"proof-stakes.jpg", "Plot number board at a VSS layout",
				"Plot stakes and numbering on the ground", false, false},
			{"proof-visit.jpg", "Site visit walkthrough at a VSS layout",
				"Site visit walkthrough", false, false},
			{"proof-roads.jpg", "Finished internal roads at a VSS layout",
				"Finished internal roads", false, false},
			{"aerial-close.jpg", "Aerial close-up of a completed VSS layout",
				"Layout overview on handover day", false, true},
		},
		"Stats": []struct {
			Target int
			Label  string
			Req    bool
		}{
			{4, "Layouts delivered", false},
			{600, "Plots delivered", true},
			{120, "Acres developed", true},
			{100, "% DTCP approved", true},
		},
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	render(w, r, "about", Page{
		Title:       "About — VSS Associates",
		Data:        aboutData(),
		DtcpNote:    "Every layout sold here is DTCP approved — approval numbers published per project",
		SampleFlag:  "story beats, years & proof numbers marked * await client supply",
		FlagDismiss: true,
	})
}

// contactData holds only the repeating collections (info rows, visit
// steps); headline and form copy live directly in contact.html.
func contactData() map[string]any {
	return map[string]any{
		"Info": []struct {
			Icon template.HTML
			K    string
			V    template.HTML
		}{
			{`<svg viewBox="0 0 24 24"><path d="M5 4h4l2 5-2.5 1.5a12 12 0 0 0 5 5L15 13l5 2v4a2 2 0 0 1-2 2A16 16 0 0 1 3 6a2 2 0 0 1 2-2z"/></svg>`,
				"Phone", "044 4208 7913"},
			{`<svg viewBox="0 0 24 24"><rect x="3" y="5" width="18" height="14" rx="2"/><path d="M3 7l9 6 9-6"/></svg>`,
				"Email", "contact@vssassociates.com"},
			{`<svg viewBox="0 0 24 24"><path d="M12 21s-7-5.5-7-11a7 7 0 1 1 14 0c0 5.5-7 11-7 11z"/><circle cx="12" cy="10" r="2.5"/></svg>`,
				"Office", template.HTML("No. 11, West Karikalan Street, 1st Floor,<br>(Opp. to SBI Bank), Adambakkam,<br>Chennai – 600 088")},
			{`<svg viewBox="0 0 24 24"><circle cx="12" cy="12" r="9"/><path d="M12 7v5l3 3"/></svg>`,
				"Working hours", template.HTML("Mon – Sat, 8.00 – 18.00<br>Sunday closed")},
		},
		"Steps": []struct{ T, D string }{
			{"Share your interest", "Form or phone — whichever is easier."},
			{"We confirm the slot", "Within a working day, by phone."},
			{"Walk the layout", "Plot plan and approval documents in hand."},
			{"Decide at leisure", "No on-the-spot pressure; prices hold as stated."},
		},
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	render(w, r, "contact", Page{
		Title:    "Contact / Book a Visit — VSS Associates",
		Data:     contactData(),
		DtcpNote: "All layouts marketed are subject to DTCP approval — approval numbers published per project",
	})
}

func projects(w http.ResponseWriter, r *http.Request) {
	rows := projectsData()
	for i := range rows {
		rows[i].First = i == 0
		rows[i].Even = i%2 == 1
	}
	render(w, r, "projects", Page{
		Title:      "Projects — VSS Associates",
		Data:       map[string]any{"Rows": rows},
		DtcpNote:   "DTCP status published per project — approved before offered",
		SampleFlag: "extents, counts, prices marked * await client supply",
	})
}

func projectDetail(w http.ResponseWriter, r *http.Request) {
	render(w, r, "project-detail", Page{
		Title:       "Landmark City — VSS Associates",
		Data:        projectDetailData(),
		DtcpNote:    "All layouts DTCP approved — approval numbers published per project",
		SampleFlag:  "plot facts, prices & stats marked * await client supply",
		FlagDismiss: true,
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, r, "home", Page{
		Title:      "VSS Associates — DTCP-approved plots in Chennai",
		Data:       homeData(),
		DtcpNote:   "✓ Every layout sold here is DTCP approved — approval numbers published on each project page.",
		SampleFlag: "values marked * await client supply",
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
	mux.HandleFunc("GET /projects", projects)
	mux.HandleFunc("GET /projects/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/projects/" {
			http.Redirect(w, r, "/projects", http.StatusPermanentRedirect)
			return
		}
		notFound(w, r)
	})
	mux.HandleFunc("GET /projects/{slug}", func(w http.ResponseWriter, r *http.Request) {
		if r.PathValue("slug") == "landmark-city" {
			projectDetail(w, r)
			return
		}
		placeholder("Project")(w, r)
	})
	mux.HandleFunc("GET /about", about)
	mux.HandleFunc("GET /contact", contact)
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
