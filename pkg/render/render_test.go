package render

import "testing"

func TestRender_Render(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		tmpl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "render string",
			fields: fields{
				data: map[string]interface{}{
					"key1": "value1",
				},
			},
			args: args{
				tmpl: "value is {{ .key1 }}",
			},
			want: "value is value1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Render{
				data: tt.fields.data,
			}
			got, err := r.Render(tt.args.tmpl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Render.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
