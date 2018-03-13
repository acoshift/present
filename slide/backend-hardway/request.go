type Request struct {
	Method           string
	URL              *url.URL
	Proto            string // "HTTP/1.0"
	Header           Header
	Body             io.ReadCloser
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Host             string
	RemoteAddr       string
	RequestURI       string
	TLS              *tls.ConnectionState
}
