package word2number

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFloat(t *testing.T) {
	c, err := NewConverter("en")
	if err != nil {
		t.Fatal(err)
	}
	words := "one and a half"
	t.Run(fmt.Sprint("testcase Float"), func(t *testing.T) {
		if got := c.Words2Number(words); got != 1.5 {
			t.Errorf("Converter.Words2Number(%s) = %v, want %v", words, got, 1.5)
		}
	})
}

func TestFloat2(t *testing.T) {
	c, err := NewConverter("en")
	if err != nil {
		t.Fatal(err)
	}
	words := "one point 5"
	t.Run(fmt.Sprint("testcase Float"), func(t *testing.T) {
		if got := c.Words2Number(words); got != 1.5 {
			t.Errorf("Converter.Words2Number(%s) = %v, want %v", words, got, 1.5)
		}
	})
}

func TestWordsToNumberArray(t *testing.T) {
	c, err := NewConverter("en")
	if err != nil {
		t.Fatal(err)
	}
	words := "take one glass and pour one and a half cups of sugar in it"
	shouldGet := []float64{1, 1.5}
	t.Run(fmt.Sprint("testcase String Convert"), func(t *testing.T) {
		if got := c.Words2NumberArray(words); !reflect.DeepEqual(got, shouldGet) {
			t.Errorf("Converter.Words2NumberArray(%s) = %v, want %v", words, got, shouldGet)
		}
	})
}

func TestConvertNumbersInFullString3(t *testing.T) {
	c, err := NewConverter("en")
	if err != nil {
		t.Fatal(err)
	}
	words := "2 and 1/2 ounce of something and 1/4 ounce of something else"
	shouldGet := "2.5 ounce of something and 0.25 ounce of something else"
	t.Run(fmt.Sprint("testcase String Convert"), func(t *testing.T) {
		if got := c.ReplaceNumbersInWordForm(words); got != shouldGet {
			t.Errorf("Converter.ReplaceNumbersInWordForm(%s) = %v, want %v", words, got, shouldGet)
		}
	})
}

func TestConvertNumbersInFullString(t *testing.T) {
	c, err := NewConverter("en")
	if err != nil {
		t.Fatal(err)
	}
	words := "take one glass and pour one and a half cups of sugar in it"
	shouldGet := "take 1 glass and pour 1.5 cups of sugar in it"
	t.Run(fmt.Sprint("testcase String Convert"), func(t *testing.T) {
		if got := c.ReplaceNumbersInWordForm(words); got != shouldGet {
			t.Errorf("Converter.ReplaceNumbersInWordForm(%s) = %v, want %v", words, got, shouldGet)
		}
	})
}

func TestConvertNumbersInFullString205(t *testing.T) {
	c, err := NewConverter("en")
	if err != nil {
		t.Fatal(err)
	}
	words := "two hundred and five"
	shouldGet := "205"
	t.Run(fmt.Sprint("testcase String Convert"), func(t *testing.T) {
		if got := c.ReplaceNumbersInWordForm(words); got != shouldGet {
			t.Errorf("Converter.ReplaceNumbersInWordForm(%s) = %v, want %v", words, got, shouldGet)
		}
	})
}

func TestConvertNumbersInFullStringComplex(t *testing.T) {
	c, err := NewConverter("en")
	if err != nil {
		t.Fatal(err)
	}
	words := "with two hundred and five bottles you and I can make twenty two cakes"
	shouldGet := "with 205 bottles you and I can make 22 cakes"
	t.Run(fmt.Sprint("testcase String Convert"), func(t *testing.T) {
		if got := c.ReplaceNumbersInWordForm(words); got != shouldGet {
			t.Errorf("Converter.ReplaceNumbersInWordForm(%s) = %v, want %v", words, got, shouldGet)
		}
	})
}

func TestConvertNumbersInFullString2(t *testing.T) {
	c, err := NewConverter("en")
	if err != nil {
		t.Fatal(err)
	}
	words := "Let cool, then pour into a glass jar and seal tightly with a lid."
	shouldGet := words
	t.Run(fmt.Sprint("testcase String Convert"), func(t *testing.T) {
		if got := c.ReplaceNumbersInWordForm(words); got != shouldGet {
			t.Errorf("Converter.ReplaceNumbersInWordForm(%s) = %v, want %v", words, got, shouldGet)
		}
	})
}

