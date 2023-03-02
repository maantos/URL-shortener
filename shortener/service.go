package shortener

// Redirect service defines bussines logic for redirectService
// This is part of app & domain
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
