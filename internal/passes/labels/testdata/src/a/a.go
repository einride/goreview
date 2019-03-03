package a

func A() {
GoodLabel:
	for i := 0; i < 100; i++ {
		if i == 54 {
			break GoodLabel
		}
	}
badLabel: // want "labels must use CamelCase"
	for i := 0; i < 100; i++ {
		if i == 54 {
			break badLabel
		}
	}
Bad_label: // want "labels must use CamelCase"
	for i := 0; i < 100; i++ {
		if i == 54 {
			break Bad_label
		}
	}
}
