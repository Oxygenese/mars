module github.com/mars-projects/mars/app/system

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-kratos/kratos/v2 v2.2.0
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/mars-projects/mars v0.0.0-20220311054936-94f2008a0a4e
	golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
	gorm.io/gorm v1.22.5
)

replace github.com/mars-projects/mars => ../../
