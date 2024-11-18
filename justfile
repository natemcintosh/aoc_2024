# Run all tests
default: test-all

# Test all days
test-all:
    #!/home/natemcintosh/.cargo/bin/nu
    # Have to use bash here because nu incorrectly interprets `./...`
    # This syntax is a go shortcut for testing everything it can find.
    /usr/bin/bash -c "go test ./..."

alias td := test-day

# Test one day
test-day day:
    #!/home/natemcintosh/.cargo/bin/nu
    let sday = ({{ day }} | into string)
    let formatted_day = if ($sday | str length) == 1 { ['0', $sday] | str join } else { $sday }
    print $"Running tests for day ($formatted_day)"
    go test -v $"./day($formatted_day)"

alias rd := run-day

# Run a specific day
run-day day:
    #!/home/natemcintosh/.cargo/bin/nu
    let sday = ({{ day }} | into string)
    let formatted_day = if ($sday | str length) == 1 { ['0', $sday] | str join } else { $sday }
    go run $"./day($formatted_day)/main.go"
