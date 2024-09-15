package main

import "fmt"
import st "strings"

var p = fmt.Println

func main() {
	p("Contains:", st.Contains("test", "es"))
	p("Count:", st.Count("test", "t"))
	p("HasPrefix:", st.HasPrefix("test", "te"))
	p("HasSuffix", st.HasSuffix("test", "st"))
	p("Index:", st.Index("test", "e"))
	p("Join:", st.Join([]string{"a", "b"}, "-"))
	p("Replace:   ", st.Replace("foo", "o", "0", -1))
	p("Replace:   ", st.Replace("foo", "o", "0", 1))
	p("Split:", st.Split("a-b-c", "-"))
	p("ToLower", st.ToLower("TEST"))
	p("ToUpper", st.ToUpper("hello"))
}
