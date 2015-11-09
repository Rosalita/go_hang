package main
import(
  "fmt"
  "strings"
  "regexp"
)

func main(){
  stage_of_death := 0
  guess :=""
  guessed_letters :=""
  word :="test" //to do make a long list of words and pick one at random.
  fmt.Printf("H A N G M A N\n")
  for {
      draw_hangman(stage_of_death)
      draw_dashes(len(word))
      fmt.Printf("Guess a letter: ")
      fmt.Scanln(&guess)
      isALetter, _ := regexp.MatchString("^[a-zA-Z]",guess)
      if isALetter == false{
        fmt.Printf("That's not a letter! Try again\n")
      } else if (len(guess) > 1){
        fmt.Printf("You entered more than 1 character! Try again\n")
      } else if strings.Contains(guessed_letters, guess){
          fmt.Printf("You have already guessed that letter! Try again\n")
      } else if strings.Contains(word, guess){
          fmt.Printf("The letter you guessed is in the word\n")
          guessed_letters += guess
          fmt.Printf(guessed_letters)
      } else {
          fmt.Printf("The letter you guessed is not in the word\n")
          stage_of_death ++
          guessed_letters += guess
          fmt.Printf(guessed_letters)
      }
  }
}
func draw_hangman(stage_of_death int){
    switch stage_of_death{
        case 0:
            fmt.Printf(" +---+\n")
            fmt.Printf(" |   |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("========\n")

        case 1:
            fmt.Printf(" +---+\n")
            fmt.Printf(" |   |\n")
            fmt.Printf(" O   |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("========\n")

        case 2:
            fmt.Printf(" +---+\n")
            fmt.Printf(" |   |\n")
            fmt.Printf(" O   |\n")
            fmt.Printf(" |   |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("========\n")

        case 3:
            fmt.Printf(" +---+\n")
            fmt.Printf(" |   |\n")
            fmt.Printf(" O   |\n")
            fmt.Printf("/|   |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("========\n")

        case 4:
            fmt.Printf(" +---+\n")
            fmt.Printf(" |   |\n")
            fmt.Printf(" O   |\n")
            fmt.Printf("/|\\  |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("========\n")

        case 5:
            fmt.Printf(" +---+\n")
            fmt.Printf(" |   |\n")
            fmt.Printf(" O   |\n")
            fmt.Printf("/|\\  |\n")
            fmt.Printf("/    |\n")
            fmt.Printf("     |\n")
            fmt.Printf("     |\n")
            fmt.Printf("========\n")

        case 6:
            fmt.Printf(" +---+\n")
            fmt.Printf(" |   |\n")
            fmt.Printf(" O   |\n")
            fmt.Printf("/|\\  |\n")
            fmt.Printf("/ \\  |\n")
            fmt.Printf("     |\n")
            fmt.Printf("R.I.P|\n")
            fmt.Printf("========\n")
    }
}
func draw_dashes(number_of_dashes int){
    var i int
    for i = 0; i < number_of_dashes; i++{
        fmt.Printf("_")
    }
    fmt.Printf("\n")
}
