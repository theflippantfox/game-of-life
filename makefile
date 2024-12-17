dev:
	go run main.go

build:clean
	go -o bin/game_of_life main.go

clean:
       rm -rf bin
