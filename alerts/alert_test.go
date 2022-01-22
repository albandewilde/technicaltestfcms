package alerts

import (
	"testing"
)

func TestNewAlert(t *testing.T) {
	type fields struct {
		ID       string
		UserID   string
		Date     string
		Name     string
		MaxPrice int
		MinPrice int
		City     string
	}
	cases := []struct {
		name        string
		expected    Alert
		expectedErr string
		fields      fields
	}{
		{
			name: "Create a simple new alrte",
			expected: Alert{
				"id",
				"userId",
				"1970-01-01",
				"an alert",
				0,
				10,
				"Nantes",
			},
			expectedErr: "",
			fields: fields{
				"id",
				"userId",
				"1970-01-01",
				"an alert",
				0,
				10,
				"Nantes",
			},
		},
		{
			name:        "Invalid month",
			expected:    Alert{},
			expectedErr: "parsing time \"1970-14-24\": month out of range",
			fields: fields{
				"id",
				"userId",
				"1970-14-24",
				"an alert",
				0,
				10,
				"Nantes",
			},
		},
		{
			name:        "Invalid day",
			expected:    Alert{},
			expectedErr: "parsing time \"1970-10-54\": day out of range",
			fields: fields{
				"id",
				"userId",
				"1970-10-54",
				"an alert",
				0,
				10,
				"Nantes",
			},
		},
		{
			name:        "Empty string",
			expected:    Alert{},
			expectedErr: "parsing time \"\" as \"2006-01-02\": cannot parse \"\" as \"2006\"",
			fields: fields{
				"id",
				"userId",
				"",
				"an alert",
				0,
				10,
				"Nantes",
			},
		},
	}

	for _, test := range cases {
		got, err := NewAlert(
			test.fields.ID,
			test.fields.UserID,
			test.fields.Date,
			test.fields.Name,
			test.fields.MaxPrice,
			test.fields.MinPrice,
			test.fields.City,
		)

		if err != nil && err.Error() != test.expectedErr {
			t.Errorf(
				"Test: %s failed, expected error %s, but got %s",
				test.name,
				test.expectedErr,
				err.Error(),
			)
		}

		if test.expected != got {
			t.Errorf(
				"Test: %s failed, expected %+v, but got %+v",
				test.name,
				test.expected,
				got,
			)
		}
	}
}
