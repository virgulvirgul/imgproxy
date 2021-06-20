package main

import "net/http"

var landingTmpl = []byte(`
<!doctype html>
<html>
	<head><title>image server</title></head>
	<body>
		<h1>image server</h1>
	</body>
</html>
`)

func handleLanding(reqID string, rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Tyle", "text/html")
	rw.WriteHeader(200)
	rw.Write(landingTmpl)
}
