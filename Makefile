gen:
	protoc --proto_path=proto --go_out=gen proto/*.proto

clean:
	rm -rf gen/*

run:
	go run main.go