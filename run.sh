#!/bin/bash

go build -o bookings cmd/web/*.go 
./bookings -dbname=bookings -dbuser=sshtepan -cache=false -production=false
