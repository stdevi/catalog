package controllers

func (s *Server) initRoutes() {
	s.Router.HandleFunc("/api/categories", s.GetCategories).Methods("GET")
}
