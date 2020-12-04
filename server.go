package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"path"

	// Logging
	"github.com/unrolled/logger"

	// Stats/Metrics
	"github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-metrics/exp"
	"github.com/thoas/stats"

	"github.com/NYTimes/gziphandler"
	"github.com/julienschmidt/httprouter"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

// Counters ...
type Counters struct {
	r metrics.Registry
}

func NewCounters() *Counters {
	counters := &Counters{
		r: metrics.NewRegistry(),
	}
	return counters
}

func (c *Counters) Inc(name string) {
	metrics.GetOrRegisterCounter(name, c.r).Inc(1)
}

func (c *Counters) Dec(name string) {
	metrics.GetOrRegisterCounter(name, c.r).Dec(1)
}

func (c *Counters) IncBy(name string, n int64) {
	metrics.GetOrRegisterCounter(name, c.r).Inc(n)
}

func (c *Counters) DecBy(name string, n int64) {
	metrics.GetOrRegisterCounter(name, c.r).Dec(n)
}

// Server ...
type Server struct {
	bind      string
	root      string
	config    Config
	templates *Templates
	router    *httprouter.Router

	// Logger
	logger *logger.Logger

	// Stats/Metrics
	counters *Counters
	stats    *stats.Stats
}

func (s *Server) render(name string, w http.ResponseWriter, ctx interface{}) {
	buf, err := s.templates.Exec(name, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type IndexContext struct {
	Root     string
	NoteList []*Note
}

// IndexHandler ...
func (s *Server) IndexHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s.counters.Inc("n_index")

		var noteList []*Note
		err := db.All(&noteList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx := &IndexContext{
			Root:     s.root,
			NoteList: noteList,
		}

		s.render("index", w, ctx)
	}
}

// EditContext ...
type EditContext struct {
	Root  string
	ID    int
	Title string
	Body  string
}

// EditHandler ...
func (s *Server) EditHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		s.counters.Inc("n_edit")

		var id string

		id = p.ByName("id")
		if id == "" {
			id = r.FormValue("id")
		}

		var (
			note Note
			body []byte
		)

		if id != "" {
			i, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				log.Printf("error parsing id %s: %s", id, err)
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}

			err = db.One("ID", i, &note)
			if err != nil {
				log.Printf("error looking up note %d: %s", i, err)
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}

			body, err = note.LoadBody(s.config.data)
			if err != nil {
				log.Printf("error loading note body for %d: %s", i, err)
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}
		}

		ctx := &EditContext{
			Root:  s.root,
			ID:    note.ID,
			Title: note.Title,
			Body:  string(body),
		}

		s.render("edit", w, ctx)
	}
}

// ViewContext ...
type ViewContext struct {
	Root  string
	ID    int
	Title string
	HTML  template.HTML
}

