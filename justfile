set quiet

run day:
    go run ./day$(printf "%02d" {{day}})

test day:
    go run ./day$(printf "%02d" {{day}}) -test

new day:
    #!/usr/bin/env bash
    set -euo pipefail

    DAY="$(printf "%02d" {{day}})"
    NAME="day$DAY"
    mkdir -p "$NAME"
    cd "$NAME"
    curl -sS "https://adventofcode.com/2024/day/{{day}}/input" \
        --cookie "session=$(cat ../.session)" > \
        "input.txt"
    touch "test_input.txt"
    go mod init "github.com/paolostyle/advent-of-code-2024/$NAME"
    cd ..
    go work edit -use "./$NAME"

    cat > "$NAME/$NAME.go" << EOF
    package main

    import (
        "fmt"
        "time"

        "github.com/paolostyle/advent-of-code-2024/common"
    )

    func part1(input string) int {
        return 0
    }

    func part2(input string) int {
        return 0
    }

    func main() {
        defer common.TimeTrack(time.Now())
        input := common.ReadInput({{day}})
        fmt.Println("DAY $DAY")
        fmt.Println("Part 1: ", part1(input))
        fmt.Println("Part 2: ", part2(input))
    }
    EOF
