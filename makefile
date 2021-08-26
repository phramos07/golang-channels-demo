SHELL := /bin/bash

.PHONY: write read routine

write:
	go run ./deadlocks/deadlockWrite.go

read:
	go run ./deadlocks/deadlockRead.go

routine:
	go run ./deadlocks/noDeadlock.go

seq:
	go run ./sequential/sequential.go $(NUMBER_OF_REQS)

cur:
	go run ./concurrent/concurrent.go $(NUMBER_OF_REQS)

time-seq: 
	time $(MAKE) seq

time-cur:
	time $(MAKE) cur