// ViewHandler ...
func (s *Server) ViewHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		s.counters.Inc("n_view")

		var id string

		id = p.ByName("id")
		if id == "" {
			id = r.FormValue("id")
		}

		if id == "" {
			log.Printf("no id specified to view: %s", id)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Printf("error parsing id %s: %s", id, err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		var note Note
		err = db.One("ID", i, &note)
		if err != nil {
			log.Printf("error looking up note %d: %s", i, err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		body, err := note.LoadBody(s.config.data)
		if err != nil {
			log.Printf("error loading note body for %d: %s", i, err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		unsafe := blackfriday.MarkdownCommon(body)
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

		ctx := &ViewContext{
			Root:  s.root,
			ID:    note.ID,
			Title: note.Title,
			HTML:  template.HTML(html),
		}

		s.render("view", w, ctx)
	}
}

// DeleteHandler ...
func (s *Server) DeleteHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		s.counters.Inc("n_delete")

		var id string

		id = p.ByName("id")
		if id == "" {
			id = r.FormValue("id")
		}

		if id == "" {
			log.Printf("no id specified to delete: %s", id)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Printf("error parsing id %s: %s", id, err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		var note Note
		err = db.One("ID", i, &note)
		if err != nil {
			log.Printf("error looking up note %d: %s", i, err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		err = db.DeleteStruct(&note)
		if err != nil {
			log.Printf("error deleting note %d: %s", i, err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		err = note.DeleteBody(s.config.data)
		if err != nil {
			log.Printf("error deleting note body %d: %s", i, err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, s.root, http.StatusFound)
	}
}

// SaveHandler ...
func (s *Server) SaveHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		s.counters.Inc("n_save")

		var note Note

		redirect := s.root

		sid := p.ByName("id")
		id := SafeParseInt(sid, 0)

		title := r.FormValue("title")
		body := r.FormValue("body")
		// TODO: Save tags
		// tags := r.FormValue("tags")

		if id > 0 {
			err := db.One("ID", id, &note)
			if err != nil {
				log.Printf("error looking up note %d: %s", id, err)
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}
			redirect += "view/" + sid
		} else {
			note = *NewNote(title)
		}

		// TODO: Save tags
		err := db.Save(&note)
		if err != nil {
			log.Printf("error saving note: %s", err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		err = note.SaveBody(s.config.data, []byte(body))
		if err != nil {
			log.Printf("error saving note body: %s", err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, redirect, http.StatusFound)
	}
}

// StatsHandler ...
func (s *Server) StatsHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		bs, err := json.Marshal(s.stats.Data())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(bs)
	}
}

// ListenAndServe ...
func (s *Server) ListenAndServe() {
	log.Fatal(
		http.ListenAndServe(
			s.bind,
			s.logger.Handler(
				s.stats.Handler(
					gziphandler.GzipHandler(
						s.router,
					),
				),
			),
		),
	)
}

func (s *Server) initRoutes() {
	if s.root[len(s.root)-1] != '/' {
		s.root += "/"
	}
	if s.root[0] != '/' {
		s.root = "/" + s.root
	}
	s.router.Handler("GET", path.Join(s.root, "/debug/metrics"), exp.ExpHandler(s.counters.r))
	s.router.GET(path.Join(s.root, "/debug/stats"), s.StatsHandler())

	if s.root == "/" {
		s.router.GET("/", s.IndexHandler())
	} else {
		s.router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			http.Redirect(w, r, s.root, http.StatusTemporaryRedirect)
		})
		s.router.GET(s.root, s.IndexHandler())
	}

	s.router.GET(path.Join(s.root, "/new"), s.EditHandler())
	s.router.POST(path.Join(s.root, "/save"), s.SaveHandler())
	s.router.POST(path.Join(s.root, "/save/:id"), s.SaveHandler())

	s.router.GET(path.Join(s.root, "/edit/:id"), s.EditHandler())
	s.router.GET(path.Join(s.root, "/view/:id"), s.ViewHandler())
	s.router.GET(path.Join(s.root, "/delete/:id"), s.DeleteHandler())
}

func loadAsset(filepath string) string {
	file, ok := Assets.Files[filepath]
	if !ok {
		panic("can't open: " + filepath)
	}
	return string(file.Data)
}

// NewServer ...
func NewServer(bind string, config Config, root string) *Server {
	server := &Server{
		bind:      bind,
		root:      root,
		config:    config,
		router:    httprouter.New(),
		templates: NewTemplates("base"),

		// Logger
		logger: logger.New(logger.Options{
			Prefix:               "notes",
			RemoteAddressHeaders: []string{"X-Forwarded-For"},
			OutputFlags:          log.LstdFlags,
		}),

		// Stats/Metrics
		counters: NewCounters(),
		stats:    stats.New(),
	}

	// Templates
	

	indexTemplate := template.New("index")
	template.Must(indexTemplate.Parse(loadAsset("/index.html")))
	template.Must(indexTemplate.Parse(loadAsset("/base.html")))

	editTemplate := template.New("edit")
	template.Must(editTemplate.Parse(loadAsset("/edit.html")))
	template.Must(editTemplate.Parse(loadAsset("/base.html")))

	viewTemplate := template.New("view")
	template.Must(viewTemplate.Parse(loadAsset("/view.html")))
	template.Must(viewTemplate.Parse(loadAsset("/base.html")))

	server.templates.Add("edit", editTemplate)
	server.templates.Add("view", viewTemplate)
	server.templates.Add("index", indexTemplate)

	server.initRoutes()

	return server
}
