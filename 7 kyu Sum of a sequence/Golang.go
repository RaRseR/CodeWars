func SequenceSum(start, end, step int) (result int) {
  for i := start; i <= end; i += step {
    result += i
  }
  
  return result
}
