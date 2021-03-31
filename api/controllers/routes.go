package controllers

import "github.com/kartuludus/wallaster-be/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")


	//Users routes
	s.Router.HandleFunc("/customers", middlewares.SetMiddlewareJSON(s.CreateCustomer)).Methods("POST")
	s.Router.HandleFunc("/customers", middlewares.SetMiddlewareJSON(s.GetCustomers)).Methods("GET")
	s.Router.HandleFunc("/customers/{id}", middlewares.SetMiddlewareJSON(s.GetCustomer)).Methods("GET")
	s.Router.HandleFunc("/customers/search/{search}", middlewares.SetMiddlewareJSON(s.FindCustomers)).Methods("GET")
	s.Router.HandleFunc("/customers/update", middlewares.SetMiddlewareJSON(s.UpdateCustomer)).Methods("POST")

}