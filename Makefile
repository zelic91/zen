test:
	go run main.go run -c zen.yaml -t ../testgen
	cd ../testgen && make init && code .

untest:
	rm -rf ../testgen

.PHONY: test untest