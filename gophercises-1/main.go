package main

import (

    "flag"
    "fmt"
    "os"
    "encoding/csv"
    "math/rand"
    "strings"
    "time"

)

func main(){
    csvFilename := flag.String("csv", "problems.csv", "a csv file. format: 'question','answer'")
    timeLimit := flag.Int("tpq", 30, "Time limit per question, in seconds")
    randomize := flag.Bool("random", false, "Whether to randomize the questions, boolean")
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
    
    if *randomize {
        problems = shuffleSlice(problems)
    }

    timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

    amountCorrect := 0
    for i, p := range problems {
        fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
        answerChannel := make(chan string)
        go func() {
            var answer string
            fmt.Scanf("%s\n", &answer)
            answerChannel <- strings.ToLower(answer)
        }()

        select {
        case <-timer.C:
            fmt.Printf("Time's up! You scored %d out of %d.\n", amountCorrect, len(problems))
            return
        case answer := <-answerChannel:
            if answer == p.answer {
                amountCorrect++
            }
        }
    }
    fmt.Printf("You got %d right / %d questions", amountCorrect, len(problems))
}

func parseLines(lines [][]string) []problem {
    toret := make([]problem, len(lines))
    for i, line := range lines {
        toret[i] = problem{
            question: line[0],
            answer: strings.ToLower(strings.TrimSpace(line[1])),
        }
    }
    return toret
}

func shuffleSlice(problems []problem) []problem {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    toret := make([]problem, len(problems))
    permutation := r.Perm(len(problems))
    for i, valindex := range permutation {
        toret[i] = problems[valindex]
    }
    return toret
} 

type problem struct {
    question string
    answer string
}

