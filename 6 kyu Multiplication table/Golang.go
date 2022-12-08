func MultiplicationTable(size int) (res [][]int) {
    if size == 0 {
      return res
    }
  
    if size == 1 {
      return [][]int{{1}}
    }
  
    for i := 1; i <= size; i ++ {
        x := []int{}
      
        for j := 1; j <= size; j ++ {
            x = append(x, i * j)
        }
      
        res = append(res, x)
    }
  
    return res
}
