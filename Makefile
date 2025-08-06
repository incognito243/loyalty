# Default log files
DAY1 ?= mock_data/log_day0.csv
DAY2 ?= mock_data/log_day1.csv

.PHONY: all build clean test run

all: clean build test run

build:
	go build -o loyalty main.go

test:
	go test -v ./loyal

run:
	go run main.go $(DAY1) $(DAY2)

clean:
	rm -f loyalty
