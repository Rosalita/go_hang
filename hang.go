package main
import(
  "fmt"
  "strings"
  "regexp"
)

func main(){
  again:= play_hangman()
  for {
    if again == "y"{
      again = play_hangman()
    } else if again == "n"{
      break
    }
  }
}

func play_hangman()string{
  stage_of_death := 0
  has_guessed_1_letter := false
  guess :=""
  guessed_letters :=""
  again :=""
  dashes :=""
  newdashes:=""
  word :="test" //to do make a long list of words and pick one at random.
  fmt.Printf("H A N G M A N\n")
  for {
      draw_hangman(stage_of_death)
      if stage_of_death == 6{
           fmt.Printf("Oh dear hangman is dead\n")
           for{
               fmt.Printf("Play again? (y/n) \n")
               fmt.Scanln(&again)
               isYorN, _ := regexp.MatchString("^y|Y|n|N",again)
               if isYorN == false{
                   fmt.Printf("You didn't type 'y' or 'n'! Try again\n")
               } else if (len(again) > 1){
                   fmt.Printf("You entered more than 1 character! Try again\n")
               } else if strings.ToLower(again) == "y"{
                   return "y"
               } else if strings.ToLower(again) == "n"{
                   return "n"
               }
           }
      }
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
    dashes:=""
    i:=0
    for i = 0; i < wordlen; i++{
        dashes += "_"
    }
    return dashes
}
func revealdashes(word string, guess string, dashes string) string{
    newdashes:=""
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
