package main

import (
	"compress/gzip"
	"mime"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.ListenAndServe("[::1]:3000", gzipMiddleware(http.HandlerFunc(index)))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := []byte(`
		<!doctype html>
		<title>Compression Test</title>
		<h1>Hello</h1>
		<p>The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymphs grab quick-jived waltz. Brick quiz whangs jumpy veldt fox. Bright vixens jump; dozy fowl quack. Quick wafting zephyrs vex bold Jim. Quick zephyrs blow, vexing daft Jim. Sex-charged fop blew my junk TV quiz. How quickly daft jumping zebras vex. Two driven jocks help fax my big quiz. Quick, Baz, get my woven flax jodhpurs! "Now fax quiz Jack!" my brave ghost pled. Five quacking zephyrs jolt my wax bed. Flummoxed by job, kvetching W. zaps Iraq. Cozy sphinx waves quart jug of bad milk. A very bad quack might jinx zippy fowls. Few quips galvanized the mock jury box. Quick brown dogs jump over the lazy fox. The jay, pig, fox, zebra, and my wolves quack! Blowzy red vixens fight for a quick jump. Joaquin Phoenix was gazed by MTV for luck. A wizard’s job is to vex chumps quickly in fog. Watch "Jeopardy!", Alex Trebek's fun TV quiz game. Woven silk pyjamas exchanged for blue quartz. Brawny gods just</p>
		<p>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem. Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc,</p>
	`)
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func gzipMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}

		if r.Header.Get("Sec-WebSocket-Key") != "" {
			h.ServeHTTP(w, r)
			return
		}

		wh := w.Header()

		if wh.Get("Content-Encoding") != "" {
			h.ServeHTTP(w, r)
			return
		}

		wh.Add("Vary", "Accept-Encoding")

		gw := &gzipResponseWriter{
			ResponseWriter: w,
		}
		defer gw.Close()

		h.ServeHTTP(gw, r)
	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	gzipWriter  *gzip.Writer
	wroteHeader bool
}

var allowCompressType = map[string]bool{
	"text/plain":               true,
	"text/html":                true,
	"text/css":                 true,
	"text/xml":                 true,
	"text/javascript":          true,
	"application/x-javascript": true,
	"application/xml":          true,
}

func (w *gzipResponseWriter) init() {
	h := w.Header()

	if h.Get("Content-Encoding") != "" {
		return
	}

	if sl := h.Get("Content-Length"); sl != "" {
		l, _ := strconv.Atoi(sl)
		if l > 0 && l < 860 {
			return
		}
	}

	// skip if no match type
	ct, _, err := mime.ParseMediaType(h.Get("Content-Type"))
	if err != nil {
		ct = "application/octet-stream"
	}

	if !allowCompressType[ct] {
		return
	}

	// START 1 OMIT
	w.gzipWriter = gzip.NewWriter(w.ResponseWriter)
	// END 1 OMIT
	h.Del("Content-Length")
	h.Set("Content-Encoding", "gzip")
}

func (w *gzipResponseWriter) Write(p []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	if w.gzipWriter != nil {
		return w.gzipWriter.Write(p)
	}
	return w.ResponseWriter.Write(p)
}

func (w *gzipResponseWriter) Close() {
	if w.gzipWriter == nil {
		return
	}
	// START 2 OMIT
	w.gzipWriter.Close()
	// END 2 OMIT
}

func (w *gzipResponseWriter) WriteHeader(code int) {
	if w.wroteHeader {
		return
	}
	w.wroteHeader = true
	w.init()
	w.ResponseWriter.WriteHeader(code)
}
