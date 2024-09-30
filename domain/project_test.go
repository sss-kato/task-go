package domain

import (
	"reflect"
	"testing"
)

func TestNewProject(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProject(tt.args.pnm, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProject() = %v, want %v", got, tt.want)
			}
		})
	}
}
