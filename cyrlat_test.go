package cyrlat

import "testing"

func TestLatin(t *testing.T) {

	someTexts := []string{"Перед написанием этой библиотеки я не нашел аналагов на Go.",
		"Мне нужна была подобная библиотека.",
		"Довелося трішки згаяти часу на її написання."}

	for _, cyr := range someTexts {
		t.Log(cyr)
		t.Log(Latin(cyr))
		t.Logf("Lower: %s\n", Latin(cyr, Lower))
		t.Logf("Upper: %s\n", Latin(cyr, Upper))
		t.Logf("UnSpace: %s\n", Latin(cyr, UnSpace))
		t.Logf("Lower, UnSpace: %s\n", Latin(cyr, Lower, UnSpace))
		t.Logf("Upper, UnSpace: %s\n", Latin(cyr, Upper, UnSpace))
		t.Logf("Upper, Lower, UnSpace: %s\n", Latin(cyr, Upper, Lower, UnSpace))
	}
}
