package db

import (
	"database/sql"
	"testing"
)

func TestDelete(t *testing.T) {
	dbconn, err := InitDB(DSN)
	if err != nil {
		t.Errorf("Error connect to db %v", err)
	}
	defer dbconn.Close()

	type args struct {
		db *sql.DB
		id int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Success", args: args{db: dbconn, id: 1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Delete(tt.args.db, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
