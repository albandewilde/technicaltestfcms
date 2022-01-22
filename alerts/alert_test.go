package alerts

import "testing"

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
		name     string
		expected Alert
		fields   fields
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
	}

	for _, test := range cases {
		got := NewAlert(
			test.fields.ID,
			test.fields.UserID,
			test.fields.Date,
			test.fields.Name,
			test.fields.MaxPrice,
			test.fields.MinPrice,
			test.fields.City,
		)

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
