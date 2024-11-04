package domain

import (
	"reflect"
	"testing"
)

func TestNewProject(t *testing.T) {

	test_project_1 := makeTestProject("project_test_1", 1)

	test_project_2 := makeTestProject("0123456789012345678901234567890", 2)

	type args struct {
		pnm string
		uid int
	}
	tests := []struct {
		name    string
		args    args
		want    ProjectIF
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", args{pnm: test_project_1.GetProjectName(), uid: test_project_1.GetUserID()}, test_project_1, false},
		{"case2", args{pnm: test_project_2.GetProjectName(), uid: test_project_2.GetUserID()}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProject(tt.args.pnm, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.name == "case1" {
				gottype := reflect.TypeOf(got)
				if !gottype.Implements(reflect.TypeOf((*ProjectIF)(nil)).Elem()) {
					t.Errorf("%v does not implement ProjectIF", got)
					return
				}

				if got.GetProjectName() != test_project_1.GetProjectName() {
					t.Errorf("NewProject() got = %v, want %v", got.GetProjectName(), tt.want.GetProjectName())
					return
				}

				if got.GetUserID() != test_project_1.GetUserID() {
					t.Errorf("NewProject() got = %v, want %v", got.GetUserID(), tt.want.GetUserID())
					return
				}

			}

			if tt.name == "case2" {
				emsg := err.Error()
				if emsg != ErrorMsg11 {
					t.Errorf("Error Message got = %v, want %v", err, ErrorMsg11)
					return
				}

			}
		})
	}
}

type test_project struct {
	name string
	uid  int
}

func makeTestProject(name string, uid int) ProjectIF {

	return &test_project{name, uid}
}

func (tp *test_project) GetProjectName() string {

	return tp.name
}

func (tp *test_project) GetUserID() int {

	return tp.uid
}
