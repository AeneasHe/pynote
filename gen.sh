#!/bin/sh
cd gen
go run gen.go
mv assets.go ../asset
cd ..