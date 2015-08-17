package api

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type ApiHandler struct {
	Api rest.Api
}

func NewApi() *ApiHander {
	handler := ApiHandler{
		Api: rest.NewApi(),
	}
	handler.Api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/jobs", handler.CreateJob),
		rest.Get("/jobs", handler.GetJobs),
		rest.Get("/jobs/:id", handler.GetJob),
		rest.Put("/jobs/:id", handler.UpdateJob),
		rest.Delete("/jobs/:id", handler.DeleteJob),
		rest.Post("/procs", handler.StartProc),
		rest.Get("/procs/:id", handler.GetProc),
		rest.Delete("/procs/:id", handler.StopProc),
		rest.Get("/jobs/:id/procs", handler.GetProcsByJob),
		rest.Get("/jobs/:id/log", handler.GetProcLogByJob),
	)
	if err != nil {
		log.Fatal(err)
	}
	handler.Api.SetApp(router)
	return &handler
}

func (h *ApiHandler) Start() {
	http.ListenAndServe(":8080", h.Api.MakeHandler())
}

func (h *ApiHandler) CreateJob(w rest.ResponseWriter, r *rest.Request) {
}

func (h *ApiHandler) GetJobs(w rest.ResponseWriter, r *rest.Request) {
}

func (h *ApiHandler) GetJob(w rest.ResponseWriter, r *rest.Request) {
}

func (h *ApiHandler) UpdateJob(w rest.ResponseWriter, r *rest.Request) {
}

func (h *ApiHandler) DeleteJob(w rest.ResponseWriter, r *rest.Request) {
}

func (h *ApiHandler) StartProc(w rest.ResponseWriter, r *rest.Request) {
}

func (h *ApiHandler) GetProc(w rest.ResponseWriter, r *rest.Request) {
}

func (h *ApiHandler) StopProc(w rest.ResponseWriter, r *rest.Request) {
}

func (h *ApiHandler) GetProcsByJob(w rest.ResponseWriter, r *rest.Request) {
}

func (h *ApiHandler) GetProcLogByJob(w rest.ResponseWriter, r *rest.Request) {
}
