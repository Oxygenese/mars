module github.com/mars-projects/mars/app/auth

go 1.16

require (
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20220310144244-ac99a5c877c4
	github.com/go-kratos/kratos/v2 v2.2.0
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/mars-projects/mars v0.0.0-20220314054335-f6b1dcb439ba
	github.com/mars-projects/oauth2/v4 v4.4.2
	golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70
)

replace github.com/mars-projects/mars => ../../
