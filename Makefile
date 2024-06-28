BIN_NAME=g-electra-web
VER=0.0.1

ui:
	cd ui && npm install && npm run build && cd ..

build: ui


run: build

