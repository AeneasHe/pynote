@echo off
go generate
go build -ldflags "-H windowsgui" -o dist/windows/Pynote.exe
