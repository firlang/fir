all:
	make clean
	make build
	make run

clean:
	rm -rf debug/fir

build:
	go build -o debug/fir .

run:
	./debug/fir