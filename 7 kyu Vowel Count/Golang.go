func GetCount(str string) (count int) {
  vowels := map[rune]bool{'a' : true, 'e' : true, 'i' : true, 'o' : true, 'u' : true}
  
  for _, letter := range str {
    if _, ok := vowels[letter]; ok {
      count++
    }
  }
  
  return count
}
