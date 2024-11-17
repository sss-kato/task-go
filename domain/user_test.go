package domain

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	nm_test1 := "test_name1"
	pw_test1 := "test_pass1"
	ma_test1 := "test_mail_1@gmail.com"
	test_user1 := makeTestUser(nm_test1, pw_test1, ma_test1)

	nm_test2 := "1234567890123456"
	pw_test2 := "test_pass2"
	ma_test2 := "test_mail_2@gmail.com"
	test_user2 := makeTestUser(nm_test2, pw_test2, ma_test2)

	nm_test3 := "1234"
	pw_test3 := "test_pass2"
	ma_test3 := "test_mail_3@gmail.com"
	test_user3 := makeTestUser(nm_test3, pw_test3, ma_test3)

	nm_test4 := "test_name4"
	pw_test4 := "1234567890123456"
	ma_test4 := "test_mail_3@gmail.com"
	test_user4 := makeTestUser(nm_test4, pw_test4, ma_test4)

	nm_test5 := "test_name5"
	pw_test5 := "1234"
	ma_test5 := "test_mail_5@gmail.com"
	test_user5 := makeTestUser(nm_test5, pw_test5, ma_test5)

	type args struct {
		nm string
		pw string
		ma string
	}
	tests := []struct {
		name    string
		args    args
		want    UserIF
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", args{nm_test1, pw_test1, ma_test1}, test_user1, false},
		{"case2", args{nm_test2, pw_test2, ma_test2}, test_user2, true},
		{"case3", args{nm_test3, pw_test3, ma_test3}, test_user3, true},
		{"case4", args{nm_test4, pw_test4, ma_test4}, test_user4, true},
		{"case5", args{nm_test5, pw_test5, ma_test5}, test_user5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.nm, tt.args.pw, tt.args.ma)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.name == "case1" {
				checkUserTestCase1(got, tt.want, t)
			}

			if tt.name == "case2" {
				checkUserTestCase2(err, t)
			}

			if tt.name == "case3" {
				checkUserTestCase3(err, t)
			}

			if tt.name == "case4" {
				checkUserTestCase4(err, t)
			}

			if tt.name == "case5" {
				checkUserTestCase5(err, t)
			}

		})
	}
}

func checkUserTestCase1(got UserIF, want UserIF, t *testing.T) {

	gottype := reflect.TypeOf(got)

	if !gottype.Implements(reflect.TypeOf((*UserIF)(nil)).Elem()) {
		t.Errorf("%v does not implement ProjectIF", got)
		return
	}

	if got.GetMailAdress() != want.GetMailAdress() {
		t.Errorf("NewProject() got = %v, want %v", got.GetMailAdress(), want.GetMailAdress())
		return
	}

	if got.GetPassWord() != want.GetPassWord() {
		t.Errorf("NewProject() got = %v, want %v", got.GetPassWord(), want.GetPassWord())
		return
	}

	if got.GetName() != want.GetName() {
		t.Errorf("NewProject() got = %v, want %v", got.GetName(), want.GetName())
		return
	}

}

func checkUserTestCase2(err error, t *testing.T) {
	emsg := err.Error()

	if emsg != ErrorMsg01 {
		t.Errorf("NewProject() got = %v, want %v", emsg, ErrorMsg01)
		return
	}

}

func checkUserTestCase3(err error, t *testing.T) {
	emsg := err.Error()

	if emsg != ErrorMsg02 {
		t.Errorf("NewProject() got = %v, want %v", emsg, ErrorMsg02)
		return
	}

}

func checkUserTestCase4(err error, t *testing.T) {
	emsg := err.Error()

	if emsg != ErrorMsg03 {
		t.Errorf("NewProject() got = %v, want %v", emsg, ErrorMsg03)
		return
	}

}

func checkUserTestCase5(err error, t *testing.T) {
	emsg := err.Error()

	if emsg != ErrorMsg04 {
		t.Errorf("NewProject() got = %v, want %v", emsg, ErrorMsg04)
		return
	}

}

type test_user struct {
	name       string
	password   string
	mailadress string
}

func makeTestUser(nm string, pw string, ma string) UserIF {
	hashedPw := hashedPassword(pw)
	return &test_user{nm, hashedPw, ma}
}

func (tu *test_user) GetMailAdress() string {

	return tu.mailadress
}

func (tu *test_user) GetPassWord() string {

	return tu.password
}

func (tu *test_user) GetName() string {

	return tu.name
}
