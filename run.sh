#!/bin/bash
redis-server &
go build -o enviro-check *.go && ./enviro-check