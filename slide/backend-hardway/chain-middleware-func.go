func chain(ms ...middleware) middleware {
	return func(h http.Handler) http.Handler {
		for i := len(ms); i > 0; i-- {
			h = ms[i-1](h)
		}
		return h
	}
}

h := chain(
	m1,
	m2,
	m3,
	m4,
)(handler)
