package argparse

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Parser
	}{
		{
			name: "test new parser",
			want: &Parser{args: make([]*arg, 0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_addArg(t *testing.T) {
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		a *arg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test addArg to parser",
			fields: fields{
				args:   nil,
				parsed: false,
			},
			args: args{
				a: &arg{sname: "a"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if err := p.addArg(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("Parser.addArg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_typeVar(t *testing.T) {
	var s string
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		i           interface{}
		defVal      interface{}
		short       string
		long        string
		description string
		size        int
		unique      bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test parser typeVar func",
			fields: fields{},
			args: args{
				i:      &s,
				defVal: "",
				short:  "s",
				long:   "string",
				size:   1,
				unique: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			p.typeVar(tt.args.i, tt.args.defVal, tt.args.short, tt.args.long, tt.args.description, tt.args.size, tt.args.unique, nil)
		})
	}
}

func TestParser_Parse(t *testing.T) {
	var conf string
	var port int
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		a []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test parser run Parse",
			fields: fields{
				args: []*arg{
					&arg{
						sname:  "c",
						size:   1,
						unique: true,
						value:  &conf,
					},
					&arg{
						sname:  "p",
						size:   1,
						unique: true,
						value:  &port,
					},
				},
			},
			args: args{
				a: []string{
					"-c", "/root/config.ini",
					"-p", "3000",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if err := p.Parse(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_parse(t *testing.T) {
	var conf string
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		args *[]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test parser parse func",
			fields: fields{
				args: []*arg{
					&arg{
						sname:  "c",
						size:   1,
						unique: true,
						value:  &conf,
					},
				},
			},
			args:    args{args: &[]string{"-c", "/config.yml"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if err := p.parse(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Parser.parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_parseArguemtns(t *testing.T) {
	var conf string
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		args *[]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test parser parseArguments",
			fields: fields{
				args: []*arg{
					&arg{
						sname:  "c",
						size:   1,
						unique: true,
						value:  &conf,
					},
				},
			},
			args:    args{args: &[]string{"-vv"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if err := p.parseArguemtns(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseArguemtns() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_String(t *testing.T) {
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		defValue string
		short    string
		long     string
		desc     string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add string flag to parser",
			fields: fields{},
			args: args{
				short: "c",
				long:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if got := p.String(tt.args.defValue, tt.args.short, tt.args.long, tt.args.desc, nil); got == nil {
				t.Errorf("Parser.String() = %v", got)
			}
		})
	}
}

func TestParser_StringVar(t *testing.T) {
	var s string
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		i      *string
		defVal string
		short  string
		long   string
		desc   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add string flag to parser with ptr",
			fields: fields{},
			args: args{
				i:     &s,
				short: "s",
				long:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			p.StringVar(tt.args.i, tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, nil)
		})
	}
}

func TestParser_Int(t *testing.T) {
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		defVal int
		short  string
		long   string
		desc   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add int flag to parser",
			fields: fields{},
			args: args{
				short: "i",
				long:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if got := p.Int(tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, nil); got == nil {
				t.Errorf("Parser.Int() = %v", got)
			}
		})
	}
}

func TestParser_IntVar(t *testing.T) {
	var i int
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		i      *int
		defVal int
		short  string
		long   string
		desc   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add int flag to parser (ptr)",
			fields: fields{},
			args: args{
				i:     &i,
				short: "i",
				long:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			p.IntVar(tt.args.i, tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, nil)
		})
	}
}

func TestParser_Bool(t *testing.T) {
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		defVal bool
		short  string
		long   string
		desc   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add bool flag to parser",
			fields: fields{},
			args: args{
				short: "b",
				long:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if got := p.Bool(tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, nil); got == nil {
				t.Errorf("Parser.Bool() = %v", got)
			}
		})
	}
}

func TestParser_BoolVar(t *testing.T) {
	var b bool
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		i      *bool
		defVal bool
		short  string
		long   string
		desc   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add bool flag to parser (ptr)",
			fields: fields{},
			args: args{
				i:     &b,
				short: "b",
				long:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			p.BoolVar(tt.args.i, tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, nil)
		})
	}
}

func TestParser_Float(t *testing.T) {
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		defVal float64
		short  string
		long   string
		desc   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add float flag to parser",
			fields: fields{},
			args: args{
				short: "f",
				long:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if got := p.Float(tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, nil); got == nil {
				t.Errorf("Parser.Float() = %v", got)
			}
		})
	}
}

func TestParser_FloatVar(t *testing.T) {
	var f float64
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		i      *float64
		defVal float64
		short  string
		long   string
		desc   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add float flag to parser (ptr)",
			fields: fields{},
			args: args{
				i:     &f,
				short: "f",
				long:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			p.FloatVar(tt.args.i, tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, nil)
		})
	}
}

func TestParser_StringSlice(t *testing.T) {
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		defVal []string
		short  string
		long   string
		desc   string
		opts   *Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add string slice to parser",
			fields: fields{},
			args: args{
				short: "d",
				long:  "",
				opts:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if got := p.StringSlice(tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, tt.args.opts); got == nil {
				t.Errorf("Parser.StringSlice() = %v", got)
			}
		})
	}
}

