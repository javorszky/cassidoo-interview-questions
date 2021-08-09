package august082021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewField(t *testing.T) {
	type args struct {
		gridSize int
		fields   []Coordinate
	}
	tests := []struct {
		name    string
		args    args
		want    Field
		wantErr bool
	}{
		{
			name: "generates a smol field with no mines",
			args: args{
				gridSize: 3,
				fields:   nil,
			},
			want: Field{
				size: 3,
				fields: map[Coordinate]FieldType{
					{1, 1}: zero,
					{2, 1}: zero,
					{3, 1}: zero,
					{1, 2}: zero,
					{2, 2}: zero,
					{3, 2}: zero,
					{1, 3}: zero,
					{2, 3}: zero,
					{3, 3}: zero,
				},
			},
			wantErr: false,
		},
		{
			name: "generates a smol field with 1 mine at 1,1",
			args: args{
				gridSize: 3,
				fields: []Coordinate{
					{1, 1},
				},
			},
			want: Field{
				size: 3,
				fields: map[Coordinate]FieldType{
					{1, 1}: mine,
					{2, 1}: one,
					{3, 1}: zero,
					{1, 2}: one,
					{2, 2}: one,
					{3, 2}: zero,
					{1, 3}: zero,
					{2, 3}: zero,
					{3, 3}: zero,
				},
			},
			wantErr: false,
		},
		{
			name: "generates a smol field with 1 mine at 2,2",
			args: args{
				gridSize: 3,
				fields: []Coordinate{
					{2, 2},
				},
			},
			want: Field{
				size: 3,
				fields: map[Coordinate]FieldType{
					{1, 1}: one,
					{2, 1}: one,
					{3, 1}: one,
					{1, 2}: one,
					{2, 2}: mine,
					{3, 2}: one,
					{1, 3}: one,
					{2, 3}: one,
					{3, 3}: one,
				},
			},
			wantErr: false,
		},
		{
			name: "generates a smol field with 1 mine at 3,3",
			args: args{
				gridSize: 3,
				fields: []Coordinate{
					{3, 3},
				},
			},
			want: Field{
				size: 3,
				fields: map[Coordinate]FieldType{
					{1, 1}: zero,
					{2, 1}: zero,
					{3, 1}: zero,
					{1, 2}: zero,
					{2, 2}: one,
					{3, 2}: one,
					{1, 3}: zero,
					{2, 3}: one,
					{3, 3}: mine,
				},
			},
			wantErr: false,
		},
		{
			name: "generates a smol field with 8 mines around the edges",
			args: args{
				gridSize: 3,
				fields: []Coordinate{
					{1, 1},
					{2, 1},
					{3, 1},
					{1, 2},
					{3, 2},
					{1, 3},
					{2, 3},
					{3, 3},
				},
			},
			want: Field{
				size: 3,
				fields: map[Coordinate]FieldType{
					{1, 1}: mine,
					{2, 1}: mine,
					{3, 1}: mine,
					{1, 2}: mine,
					{2, 2}: eight,
					{3, 2}: mine,
					{1, 3}: mine,
					{2, 3}: mine,
					{3, 3}: mine,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewField(tt.args.gridSize, tt.args.fields)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestField_String(t *testing.T) {
	type fields struct {
		size   int
		fields []Coordinate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "prints the field's string representation correctly for a field with no mines",
			fields: fields{
				size:   3,
				fields: nil,
			},
			want: `___
___
___`,
		},
		{
			name: "prints the field's string representation correctly for a field with a mine at 1,1",
			fields: fields{
				size: 3,
				fields: []Coordinate{
					{1, 1},
				},
			},
			want: `x1_
11_
___`,
		},
		{
			name: "prints the field's string representation correctly for a field with a mine at 3,3",
			fields: fields{
				size: 3,
				fields: []Coordinate{
					{3, 3},
				},
			},
			want: `___
_11
_1x`,
		},
		{
			name: "prints the field's string representation correctly for a field with mines around the edges",
			fields: fields{
				size: 3,
				fields: []Coordinate{
					{1, 1},
					{2, 1},
					{3, 1},
					{1, 2},
					{3, 2},
					{1, 3},
					{2, 3},
					{3, 3},
				},
			},
			want: `xxx
x8x
xxx`,
		},
		{
			name: "prints the field's string representation correctly for a field with mines at 1,1, 1,3, 2,3",
			fields: fields{
				size: 3,
				fields: []Coordinate{
					{1, 1},
					{1, 3},
					{2, 3},
				},
			},
			want: `x1_
331
xx1`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := NewField(tt.fields.size, tt.fields.fields)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, f.String())
		})
	}
}
