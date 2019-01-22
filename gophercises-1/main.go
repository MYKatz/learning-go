package main

import (

    "flag"
    "fmt"
    "os"
    "encoding/csv"
    "strings"
)

func main(){
    csvFilename := flag.String("csv", "problems.csv", "a csv file. format: 'question','answer'")
    flag.Parse()
    
    file, err := os.Open(*csvFilename)
    
    if err != nil {
        fmt.Printf("Can't open the csv file: %s\n",*csvFilename)
        os.Exit(1)
    }
    
    r := csv.NewReader(file)
    lines, err := r.ReadAll()
    if err != nil {
        fmt.Printf("Can't use csv file... failed to parse")
        os.Exit(1)
    }
    
    problems := parseLines(lines)
    
    amountCorrect := 0
    for i, p := range problems {
        fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
        var answer string
        fmt.Scanf("%s\n", &answer)
        if answer == p.answer {
            amountCorrect++
        }
    }
    fmt.Printf("You got %d right / %d questions", amountCorrect, len(problems))
}

func parseLines(lines [][]string) []problem {
    toret := make([]problem, len(lines))
    for i, line := range lines {
        toret[i] = problem{
            question: line[0],
            answer: strings.TrimSpace(line[1]),
        }
    }
    return toret
}

type problem struct {
    question string
    answer string
}
