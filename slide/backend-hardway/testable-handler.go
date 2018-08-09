func createHandler(db *sql.DB) http.Handler {
	// init data
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

	})
}

// ------------------------------

type handler struct {
	db *sql.DB
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
