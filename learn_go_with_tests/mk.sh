#!/bin/bash
l="
01_hello_world
02_integers
03_for
04_arrays
05_structs
06_pointers
07_maps
08_di
09_mocking
10_concurrency
11_select
"

echo "$l" | while read d; do
    if [ -z "$d" ]; then
        continue
    fi
    mkdir -p "$d" && echo "package main" > "${d}/main.go"


done
