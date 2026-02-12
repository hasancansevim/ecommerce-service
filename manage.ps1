param (
    [string]$cmd = "run"
)

switch ($cmd) {
    "swagger" { swag init }
    "tidy"    { go mod tidy }
    "run"     { swag init; go run . }
    "up"      { swag init; docker-compose up --build }
    "up-d"    { swag init; docker-compose up -d --build }
    "down"    { docker-compose down }
    "infra"   { docker-compose up -d postgres redis rabbitmq elasticsearch }
    default   { Write-Host "Usage: .\manage.ps1 [swagger|run|up|down|infra]" }
}