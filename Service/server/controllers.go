package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Controllers struct {
	router *httprouter.Router
}

func (c *Controllers) ConnectControllers(s *ServiceOptions, controllers func()) error {
	c.router = httprouter.New()
	controllers()
	http.ListenAndServe(s.Host+":"+s.Port, c.router)
	return nil
}

func (c *Controllers) ShowOrderByID(cache *DataBaseCache) {

	c.router.GET("/showOrder/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		order := cache.GetCache(ps[0].Value)

		tmpl, err := template.ParseFiles("templates/showOrderByID.html")
		if err != nil {
			log.Print(err)
		} else {
			err = tmpl.Execute(w, order)
			if err != nil {
				log.Print(err)
			}
		}

	})
}
func (c *Controllers) CreateNewOrder(mb *MessageBroker, s *ServiceOptions) {

	c.router.GET("/createNewOrder", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		tmpl, err := template.ParseFiles("templates/createNewOrder.html")
		if err != nil {
			log.Print(err)
		} else {
			err = tmpl.Execute(w, s)
			if err != nil {
				log.Print(err)
			}
		}

	})
	c.router.PUT("/createNewOrder", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		t, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		} else {
			mb.PublishMessageBroker(t)
		}
	})
}
