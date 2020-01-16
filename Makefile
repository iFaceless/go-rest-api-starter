all: web admin

web:
	go build -o bin/starter-web cmd/web/main.go

admin:
	go build -o bin/starter-admin cmd/admin/main.go