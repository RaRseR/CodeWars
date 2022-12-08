func BreakChocolate(n, m int) int {
  if n * m <= 1 {
    return 0
  }
  
  return n * m - 1
}
