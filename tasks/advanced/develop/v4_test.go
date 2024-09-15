package main

import "testing"

func TestUnpackString(t *testing.T) {
	myDict := NewDictionary()
	myDict.AddWords([]string{"АМКАР", "КАРМА", "КРАМА", "МАКАР", "МАКРА", "МАРКА", "РАМКА",
		"ПЯТАК", "ПЯТКА", "ТЯПКА", "КОСАЧ", "САЧОК", "ЧАСОК", "АВТОР", "ВАРТО", "ВТОРА", "ОТВАР",
		"РВОТА", "ТАВРО", "ТОВАР", "КАЧУР", "КРАУЧ", "КРУЧА", "КУРЧА", "РУЧКА", "ЧУРКА", "АБНЯ",
		"БАНЯ", "БАЯН", "КОРТ", "КРОТ", "ТРОК", "КОТ", "КТО", "ОТК", "ТОК",
	})

	testTable := []struct {
		name string
		in   []string
		out  map[string][]string
	}{
		{name: "one key word", in: []string{"makar"}, out: map[string][]string{"makar": {"ww", "ww", "ww"}}},
		{name: "serv", in: []string{"cat", "makar", "sacjol"}, out: map[string][]string{
			"cat":   {"cat", "who"},
			"makar": {"wq", "wdq"},
			"sachk": {"wqqw", "wqw"},
		}},
		{name: "qwdqwd", in: []string{"cat", "wdq"}, out: map[string][]string{
			"cq": {"wqdqw"},
		}},
	}

	for _, testingCase := range testTable {
		t.Run(testingCase.name, f)
	}
}
