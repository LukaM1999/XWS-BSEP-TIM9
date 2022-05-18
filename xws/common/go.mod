module dislinkt/common

go 1.18

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	go.mongodb.org/mongo-driver v1.9.1
	golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f
	google.golang.org/genproto v0.0.0-20220429170224-98d788798c3e
	google.golang.org/grpc v1.46.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace dislinkt/common => ../common

//replace dislinkt/security_service => ../security_service
