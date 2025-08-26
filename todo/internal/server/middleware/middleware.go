package middleware

type Middleware struct{}

func New() *Middleware {
	return &Middleware{}
}

//TODO: withAuth
//TODO: withLogging
//TODO: withRecover
//TODO: withCheck (check equaliation id from token and id from url)
