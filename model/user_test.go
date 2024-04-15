package model

import (
	"fmt"
	"testing"
	"time"
)

func TestUser_ValidateName(t *testing.T) {
	a := "test"
	fmt.Print(a)
	type fields struct {
		ID         uint
		Name       string
		Password   string
		Mailadress string
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{Name: "12345"}, false},
		{"case2", fields{Name: "1234"}, true},
		{"case3", fields{Name: "123456789012345"}, false},
		{"case4", fields{Name: "1234567890123456"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Password:   tt.fields.Password,
				Mailadress: tt.fields.Mailadress,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if err := u.ValidateName(); (err != nil) != tt.wantErr {
				t.Errorf("User.ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_ValidatePassword(t *testing.T) {
	type fields struct {
		ID         uint
		Name       string
		Password   string
		Mailadress string
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{Password: "12345"}, false},
		{"case2", fields{Password: "1234"}, true},
		{"case3", fields{Password: "123456789012345"}, false},
		{"case4", fields{Password: "1234567890123456"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Password:   tt.fields.Password,
				Mailadress: tt.fields.Mailadress,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if err := u.ValidatePassword(); (err != nil) != tt.wantErr {
				t.Errorf("User.ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_ValidateMailAdress(t *testing.T) {
	type fields struct {
		ID         uint
		Name       string
		Password   string
		Mailadress string
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		// 30 15
		{"case1", fields{Mailadress: "12345@gmail.com"}, false},
		{"case2", fields{Mailadress: "1234"}, true},
		{"case3", fields{Mailadress: "12345678901234567890@gmail.com"}, false},
		{"case4", fields{Mailadress: "123456789012345678901@gmail.com"}, true},
		{"case5", fields{Mailadress: "12345gmail.com"}, true},
		{"case5", fields{Mailadress: "12345@@gmail.com"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Password:   tt.fields.Password,
				Mailadress: tt.fields.Mailadress,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if err := u.ValidateMailAdress(); (err != nil) != tt.wantErr {
				t.Errorf("User.ValidateMailAdress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
