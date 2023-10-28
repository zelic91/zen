test:
	go run main.go create -c zen.yaml -t testgen
	cd testgen && make init && code .

untest:
	rm -rf testgen

create:
	go run main.go create

.PHONY: test untest create