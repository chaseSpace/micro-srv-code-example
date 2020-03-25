package pkg

import "time"

/*
app共享的代码
*/

func GetTextByTime(t ...time.Time) []byte {
	var tmp time.Time
	if len(t) == 0 {
		tmp = time.Now()
	} else {
		tmp = t[0]
	}
	text, _ := tmp.MarshalText()
	return text
}
