objects = main.go hist.go util.go plot.go label.go
run = go run
coverout = coverage.out


wtest:
	go test
	$(run) $(objects)

coverage_out:
	go test -coverprofile=${coverout}

coverage: coverage_out
	go tool cover -html=coverage.out

main:
	$(run) $(objects)
