.PHONY: write read routine

write:
	go run ./deadlockWrite.go

read:
	go run ./deadlockRead.go

routine:
	go run ./noDeadlock.go