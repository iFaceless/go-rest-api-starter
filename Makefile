all: web admin

web:
	go build -o bin/starter-web cmd/web/main.go

admin:
	go build -o bin/starter-admin cmd/admin/main.go

test:
	DOT_ENV_FILE=`pwd`/.env_unittest go test -p 1 -gcflags "all=-N -l" -coverprofile=coverage.out ./pkg/...