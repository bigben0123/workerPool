@echo off

setlocal

if exist install.cmd goto ok
echo install.cmd must be run from its folder
goto end

: ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

rem gofmt -w src

go install main

set GOPATH=OLDGOPATH

:end
echo finished