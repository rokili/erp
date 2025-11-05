@echo off
echo Initializing ERP database...
cd tools
go run init_db.go
echo Database initialization completed.
pause