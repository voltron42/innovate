package factors_test

import (
  "."
  "testing"
)

func Test7(t *testing.T) {
  list := factors.Factor(7)
  if len(list) != 0 {
    t.Fatal("7 should have no factors")
  }
}
