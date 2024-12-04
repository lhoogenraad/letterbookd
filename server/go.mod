module github.com/lhoogenraad/letterbookd

go 1.20

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v1.5.5
	github.com/go-chi/cors v1.2.1
	github.com/go-sql-driver/mysql v1.8.1
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/crypto v0.22.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)

replace server/tools => ../server
