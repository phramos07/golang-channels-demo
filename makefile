.PHONY: write read

write:
	go run ./deadlockWrite.go

read:
	go run ./deadlockRead.go

