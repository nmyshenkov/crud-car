package db

import (
	"crud-car/cars"
	"database/sql"
	"testing"
)

func TestInsert(t *testing.T) {
	dbconn, err := InitDB(DSN)
	if err != nil {
		t.Errorf("Error connect to db %v", err)
	}
	defer dbconn.Close()

	type args struct {
		db *sql.DB
		r  cars.Car
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				db: dbconn, r: cars.Car{
					Company:  "audi",
					Model:    "a30",
					Capacity: 60,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Insert(tt.args.db, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
