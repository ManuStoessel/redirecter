package Handler

import (
	Store "github.com/ManueStoessel/redirecter/store"
)

type RedirecterHandler struct {
	Store *Store.Store
}

type LongURL struct {
	URL string `json:"url"`
}
