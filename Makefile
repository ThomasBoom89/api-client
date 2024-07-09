install:
	make install-wails
	wails dev

install-wails:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest

install-dependencies:
	make install-backend-dependencies
	make install-frontend-dependencies

install-backend-dependencies:
	go get ./...

install-frontend-dependencies:
	cd frontend && npm install

ent-generate:
	go generate ./src/ent

test:
	go test ./...

