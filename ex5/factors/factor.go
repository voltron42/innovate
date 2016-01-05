package factors

func Factor(n int) []int {
  out := []int{}
  for x := 2; x < n; x++ {
    if n % x == 0 {
      out = append(out, x)
    }
  }
  return out
}
