package main
import(
  "fmt"
  "strings"
  "regexp"
)

func main(){
  stage_of_death := 0
  has_guessed_1_letter := false
  guess :=""
  guessed_letters :=""
  var dashes string
  var newdashes string
  word :="test" //to do make a long list of words and pick one at random.
  fmt.Printf("H A N G M A N\n")
  for {
      draw_hangman(stage_of_death)
      if has_guessed_1_letter == false{
          dashes = hideword(len(word))
          fmt.Printf("%s\n",dashes)
      } else {
          fmt.Printf("%s\n", newdashes)
      }
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

          if has_guessed_1_letter == false{
              updateddashes:= revealdashes(word, guess, dashes)
              newdashes = updateddashes
          } else {
              updateddashes:= revealdashes(word,guess,newdashes)
              newdashes = updateddashes
          }
          
          has_guessed_1_letter = true
      } else {
          fmt.Printf("The letter you guessed is not in the word\n")
          stage_of_death ++
          guessed_letters += guess
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
func hideword(wordlen int) string {
    var dashes string
    var i int
    for i = 0; i < wordlen; i++{
        dashes += "_"
    }
    return dashes
}
func revealdashes(word string, guess string, dashes string) string{
    var newdashes string
    for i, r := range dashes {
        c := string(r)
        if c != "_"{
            newdashes += c

        } else {
            var letter = string(word[i])
                if guess == letter{
                    newdashes += guess
                } else {
                   newdashes += "_"
                }
        }
    }
    return newdashes
}
