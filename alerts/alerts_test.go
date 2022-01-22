package alerts

import (
	"testing"
)

var alerts = Alerts{
	usersAlerts: map[string][]Alert{
		"bob": {
			Alert{"a", "bob", "0000-00-00", "alrt1", 0, 10, "Nantes"},
			Alert{"b", "bob", "0000-00-00", "alrt2", 1, 50, "Orléans"},
		},
		"alice": {
			Alert{"c", "alice", "0000-00-00", "my_alerte", 100, 1000, "Paris"},
			Alert{"d", "alice", "0000-00-00", "my second alerte", 5, 5, "Rennes"},
			Alert{"e", "alice", "0000-00-00", "my third", 4, 9, "Lyon"},
		},
	},
}

func TestListAlerts(t *testing.T) {
	cases := []struct {
		name        string
		expected    []Alert
		expectedErr string
		uname       string
	}{
		{
			name: "Get user list",
			expected: []Alert{
				{"a", "bob", "0000-00-00", "alrt1", 0, 10, "Nantes"},
				{"b", "bob", "0000-00-00", "alrt2", 1, 50, "Orléans"},
			},
			expectedErr: "",
			uname:       "bob",
		},
		{
			name:        "Non existing user",
			expected:    []Alert{},
			expectedErr: "User not found",
			uname:       "rober",
		},
	}

	for _, test := range cases {

		got, err := alerts.List(test.uname)

		if err != nil && err.Error() != test.expectedErr {
			t.Errorf(
				"Test: %s failed, expected error %s, but got %s",
				test.name,
				test.expectedErr,
				err.Error(),
			)
		}

		if !equals(test.expected, got) {
			t.Errorf(
				"Test: %s failed, expected %+v, but got %+v",
				test.name,
				test.expected,
				got,
			)
		}
	}
}

func equals(a, b []Alert) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
