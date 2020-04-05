package controllers

func (s *Server) initRoutes() {
	// Category routes
	s.Router.HandleFunc("/api/categories", s.GetCategories).Methods("GET")
	s.Router.HandleFunc("/api/categories", s.CreateCategory).Methods("POST")
	s.Router.HandleFunc("/api/categories/{id}", s.UpdateCategory).Methods("PUT")
	s.Router.HandleFunc("/api/categories/{id}", s.DeleteCategory).Methods("DELETE")

	// Product routes
	s.Router.HandleFunc("/api/products", s.GetProducts).Methods("GET")
	s.Router.HandleFunc("/api/products/category/{id}", s.GetProductsByCategoryId).Methods("GET")
	s.Router.HandleFunc("/api/products", s.CreateProduct).Methods("POST")
	s.Router.HandleFunc("/api/products/{id}", s.UpdateProduct).Methods("PUT")
	s.Router.HandleFunc("/api/products/{id}", s.DeleteProduct).Methods("DELETE")
}
