package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// User Указываем структуру пользователя
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// Server Указываем структуру Сервера
type Server struct {
	// * - используем для распаковки большой сущности
	*mux.Router
	// Создаем массив пользователей
	users []User
}

// NewServer Функция для создания сервера
func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		users:  []User{},
	}
	s.routes()
	return s
}

// Создаем руты, на которых будет работать сервер
func (s *Server) routes() {
	s.HandleFunc("/users", s.getUsers()).Methods("GET")
	s.HandleFunc("/users", s.createUser()).Methods("POST")
	s.HandleFunc("/", s.ping()).Methods("GET")
}

func (s *Server) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем заголовок application/json, т.к. возвращаем данные в формате json
		w.Header().Set("Content-Type", "application/json")
		// Записываем в тело ответа pong и энкодим его в json
		if err := json.NewEncoder(w).Encode("pong"); err != nil {
			// Рейзим http ошибку, в случае каких - то проблем
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) getUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) createUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем заголовок application/json, т.к. возвращаем данные в формате json
		w.Header().Set("Content-Type", "application/json")
		// Создаем объект структуры User
		var i User
		// Создаем объект decoder json и декодируем тело запроса,
		// а также распаковываем и засовываем это в i
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			// Рейзим ошибку, если что - то не так
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		i.ID = uuid.New()
		s.users = append(s.users, i)
		// Возвращаем 201, в случае успешного создания
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
