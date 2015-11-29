package main
import(
  "fmt"
  "strings"
  "regexp"
  "math/rand"
  "time"
  "io/ioutil"
)
func init(){
    rand.Seed(time.Now().UnixNano())
}

func main(){
  wins:=0
  loses:=0
  numletters := 6
  again, has_won:= play_hangman(numletters)
  for {
    if has_won == true{
      wins ++
      if numletters > 4 {
        numletters --
      }

    } else {
      loses ++
      if numletters < 6 {
        numletters ++
      }
    }
    if again == "y"{
      fmt.Printf("------------------------\n")
      fmt.Printf("    Current Score\n")
      fmt.Printf("  %d: wins, %d: loses\n",wins, loses)
      fmt.Printf("------------------------\n")
      again, has_won = play_hangman(numletters)
    } else if again == "n"{
      break
    }

  }
}

func play_hangman(numletters int)(playagain string, is_winner bool,){
  stage_of_death :=0
  gamemode:= 0
  has_guessed_1_letter :=false
  has_won :=false
  guess :=""
  guessed_letters :=""
  again :=""
  dashes :=""
  newdashes:=""
  word := random_word(numletters)
  fmt.Printf("H A N G M A N\n")
  for {
      fmt.Println("Select game mode:")
      fmt.Println("1. Only use Common words (easy)")
      fmt.Println("2. Use all words (hard)")
      fmt.Scanln(&gamemode)
      break

  }
  for {
      draw_hangman(stage_of_death)
      if stage_of_death == 6{
           fmt.Printf("Oh dear hangman is dead\n")
           fmt.Printf("The word that could have saved him was %s\n", word)
           for{
               fmt.Printf("Play again? (y/n) \n")
               fmt.Scanln(&again)
               isYorN, somekindoferror := regexp.MatchString("^y|Y|n|N",again)
                   if somekindoferror!= nil{
                     fmt.Printf("Something has gone horribly wrong. ")
                     fmt.Printf("exiting with error can not regex match %v", again)
                     return
                   }
               if isYorN == false{
                   fmt.Printf("You didn't type 'y' or 'n'! Try again\n")
               } else if (len(again) > 1){
                   fmt.Printf("You entered more than 1 character! Try again\n")
               } else if strings.ToLower(again) == "y"{
                   return "y", false
               } else {
                   return "n", false
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

      isALetter, somekindoferror := regexp.MatchString("^[a-zA-Z]",guess)
      if somekindoferror!= nil{
        fmt.Printf("Something has gone horribly wrong. ")
        fmt.Printf("exiting with error can not regex match %v", guess)
        return
      }

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
          if newdashes == word{
            has_won = true
          }
          if has_won == true{
            fmt.Printf("-= C O N G R A T U L A T I O N S =-\n")
            fmt.Printf("You won the game! The word was %s\n", word)
            for{
                fmt.Printf("Play again? (y/n) \n")
                fmt.Scanln(&again)
                isYorN, somekindoferror := regexp.MatchString("^y|Y|n|N",again)
                    if somekindoferror!= nil{
                      fmt.Printf("Something has gone horribly wrong. ")
                      fmt.Printf("exiting with error can not regex match %v", again)
                      return
                    }
                if isYorN == false{
                    fmt.Printf("You didn't type 'y' or 'n'! Try again\n")
                } else if (len(again) > 1){
                    fmt.Printf("You entered more than 1 character! Try again\n")
                } else if strings.ToLower(again) == "y"{
                    return "y", true
                } else {
                    return "n", true
                }
            }
          }
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
    for i := 0; i < wordlen; i++{
        dashes += "_"
    }
    return dashes
}
func revealdashes(word string, guess string, dashes string) string{
    newdashes:=""
    for i, r := range dashes {
        if c:= string(r); c != "_"{
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

func check_if_winner(newdashes string,word string)bool{
  if newdashes == word {
    return true
  }
  return false
}

func random_word(numletters int)string{
        var dataletters []byte
        var err error
        if numletters == 4{
            dataletters, err = ioutil.ReadFile("words/common4l.txt")
        } else if numletters == 5 {
            dataletters, err = ioutil.ReadFile("words/common5l.txt")
        } else if numletters == 6 {
          dataletters, err = ioutil.ReadFile("words/common6l.txt")
        }

        if err != nil{
            panic(err)
        }
        datastr:= string(dataletters)
        somewords:= strings.Split(datastr, " ")
        randnum:=rand.Intn(len(somewords)-1)

        chosenword:= somewords[randnum]
        return chosenword

}
