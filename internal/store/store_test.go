package store

import (
	"testing"

	"service/internal/user"
)

func Test_convertFriends(t *testing.T) {
	tests := []struct {
		name    string
		user    user.User
		wantRes string
	}{
		{
			"no friends",
			user.User{},
			"",
		},
		{
			"one friend",
			user.User{
				Friends: []int{1},
			},
			"1",
		},
		{
			"two friends",
			user.User{
				Friends: []int{1, 2},
			},
			"1,2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := convertFriends(tt.user); gotRes != tt.wantRes {
				t.Errorf("convertFriends() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
