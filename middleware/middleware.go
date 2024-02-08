package middleware

import "net/http"

// func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("Authorization")
// 		_, err := autorization.ValidateToken(token)
// 		if err != nil {
// 			forbidden(w, r)
// 			return
// 		}
// 		f(w, r)
// 	}
// }

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No tiene autorización"))
}
