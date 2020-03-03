.PHONY: proto envoy
proto:
	protoc proto/echo.proto \
		--go_out=plugins=grpc:server \
		--js_out=import_style=commonjs,binary:client \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:client

envoy:
	docker build -t echo/envoy:1.0 envoy

build: proto
	cd client && npx webpack main.js && cd ..
