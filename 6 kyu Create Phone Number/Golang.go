import (
  "regexp"
  "strconv"
)

func CreatePhoneNumber(numbers [10]uint) (str string) {
    regex := regexp.MustCompile(`(\d{3})(\d{3})(\d{4})`)
  
    for _, n := range numbers {
      str += strconv.FormatUint(uint64(n), 10)
    }
    
    return regex.ReplaceAllString(str, "($1) $2-$3")
}
