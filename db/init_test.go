package db

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Success", args: args{dsn: DSN}, wantErr: false},
		{name: "Bad", args: args{dsn: "DSN"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, err := InitDB(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if conn != nil {
				_ = conn.Close()
			}
		})
	}
}
