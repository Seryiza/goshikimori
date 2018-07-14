package api

// ShikimoriVersion1 is a sample for adding v2-api (in future)
type ShikimoriVersion1 interface {
	// todo: это только концепт чтобы не забыть направление.
	// Не готово для использования.
	Whoami(*User, error)
}
