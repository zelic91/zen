test:
	go run main.go create goservice -m superman.test -d testgen
	cd testgen && mv env.sample .env && make init

untest:
	rm -rf testgen

.PHONY: test untest