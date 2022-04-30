module dislinkt/api_gateway

go 1.18

replace dislinkt/common => ../common

require (
	dislinkt/common v1.0.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	go.mongodb.org/mongo-driver v1.9.0
	google.golang.org/grpc v1.46.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220429170224-98d788798c3e // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
