package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/arisnotargon/t_write_test/models/request"
)

type Route struct {
	app *App
}

func NewRoute(app *App) *Route {
	return &Route{app: app}
}

// write response,common function
func (r *Route) write(w http.ResponseWriter, response interface{}, code int) {
	b, err := json.Marshal(response)
	if err != nil {
		code = 500
		b = []byte("response marshal error")
	}

	w.WriteHeader(code)
	w.Write(b)
}

func (r *Route) writeNotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	w.Write([]byte("404 page not found\n"))
}

func (r *Route) InitRoute() {
	http.HandleFunc("/health_check", r.healthCheckRoute)
	http.HandleFunc("/age", r.ageRoute)
	http.HandleFunc("/car", r.carRoute)
	http.HandleFunc("/rate", r.rateRoute)
	http.HandleFunc("/buffer", r.bufferRoute)

}

func (r *Route) healthCheckRoute(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(200)
		w.Write([]byte("health ok\n"))
	} else {
		r.writeNotFound(w)
	}
}

func (r *Route) ageRoute(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res, code := r.app.service.GetAge()
		r.write(w, res, code)
	} else if req.Method == "POST" {
		bodyB, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("request body read error,err =>>> [%+v]", err)))
			return
		}
		p := &request.SetAge{}
		err = json.Unmarshal(bodyB, p)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("request body bind params error,err =>>> [%+v]", err)))
			return
		}
		res, code := r.app.service.SetAge(p)
		r.write(w, res, code)

	} else {
		r.writeNotFound(w)
	}
}

func (r *Route) carRoute(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res, code := r.app.service.GetCar()
		r.write(w, res, code)
	} else {
		r.writeNotFound(w)
	}
}

func (r *Route) rateRoute(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res, code := r.app.service.GetRate()
		r.write(w, res, code)
	} else {
		r.writeNotFound(w)
	}
}

func (r *Route) bufferRoute(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res, code := r.app.service.GetBuffer()
		r.write(w, res, code)
	} else {
		r.writeNotFound(w)
	}
}
