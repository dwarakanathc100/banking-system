


test:
	go test ./... -v
bench-mark:
	go test ./useraccount -bench=. -benchmem
future-test:
	go test -v bdd_test.go