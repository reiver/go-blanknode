package blanknodeid

import (
	"testing"

	"errors"
)

func TestParseLabelString(t *testing.T) {
	tests := []struct {
		Value         string
		ExpectedLabel Label
		ExpectedError error
	}{
		{
			ExpectedError: ErrEmptyString,
		},



		{
			Value:                   "_",
			ExpectedLabel: someLabel("_"),
			ExpectedError: nil,
		},
		{
			Value:                   "__",
			ExpectedLabel: someLabel("__"),
			ExpectedError: nil,
		},
		{
			Value:                   "___",
			ExpectedLabel: someLabel("___"),
			ExpectedError: nil,
		},
		{
			Value:                   "____",
			ExpectedLabel: someLabel("____"),
			ExpectedError: nil,
		},
		{
			Value:                   "_____",
			ExpectedLabel: someLabel("_____"),
			ExpectedError: nil,
		},
		{
			Value:                   "_apple",
			ExpectedLabel: someLabel("_apple"),
			ExpectedError: nil,
		},
		{
			Value:                   "_apple_",
			ExpectedLabel: someLabel("_apple_"),
			ExpectedError: nil,
		},
		{
			Value:                   "_apple_BANANA",
			ExpectedLabel: someLabel("_apple_BANANA"),
			ExpectedError: nil,
		},
		{
			Value:                   "_apple_BANANA_",
			ExpectedLabel: someLabel("_apple_BANANA_"),
			ExpectedError: nil,
		},
		{
			Value:                   "_apple_BANANA_Cherry",
			ExpectedLabel: someLabel("_apple_BANANA_Cherry"),
			ExpectedError: nil,
		},
		{
			Value:                   "_apple_BANANA_Cherry_",
			ExpectedLabel: someLabel("_apple_BANANA_Cherry_"),
			ExpectedError: nil,
		},
		{
			Value:                   "apple_BANANA_Cherry",
			ExpectedLabel: someLabel("apple_BANANA_Cherry"),
			ExpectedError: nil,
		},
		{
			Value:                   "__FUNCTION__",
			ExpectedLabel: someLabel("__FUNCTION__"),
			ExpectedError: nil,
		},
		{
			Value:                   "E_WARNING",
			ExpectedLabel: someLabel("E_WARNING"),
			ExpectedError: nil,
		},
		{
			Value:                   "__sleep",
			ExpectedLabel: someLabel("__sleep"),
			ExpectedError: nil,
		},
		{
			Value:                   "wakeup__",
			ExpectedLabel: someLabel("wakeup__"),
			ExpectedError: nil,
		},
		{
			Value:                   "_SERVER",
			ExpectedLabel: someLabel("_SERVER"),
			ExpectedError: nil,
		},
		{
			Value:                   "_abcde_12345_",
			ExpectedLabel: someLabel("_abcde_12345_"),
			ExpectedError: nil,
		},



		{
			Value:         ".",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "..",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "...",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         ".abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "abc.",
			ExpectedError: ErrLabelLastCharacterNotAllowed,
		},
		{
			Value:         ".___.",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},



		{
			Value:                   "apple.BANANA.Cherry",
			ExpectedLabel: someLabel("apple.BANANA.Cherry"),
			ExpectedError: nil,
		},
		{
			Value:                   "_._",
			ExpectedLabel: someLabel("_._"),
			ExpectedError: nil,
		},
		{
			Value:                   "_.._",
			ExpectedLabel: someLabel("_.._"),
			ExpectedError: nil,
		},
		{
			Value:                   "_..._",
			ExpectedLabel: someLabel("_..._"),
			ExpectedError: nil,
		},
		{
			Value:                   "_...._",
			ExpectedLabel: someLabel("_...._"),
			ExpectedError: nil,
		},
		{
			Value:                   "_....._",
			ExpectedLabel: someLabel("_....._"),
			ExpectedError: nil,
		},



		{
			Value:         "-",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "-abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc-",
			ExpectedLabel: someLabel("abc-"),
		},



		{
			Value:         "\u00B7",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u00B7abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u00B7",
			ExpectedLabel: someLabel("abc\u00B7"),
		},



		{
			Value:         "\u0300",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u0300abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u0300",
			ExpectedLabel: someLabel("abc\u0300"),
		},

		{
			Value:         "\u0301",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u0301abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u0301",
			ExpectedLabel: someLabel("abc\u0301"),
		},

		{
			Value:         "\u0302",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u0302abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u0302",
			ExpectedLabel: someLabel("abc\u0302"),
		},

		{
			Value:         "\u0311",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u0311abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u0311",
			ExpectedLabel: someLabel("abc\u0311"),
		},

		{
			Value:         "\u0322",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u0322abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u0322",
			ExpectedLabel: someLabel("abc\u0322"),
		},

		{
			Value:         "\u0333",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u0333abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u0333",
			ExpectedLabel: someLabel("abc\u0333"),
		},

		{
			Value:         "\u0344",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u0344abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u0344",
			ExpectedLabel: someLabel("abc\u0344"),
		},

		{
			Value:         "\u0355",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u0355abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u0355",
			ExpectedLabel: someLabel("abc\u0355"),
		},

		{
			Value:         "\u036E",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u036Eabc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u036E",
			ExpectedLabel: someLabel("abc\u036E"),
		},

		{
			Value:         "\u036F",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u036Fabc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u036F",
			ExpectedLabel: someLabel("abc\u036F"),
		},



		{
			Value:         "\u203F",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u203Fabc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u203F",
			ExpectedLabel: someLabel("abc\u203F"),
		},

		{
			Value:         "\u2040",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:         "\u2040abc",
			ExpectedError: ErrLabelFirstCharacterNotAllowed,
		},
		{
			Value:                   "abc\u2040",
			ExpectedLabel: someLabel("abc\u2040"),
		},



		{
			Value:                   "_-",
			ExpectedLabel: someLabel("_-"),
		},
		{
			Value:                   "_-abc",
			ExpectedLabel: someLabel("_-abc"),
		},

		{
			Value:                   "_\u00B7",
			ExpectedLabel: someLabel("_\u00B7"),
		},
		{
			Value:                   "_\u00B7abc",
			ExpectedLabel: someLabel("_\u00B7abc"),
		},

		{
			Value:                   "_\u0300",
			ExpectedLabel: someLabel("_\u0300"),
		},
		{
			Value:                   "_\u0300abc",
			ExpectedLabel: someLabel("_\u0300abc"),
		},

		{
			Value:                   "_\u036F",
			ExpectedLabel: someLabel("_\u036F"),
		},
		{
			Value:                   "_\u036Fabc",
			ExpectedLabel: someLabel("_\u036Fabc"),
		},

		{
			Value:                   "_\u203F",
			ExpectedLabel: someLabel("_\u203F"),
		},
		{
			Value:                   "_\u203Fabc",
			ExpectedLabel: someLabel("_\u203Fabc"),
		},

		{
			Value:                   "_\u2040",
			ExpectedLabel: someLabel("_\u2040"),
		},
		{
			Value:                   "_\u2040abc",
			ExpectedLabel: someLabel("_\u2040abc"),
		},




		{
			Value:                   "0ab75000-63e2-4216-a818-def5549e2f7d",
			ExpectedLabel: someLabel("0ab75000-63e2-4216-a818-def5549e2f7d"),
		},
		{
			Value:                   "18df01b8-9a7d-40a7-aa30-2d1275db2636",
			ExpectedLabel: someLabel("18df01b8-9a7d-40a7-aa30-2d1275db2636"),
		},
		{
			Value:                   "2a37546c-e2a7-4eba-a149-ad49b8be9794",
			ExpectedLabel: someLabel("2a37546c-e2a7-4eba-a149-ad49b8be9794"),
		},
		{
			Value:                   "35a58c9e-2755-4a3c-92c9-9794ad851f91",
			ExpectedLabel: someLabel("35a58c9e-2755-4a3c-92c9-9794ad851f91"),
		},
		{
			Value:                   "4856703a-8045-4760-a85b-5e25d1b10753",
			ExpectedLabel: someLabel("4856703a-8045-4760-a85b-5e25d1b10753"),
		},
		{
			Value:                   "5f9330a4-addc-406c-8f1f-412c58bb3cd8",
			ExpectedLabel: someLabel("5f9330a4-addc-406c-8f1f-412c58bb3cd8"),
		},
		{
			Value:                   "6f8b32d5-a3f8-409d-95f1-eb3a52e8d322",
			ExpectedLabel: someLabel("6f8b32d5-a3f8-409d-95f1-eb3a52e8d322"),
		},
		{
			Value:                   "757c6fca-7ef3-4f62-aef7-d08a76f4af51",
			ExpectedLabel: someLabel("757c6fca-7ef3-4f62-aef7-d08a76f4af51"),
		},
		{
			Value:                   "8137136e-a53c-40b3-826a-5d68b302e724",
			ExpectedLabel: someLabel("8137136e-a53c-40b3-826a-5d68b302e724"),
		},
		{
			Value:                   "90954414-08e6-4804-8eea-38ad98de0df8",
			ExpectedLabel: someLabel("90954414-08e6-4804-8eea-38ad98de0df8"),
		},
		{
			Value:                   "af49569a-6b9e-4ba0-9e9a-cb47c99d0338",
			ExpectedLabel: someLabel("af49569a-6b9e-4ba0-9e9a-cb47c99d0338"),
		},
		{
			Value:                   "b62b1754-ed63-4e74-a995-679b0d78e278",
			ExpectedLabel: someLabel("b62b1754-ed63-4e74-a995-679b0d78e278"),
		},
		{
			Value:                   "c0f8c0fc-dd57-443b-9c97-c573ea1ef937",
			ExpectedLabel: someLabel("c0f8c0fc-dd57-443b-9c97-c573ea1ef937"),
		},
		{
			Value:                   "d4e77931-9435-4f1a-b7c6-a4ead8a29b4b",
			ExpectedLabel: someLabel("d4e77931-9435-4f1a-b7c6-a4ead8a29b4b"),
		},
		{
			Value:                   "e28cfeee-3a18-4bb5-b01e-6009114a8c0f",
			ExpectedLabel: someLabel("e28cfeee-3a18-4bb5-b01e-6009114a8c0f"),
		},
		{
			Value:                   "f8dab989-fe86-45ba-b5db-ed4cb1f6b204",
			ExpectedLabel: someLabel("f8dab989-fe86-45ba-b5db-ed4cb1f6b204"),
		},
	}

	for testNumber, test := range tests {
		actualLabel, actualError := ParseLabelString(test.Value)
		if nil == test.ExpectedError && nil != actualError {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Logf("ERROR: %s", actualError)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
		if nil != test.ExpectedError && !errors.Is(actualError, test.ExpectedError) {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED-ERROR: %s", test.ExpectedError)
			t.Logf("ACTUAL-ERROR:   %s", actualError)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		{
			expected := test.ExpectedLabel
			actual := actualLabel

			if expected != actual {
				t.Errorf("For test #%d, the actual blank-node-label is not what was expected.", testNumber)
				t.Logf("EXPECTED: (%d) %q", len(expected.String()), expected)
				t.Logf("ACTUAL:   (%d) %q", len(actual.String()), actual)
				t.Logf("VALUE:    %q", test.Value)
				continue
			}
		}
	}
}
