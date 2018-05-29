package stringutil

import "testing"

func TestReverse(t *testing.T){
  cases := []struct {
    in, want string
  }{
      {"Hey", "yeH"},
      {"", ""},
      {"What??", "??tahW"},
  }
  for _, c := range cases {
    got := Reverse(c.in)
    if c.want != got {
      t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
    }
  }
}
