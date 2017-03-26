build:
	cd main; \
	GOARM=5 GOOS=linux GOARCH=arm go build -o homekit

clean:
	rm homekit
