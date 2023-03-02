package shortener

// RedirectRepository is an interface for our storages...
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
