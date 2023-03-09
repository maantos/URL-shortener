package shortener

// This is part of app & domain

// RedirectService defines bussines logic for redirectService
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
