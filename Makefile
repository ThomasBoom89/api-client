install:
	make install-wails
	make install-dependencies

dev:
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

test-frontend-e2e-ui:
	cd frontend && npm run test:integration-ui

test-frontend-unit:
	cd frontend && npm run test:unit

test-frontend-unit-watch:
	cd frontend && npm run test:unit-watch

test-frontend-unit-coverage:
	cd frontend && npm run test:unit-coverage

test-backend:
	go test ./...

backend-format-check:
	gofmt -l .

backend-format:
	gofmt -w .

frontend-format-check:
	cd frontend && npm run format:check

frontend-format:
	cd frontend && npm run format

frontend-check:
	cd frontend && npm run check
