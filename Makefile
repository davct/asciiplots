objects = main.go hist.go util.go plot.go label.go
run = go run


wtest:
	go test
	$(run) $(objects)

main:
	$(run) $(objects)
