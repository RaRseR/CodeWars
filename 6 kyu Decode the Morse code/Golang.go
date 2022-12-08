import (
  "strings"
)

func DecodeMorse(morseCode string) (result string) {
  
  if len(morseCode) == 0 {
    return ""
  }
  
  for _, word := range strings.Split(strings.TrimSpace(morseCode), "   ") {
    for _, letter := range strings.Split(word, " ") {
      result += MORSE_CODE[letter]
    }
    result += " "
    
  }
  
  return result[:len(result) - 1]
}
