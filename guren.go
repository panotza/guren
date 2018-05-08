package guren

type (
// Guren is the top-level framework instance.
Guren struct {
	Server           *http.Server
	TLSServer        *http.Server
	Listener         net.Listener
}