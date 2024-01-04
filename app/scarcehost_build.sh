#!/bin/bash

# Set the Go binary name
binary_name="CHANGE_ME"

apt-get update
apt-get install libc6

./$binary_name

