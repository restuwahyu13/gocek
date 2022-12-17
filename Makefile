test:
	go test --cover -v --failfast .

dc-test:
	docker build -t gocek --compress . && docker run gocek go test --cover -v --failfast .