package base

type T string

func (t T)Multiple(n int) string {
	tmp:=t
	for i := 0; i < n-1; i++ {
		t+=tmp
	}
	return string(t)
}
