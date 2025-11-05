Write-Host "Initializing ERP database..."
Set-Location -Path "tools"
go run init_db.go
Write-Host "Database initialization completed."