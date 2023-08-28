test:
	go run main.go create goservice -m superman.test -d testgen
	cd testgen && make init

untest:
	rm -rf testgen

create:
	go run main.go create

.PHONY: test untest create