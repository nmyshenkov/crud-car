package db

import (
	"crud-car/cars"
	"database/sql"
	"fmt"
	"testing"
)

func TestSelectOne(t *testing.T) {
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
		want    cars.Car
		wantErr bool
	}{
		{name: "Success", args: args{db: dbconn, id: 1}, want: cars.Car{ID: 1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectOne(tt.args.db, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID != tt.want.ID {
				t.Errorf("SelectOne() got id = %d, want id %d", got.ID, tt.want.ID)
			}
		})
	}
}

func TestSelect(t *testing.T) {
	dbconn, err := InitDB(DSN)
	if err != nil {
		t.Errorf("Error connect to db %v", err)
	}
	defer dbconn.Close()

	type args struct {
		db     *sql.DB
		limit  int64
		offset int64
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "Got", args: args{db: dbconn, limit: 5, offset: 0}, want: 5, wantErr: false},
		{name: "Got with limit", args: args{db: dbconn, limit: 1, offset: 1}, want: 1, wantErr: false},
		{name: "Got 0", args: args{db: dbconn, limit: 0, offset: 0}, want: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Select(tt.args.db, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				fmt.Printf("%+v\n", got)
				t.Errorf("Select() got count object = %d, want count %d", len(got), tt.want)
			}
		})
	}
}
