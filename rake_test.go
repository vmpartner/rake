package rake

import (
	"testing"
)

func TestRegexStopWords(t *testing.T) {
	re := RegexStopWords(StopWordsSlice)
	if !re.MatchString("самому") {
		t.Fatal()
	}
}

func TestRunRake(t *testing.T) {
	list := RunRake("Её называли самой красивой актрисой СССР, ею восхищались миллионы мужчин. Но она всю жизнь таила на них обиды и считала, что они приносили ей только несчастья. За что Элину Быстрицкую пытался убить единственный супруг? И почему она была одинока?")
	if len(list) <= 0 {
		t.Fatal()
	}
	if list[3].Key != "убить" && list[3].Value != 3 {
		t.Fatal()
	}
}
