package direction

import "testing"

func Test_NewCommand(t *testing.T) {
	tests := []struct {
		commandString    string
		expectedType     Type
		expectedDistance int
	}{
		{"forward 5", FORWARD, 5},
		{"down 3", DOWN, 3},
	}

	for _, test := range tests {
		command, err := NewCommand(test.commandString)

		if err != nil {
			t.Fatalf("commandString is not valid, commandString: '%s', error: '%v'", test.commandString, err)
		}

		if command.Direction != test.expectedType {
			t.Fatalf("expected type: '%s', but got: '%s'", test.expectedType, command.Direction)
		}

		if command.Distance != test.expectedDistance {
			t.Fatalf("expected distance: '%d', but got: '%d'", test.expectedDistance, command.Distance)
		}
	}
}