func TestConvert_sv(t *testing.T) {
	c, err := NewConverter("sv")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	tests := []struct {
		words string
		want  float64
	}{
		// Simple
		{"en", 1},
		{"ett", 1},
		{"två", 2},
		{"tre", 3},
		{"fyra", 4},
		{"fem", 5},
		{"sex", 6},
		{"etthundrafemtio", 150},
		{"tusen kronor och femtio öre", 1000.50},
		{"hundrafemtio procent", 1.5},
		{"hundrafemtio promille", 0.15},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("testcase-", i), func(t *testing.T) {
			if got := c.Words2Number(tt.words); got != tt.want {
				t.Errorf("Converter.Words2Number(%s) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}

func TestConverter_Words2Number(t *testing.T) {
	c, _ := NewConverter("en")
	tests := []struct {
		words string
		want  float64
	}{
		// Simple
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
		{"ten", 10},
		{"eleven", 11},
		{"twelve", 12},
		{"thirteen", 13},
		{"fourteen", 14},
		{"fifteen", 15},
		{"sixteen", 16},
		{"seventeen", 17},
		{"eighteen", 18},
		{"nineteen", 19},
		{"twenty", 20},
		{"thirty", 30},
		{"forty", 40},
		{"fifty", 50},
		{"sixty", 60},
		{"seventy", 70},
		{"eighty", 80},
		{"ninety", 90},
		{"niNeTy", 90},

		// Simple multi-word
		{"one hundred", 100},
		{"one-hundred", 100},
		{"two-hundred", 200},
		{"two thousand", 2000},
		{"two-thousand", 2000},

		// More complicated multiword
		{"two thousand three-hundred seventy five", 2375},
		{"two hundred thousand five", 200005},
		{"twenty-five thousand", 25000},
		{"two thousand three hundred seventy five", 2375},
		{"two - thousand three hundred seventy five", 2375},
		{"one million", 1000000},
		{"1 million", 1000000},
		{"1.2 million", 1200000},
		{"Forty-Eight Million, Four Hundred Thousand", 48400000},
		{"two hundred fifty thousand", 250000},
		{"two hundred and fifty thousand", 250000},
		{"two thousand and fifty million", 2050000000},
		{"one million, three hundred nine thousand", 1309000},

		// Decimals
		{"oh point twenty-five", 0.25},
		{"zero point five thousandths", 0.005},
		{"ten point fifty-five hundredths", 10.55},
		{"fifty-five hundredths", 0.55}, // decimal portion not preceded by "point" or "and"
		{"one and fifty five hundredths", 1.55},
		{"one and seven tenths", 1.7},
		{"one and seven hundredths", 1.07},
		{"one and seven thousandths", 1.007},
		{"one point seventy-seven", 1.77},
		// {"one point seven seven", 1.77}, // Doesn't work properly. treated like 1. 7+7
		{"one and seventy-seven hundredths", 1.77},
		{"one and seventy seven thousandths", 1.077},
		{"one and seventy seven hundred thousandths", 1.00077},
		{"seven-hundred-seventy-seven", 777},
		{"seven-hundred-seventy-seven", 777},
		{"fifty cents", 0.5},
		// {"one and seven-hundred-seventy-seven-thousandths", 1.777}, // Rounding error. Strange
		{"zero and seven hundredths", 0.07},
		// {"one and seven-hundred-seventy-seven ten-thousandths", 1.0777}, // ten-thousandths doesn't work. "ten" is only a multiplier to the right of the decimal
		{"one and seven-hundred-seventy-seven hundred thousandths", 1.00777},

		// Stupid versions
		{"hundred thousand", 100000},
		{"three hundred and twelve US dollars and fifty cents", 312.50},
		{"seventyfive", 75},

		// percent and cent
		{"one hundred percent", 1.00},
		{"hundred percent", 1.00},
		{"two point five percent", 0.025},
		{"ninety nine percent", .99},
		{"seventy percent", 0.7},
		{"seventy cents", .7},
		{"two hundred fifty percent", 2.50},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("testcase-", i), func(t *testing.T) {
			if got := c.Words2Number(tt.words); got != tt.want {
				t.Errorf("Converter.Words2Number(%s) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}

func TestConverter_Number2Words(t *testing.T) {
	c, _ := NewConverter("en")
	tests := []struct {
		want1  string
		want2  string
		number float64
	}{
		// Simple
		{"one", "", 1},
		{"two", "", 2},
		{"three", "", 3},
		{"one thousand", "", 1000},
		{"two million", "", 2000000},
		{"twenty two", "", 22},
		{"one hundred ten", "", 110},
		{"two thousand five hundred", "", 2500},
		{"two million five hundred fifty five thousand three hundred sixty seven", "", 2555367},

		{"", "fifty", 0.50},
		{"eighteen", "seventy three", 18.73},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("testcase-", i), func(t *testing.T) {
			if got1, got2 := c.Number2Words(tt.number, 2); got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("Converter.Words2Number(%.2f) = (%s, %s), want %s, %s", tt.number, got1, got2, tt.want1, tt.want2)
			}
		})
	}
}
