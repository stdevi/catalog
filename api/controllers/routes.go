package controllers

func (s *Server) initRoutes() {
	s.Router.HandleFunc("/api/categories", s.GetCategories).Methods("GET")
	s.Router.HandleFunc("/api/categories", s.CreateCategory).Methods("POST")
}