func TestParser_StringSliceVar(t *testing.T) {
	var s []string
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		i      *[]string
		defVal []string
		short  string
		long   string
		desc   string
		opts   *Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add string slice to parser (ptr)",
			fields: fields{},
			args: args{
				i:     &s,
				short: "d",
				long:  "",
				opts:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			p.StringSliceVar(tt.args.i, tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, tt.args.opts)
		})
	}
}

func TestParser_IntSlice(t *testing.T) {
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		defVal []int
		short  string
		long   string
		desc   string
		opts   *Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add int slice to parser",
			fields: fields{},
			args: args{
				short: "i",
				long:  "",
				opts:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if got := p.IntSlice(tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, tt.args.opts); got == nil {
				t.Errorf("Parser.IntSlice() = %v", got)
			}
		})
	}
}

func TestParser_IntSliceVar(t *testing.T) {
	var i []int
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		i      *[]int
		defVal []int
		short  string
		long   string
		desc   string
		opts   *Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add int slice to parser (ptr)",
			fields: fields{},
			args: args{
				i:     &i,
				short: "i",
				long:  "",
				opts:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			p.IntSliceVar(tt.args.i, tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, tt.args.opts)
		})
	}
}

func TestParser_FloatSliceVar(t *testing.T) {
	var f []float64
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		i      *[]float64
		defVal []float64
		short  string
		long   string
		desc   string
		opts   *Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add float slice to parser (ptr)",
			fields: fields{},
			args: args{
				i:     &f,
				short: "f",
				long:  "",
				opts:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			p.FloatSliceVar(tt.args.i, tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, tt.args.opts)
		})
	}
}

func TestParser_FloatSlice(t *testing.T) {
	type fields struct {
		args   []*arg
		parsed bool
	}
	type args struct {
		defVal []float64
		short  string
		long   string
		desc   string
		opts   *Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add float slice to parser",
			fields: fields{},
			args: args{
				short: "f",
				long:  "",
				opts:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:   tt.fields.args,
				parsed: tt.fields.parsed,
			}
			if got := p.FloatSlice(tt.args.defVal, tt.args.short, tt.args.long, tt.args.desc, tt.args.opts); got == nil {
				t.Errorf("Parser.FloatSlice() = %v", got)
			}
		})
	}
}

func TestParser_Help(t *testing.T) {
	type fields struct {
		args     []*arg
		showHelp bool
		parsed   bool
	}
	type args struct {
		short string
		long  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test add help flag to parser",
			fields: fields{},
			args: args{
				short: "h",
				long:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args:     tt.fields.args,
				showHelp: tt.fields.showHelp,
				parsed:   tt.fields.parsed,
			}
			p.Help(tt.args.short, tt.args.long)
		})
	}
}
