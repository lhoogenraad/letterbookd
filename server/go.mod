module server

go 1.19

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v1.5.5
	github.com/gorilla/schema v1.3.0
	github.com/sirupsen/logrus v1.9.3
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)

replace server/tools => ../server
