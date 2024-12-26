package handler

import "github.com/gorilla/sessions"

var store = sessions.NewCookieStore([]byte("your-secret-key"))
