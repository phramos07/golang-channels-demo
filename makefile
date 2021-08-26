SHELL := /bin/bash

.PHONY: write read routine

write:
	go run ./deadlockWrite.go

read:
	go run ./deadlockRead.go

routine:
	go run ./noDeadlock.go

seq:
	go run ./sequential.go $(NUMBER_OF_REQS)

time-seq: 
	time $(MAKE) seq