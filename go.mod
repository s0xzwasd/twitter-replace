module twitter-replace

go 1.17

replace github.com/Nerzal/gocloak/v10 v10.0.1 => github.com/shibumi/gocloak/v10 v10.1.0

require github.com/Nerzal/gocloak/v10 v10.0.1

require (
	github.com/go-resty/resty/v2 v2.6.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	golang.org/x/net v0.0.0-20210928044308-7d9f5e0b762b // indirect
)
