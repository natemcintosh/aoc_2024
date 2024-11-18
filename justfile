# Run all tests
default: test-all

# Test all days
test-all:
    #!/home/natemcintosh/.cargo/bin/nu
    echo "Running tests..."
    # For each folder starting with "day", run the tests in that folder
    ls day*
        | where type == dir
        # Run the tests in the folder
        | each { |folder| go test $"./($folder.name)" }
        # Split into columns with these column names
        | split column --regex '\s+' status folder runtime
        # Sort by the status, then folder
        | sort-by status folder

alias td := test-day

# Test one day
test-day day:
    #!/home/natemcintosh/.cargo/bin/nu
    let sday = ({{ day }} | into string)
    let formatted_day = if ($sday | str length) == 1 { ['0', $sday] | str join } else { $sday }
    print $"Running tests for day ($formatted_day)"
    go test $"./day($formatted_day)"
