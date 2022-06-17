module dislinkt/common

go 1.18

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3
	github.com/sirupsen/logrus v1.8.1
	go.mongodb.org/mongo-driver v1.9.1
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	google.golang.org/genproto v0.0.0-20220617124728-180714bec0ad
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.15.6 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	golang.org/x/net v0.0.0-20220615171555-694bf12d69de // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	golang.org/x/sys v0.0.0-20220615213510-4f61da869c0c // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace dislinkt/common => ../common

//replace dislinkt/security_service => ../security_service
