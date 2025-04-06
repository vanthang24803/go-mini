param(
    [string]$Command
)

function Show-Usage {
    Write-Host "Available commands:"
    Write-Host "  .\run.ps1 run   - Run the application"
    Write-Host "  .\run.ps1 reset - Reset and tidy go modules"
}

switch ($Command) {
    "run" {
        Write-Host "Starting application..."
        go run ./cmd/app/main.go
    }
    "reset" {
        Write-Host "Resetting modules..."
        go mod tidy
    }
    default {
        Show-Usage
    }
}