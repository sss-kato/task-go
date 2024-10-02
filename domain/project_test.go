// package domain

// import (
// 	"reflect"
// 	"testing"
// )

// func TestNewProject(t *testing.T) {

// 	test_project_1 := makeTestProject("project_test_1", 1)

// 	type args struct {
// 		pnm string
// 		uid int
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    ProjectIF
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"case1", args{pnm: test_project_1.GetProjectName(), uid: test_project_1.GetUserID()}, test_project_1, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := NewProject(tt.args.pnm, tt.args.uid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewProject() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewProject() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// type test_project struct {
// 	name string
// 	uid  int
// }

// func makeTestProject(name string, uid int) ProjectIF {

// 	return &test_project{name, uid}
// }

// func (tp *test_project) GetProjectName() string {

// 	return tp.name
// }

// func (tp *test_project) GetUserID() int {

// 	return tp.uid
// }


