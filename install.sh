#!/bin/sh
cd ./backend/krushr || return
go get github.com/stanhoenson/krushr
cd - || return

cd ./frontend/krushr || return
npm install
cd - || return
