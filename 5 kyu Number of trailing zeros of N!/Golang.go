func Zeros(n int) (result int) {
  for n > 0 {
		n /= 5
    result += n
	}
	return
}
