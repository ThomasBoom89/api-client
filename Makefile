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

test:
	make test-backend
	make test-frontend-e2e
	make test-frontend-unit

test-frontend-e2e:
	cd frontend && npm run test:integration

test-frontend-unit:
	cd frontend && npm run test:unit

test-backend:
	go test ./...

frontend-lint:
	cd frontend && npm run lint

frontend-format:
	cd frontend && npm run format

frontend-check:
	cd frontend && npm run check
