package apiserver

// import (
// 	"net/http"

// 	"github.com/ositlar/go-http/internal/handlers"
// 	"github.com/ositlar/go-http/internal/store"
// )

// type server struct {
// 	router *http.ServeMux
// 	store  store.Cache
// }

// func newServer() *server {
// 	s := &server{
// 		router: http.NewServeMux(),
// 		store:  *store.NewStore(),
// 	}
// 	s.configureRouter()
// 	return s
// }

// func (s *server) configureRouter() {
// 	s.router.HandleFunc("/create_event", handlers.CreateEventHandler) //Post
// 	// mux.HandleFunc("/update_event", handler.UpdateEventHandler)        //Post
// 	// mux.HandleFunc("/delete_event", handler.DeleteEventHandler)        //Post
// 	// mux.HandleFunc("/events_for_day", handler.EventsForDayHandler)     //Get
// 	// mux.HandleFunc("/events_for_week", handler.EventsForWeekHandler)   //Get
// 	// mux.HandleFunc("/events_for_month", handler.EventsForMonthHandler) //Get
// }

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.router.ServeHTTP(w, r)
// }
