package pkg

import "net/http"

type Sender interface {
	execute(*http.Client, *http.Request)
}
