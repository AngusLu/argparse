package argparse

import (
	"testing"
)

func Test_arg_check(t *testing.T) {
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "test argument check arg with short name",
			fields:    fields{sname: "d", size: 1},
			args:      args{s: "-d"},
			wantCount: 1,
			wantErr:   false,
		},
		{
			name:      "test argument check arg with long name",
			fields:    fields{lname: "dir", size: 1},
			args:      args{s: "--dir"},
			wantCount: 1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			gotCount, err := a.check(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("arg.check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("arg.check() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_arg_checkLongName(t *testing.T) {
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "test argument check long name",
			fields: fields{
				lname: "dir",
				size:  1,
			},
			args: args{
				s: "--dir",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			if got := a.checkLongName(tt.args.s); got != tt.want {
				t.Errorf("arg.checkLongName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arg_checkShortName(t *testing.T) {
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test argument check short name",
			fields: fields{
				sname: "d",
				size:  1,
			},
			args:    args{s: "-d"},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			got, err := a.checkShortName(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("arg.checkShortName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("arg.checkShortName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arg_reduce(t *testing.T) {
	a0 := []string{"-d"}
	a1 := []string{"--dir"}
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		pos  int
		args *[]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test argument reduce with short name",
			fields: fields{
				sname:  "d",
				size:   0,
				unique: true,
			},
			args: args{
				pos:  0,
				args: &a0,
			},
		},
		{
			name: "test argument reduce with long name",
			fields: fields{
				lname:  "dir",
				size:   0,
				unique: true,
			},
			args: args{
				pos:  0,
				args: &a1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			a.reduce(tt.args.pos, tt.args.args)
			if (*tt.args.args)[0] != "" {
				t.Errorf("arg reduce failed")
			}
		})
	}
}

func Test_arg_reduceLongName(t *testing.T) {
	a0 := []string{"--dir"}
	a1 := []string{"--file", "/config.yml"}
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		pos  int
		args *[]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test argument reduce long name 1",
			fields: fields{
				lname: "dir",
				size:  0,
			},
			args: args{
				pos:  0,
				args: &a0,
			},
		},
		{
			name: "test argument redue long name 2",
			fields: fields{
				lname: "file",
				size:  1,
			},
			args: args{
				pos:  0,
				args: &a1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			a.reduceLongName(tt.args.pos, tt.args.args)
			for i := tt.args.pos; i < tt.args.pos+tt.fields.size; i++ {
				if (*tt.args.args)[i] != "" {
					t.Errorf("arg reduce fail")
				}
			}
		})
	}
}

func Test_arg_reduceShortName(t *testing.T) {
	a0 := []string{"-d"}
	a1 := []string{"-f", "/config.yml"}
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		pos  int
		args *[]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test argument reduce short name 1",
			fields: fields{
				sname: "d",
				size:  0,
			},
			args: args{
				pos:  0,
				args: &a0,
			},
		},
		{
			name: "test argument reduce short name 2",
			fields: fields{
				sname: "f",
				size:  1,
			},
			args: args{
				pos:  0,
				args: &a1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			a.reduceShortName(tt.args.pos, tt.args.args)
			for i := tt.args.pos; i < tt.args.pos+tt.fields.size; i++ {
				if (*tt.args.args)[i] != "" {
					t.Errorf("arg reduce fail")
				}
			}
		})
	}
}

func Test_arg_name(t *testing.T) {
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test argument get name 1",
			fields: fields{
				sname: "d",
			},
			want: "-d",
		},
		{
			name: "test argument get name 2",
			fields: fields{
				lname: "dir",
			},
			want: "--dir",
		},
		{
			name: "test argument get name 3",
			fields: fields{
				sname: "d",
				lname: "dir",
			},
			want: "-d | --dir",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			if got := a.name(); got != tt.want {
				t.Errorf("arg.name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arg_parse(t *testing.T) {
	var b bool
	var s string
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test argument parse 1",
			fields: fields{
				sname:  "d",
				size:   0,
				value:  &b,
				unique: true,
			},
			args: args{
				args: []string{},
			},
			wantErr: false,
		},
		{
			name: "test argument parse 2",
			fields: fields{
				sname:  "f",
				size:   1,
				unique: true,
				value:  &s,
			},
			args:    args{args: []string{"/config.yml"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			if err := a.parse(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("arg.parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_arg_parseType(t *testing.T) {
	var s string
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test argument parseType 1",
			fields: fields{
				sname:  "f",
				value:  &s,
				size:   1,
				unique: true,
			},
			args:    args{args: []string{"/config.yml"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			if err := a.parseType(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("arg.parseType() error = %v, wantErr %v", err, tt.wantErr)
			}
			if s != (tt.args.args)[0] {
				t.Errorf("arg.parseType() string = %s, wantErr %s", s, (tt.args.args)[0])
			}
		})
	}
}

func Test_arg_parseString(t *testing.T) {
	var s string
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test argument parseString",
			fields: fields{
				value:  &s,
				unique: true,
				size:   1,
			},
			args:    args{args: []string{"value"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			if err := a.parseString(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("arg.parseString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_arg_parseInt(t *testing.T) {
	var i int
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test argument parseInt",
			fields: fields{
				value:  &i,
				size:   1,
				unique: true,
			},
			args:    args{args: []string{"123"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			if err := a.parseInt(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("arg.parseInt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_arg_parseBool(t *testing.T) {
	var b bool
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test argument parseBool",
			fields: fields{
				value:  &b,
				unique: true,
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			if err := a.parseBool(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("arg.parseBool() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_arg_parseFloat(t *testing.T) {
	var f float64
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test argument parseFloat",
			fields: fields{
				value:  &f,
				unique: true,
			},
			args:    args{args: []string{"1.23"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
			}
			if err := a.parseFloat(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("arg.parseFloat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_arg_parseStringSlice(t *testing.T) {
	var s []string
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
		opts   *Option
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test argument parseStringSlice",
			fields: fields{
				size:   1,
				value:  &s,
				unique: false,
			},
			args: args{
				args: []string{
					"a", "b", "c",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
				opts:   tt.fields.opts,
			}
			if err := a.parseStringSlice(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("arg.parseStringSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_arg_parseIntSlice(t *testing.T) {
	var i []int
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
		opts   *Option
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test argument parseIntSlice",
			fields: fields{
				value:  &i,
				size:   1,
				unique: false,
			},
			args: args{
				args: []string{"1", "2", "3"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
				opts:   tt.fields.opts,
			}
			if err := a.parseIntSlice(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("arg.parseIntSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_arg_parseFloatSlice(t *testing.T) {
	var f []float64
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
		opts   *Option
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test argument parseFloatSlice",
			fields: fields{
				value:  &f,
				size:   1,
				unique: false,
			},
			args: args{
				args: []string{"1.2", "2.3", "3.5"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
				opts:   tt.fields.opts,
			}
			if err := a.parseFloatSlice(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("arg.parseFloatSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_arg_getType(t *testing.T) {
	var s string
	var b bool
	var i int
	var f float64
	type fields struct {
		sname  string
		lname  string
		size   int
		value  interface{}
		unique bool
		parsed bool
		opts   *Option
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "test get argument type 1",
			fields: fields{value: &s},
			want:   "string",
		},
		{
			name:   "test get argument type 2",
			fields: fields{value: &b},
			want:   "bool",
		},
		{
			name:   "test get argument type 3",
			fields: fields{value: &i},
			want:   "int",
		},
		{
			name:   "test get argument type 4",
			fields: fields{value: &f},
			want:   "float",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &arg{
				sname:  tt.fields.sname,
				lname:  tt.fields.lname,
				size:   tt.fields.size,
				value:  tt.fields.value,
				unique: tt.fields.unique,
				parsed: tt.fields.parsed,
				opts:   tt.fields.opts,
			}
			if got := a.getType(); got != tt.want {
				t.Errorf("arg.getType() = %v, want %v", got, tt.want)
			}
		})
	}
}
