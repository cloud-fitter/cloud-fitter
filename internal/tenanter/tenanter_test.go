package tenanter

import (
	"testing"
)

func TestShowConfigJson(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "right", args: args{configFile: "../../config.yaml"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShowConfigJson(tt.args.configFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowConfigJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				t.Log(string(got))
			}
		})
	}
}
