
all:
	go get -d github.com/xoplog/xop-go@main
	go get -d github.com/francoispqt/onelog@master
	go get -d github.com/phuslu/log@master
	go get -d github.com/rs/zerolog@master
	go test -bench .
