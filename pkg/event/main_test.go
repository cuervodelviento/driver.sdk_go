package event

import "testing"

func TestGetUserTypeAndUserIdFromPersonId(t *testing.T) {
	test_table := []struct {
		personId string
		userId   int
		userType int
	}{
		{"nu1", 1, 1},
		{"nu12", 12, 1},
		{"nu123", 123, 1},
		{"nu1234", 1234, 1},
		{"nu12345", 12345, 1},
		{"nu123456", 123456, 1},
		{"nu1234567", 1234567, 1},
		{"vi1", 1, 3},
		{"vi12", 12, 3},
		{"vi123", 123, 3},
		{"vi1234", 1234, 3},
		{"vi12345", 12345, 3},
		{"vi123456", 123456, 3},
		{"vi1234567", 1234567, 3},
		{"em1", 1, 2},
		{"em12", 12, 2},
		{"em123", 123, 2},
		{"em1234", 1234, 2},
		{"em12345", 12345, 2},
	}

	for _, test := range test_table {
		userType, userId, err := getUserTypeAndUserIdFromPersonID(test.personId)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if userId != test.userId {
			t.Errorf("Expected userId: %d, got: %d", test.userId, userId)
		}
		if userType != test.userType {
			t.Errorf("Expected userType: %d, got: %d", test.userType, userType)
		}
	}

	_, _, err := getUserTypeAndUserIdFromPersonID("2")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	_, _, err = getUserTypeAndUserIdFromPersonID("nu")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

}
