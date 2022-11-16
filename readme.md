go install github.com/spf13/cobra-cli@latest
alias cobra-cli='~/go/bin/cobra-cli'
cobra-cli init cli_golang

go run app/main.go
