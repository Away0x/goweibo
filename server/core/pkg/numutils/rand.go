package numutils

import r "math/rand"

func RandInt(min, max int) int {
  if min >= max || min < 0 || max == 0 {
    return max
  }
  return r.Intn(max-min) + min
}
