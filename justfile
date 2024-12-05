# Run all tests
default: test-all

# Test all days
test-all:
    go test "./..."

alias td := test-day

# Test one day
test-day day:
    #!{{ home_directory() }}/.cargo/bin/nu
    let sday = ({{ day }} | into string)
    let formatted_day = if ($sday | str length) == 1 { ['0', $sday] | str join } else { $sday }
    print $"Running tests for day ($formatted_day)"
    go test -v $"./day($formatted_day)"

alias rd := run-day

# Run a specific day
run-day day:
    #!{{ home_directory() }}/.cargo/bin/nu
    let sday = ({{ day }} | into string)
    let formatted_day = if ($sday | str length) == 1 { ['0', $sday] | str join } else { $sday }
    go run $"./day($formatted_day)/main.go"

# Create the structure for a new day
new-day day:
    #!{{ home_directory() }}/.cargo/bin/nu
    mkdir $"day({{ day }})"
    touch $"day({{ day }})/input.txt"
    echo "package main" | save $"day({{ day }})/main.go"
    echo "package main" | save $"day({{ day }})/main_test.go"
