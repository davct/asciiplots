objects = main.go hist.go util.go plot.go
run = go run


wtest:
	go test
	$(run) $(objects)

main:
	$(run) $(objects)
