package main

import (
	router "github.com/346285234/bbs-server/router"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

var config Configuration

func init() {
	loadConfig()
}

func main() {

	//logger := log.New(os.Stdout, "bbs", log.LstdFlags | log.Lshortfile)

	r := httprouter.New()

	// Get
	r.GET("/topics", router.Rt.ListTopics)
	r.GET("/topic/:id", GetTopic)

	// Set
	//mux.HandleFunc("/topic/create", router.ListTopics)
	//mux.HandleFunc("/topic/delete", router.ListTopics)


	//tlsConfig := &tls.Config{
	//	// Causes servers to use Go's default ciphersuite preferences,
	//	// which are tuned to avoid attacks. Does nothing on clients.
	//	PreferServerCipherSuites: true,
	//	// Only use curves which have assembly implementations
	//	CurvePreferences: []tls.CurveID{
	//		tls.CurveP256,
	//		tls.X25519, // Go 1.8 only
	//	},
	//	MinVersion: tls.VersionTLS12,
	//	CipherSuites: []uint16{
	//		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	//		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	//		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
	//		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
	//		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	//		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	//
	//		// Best disabled, as they don't provide Forward Secrecy,
	//		// but might be necessary for some clients
	//		// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	//		// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
	//	},
	//}

	server := &http.Server{
		Addr:			config.Address,
		Handler:		r,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
		TLSConfig: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}

func check(fn func (w http.ResponseWriter,
	r *http.Request,
	p httprouter.Params)) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		fn(w, r, p)

	}
}