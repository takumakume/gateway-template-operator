package render

import (
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestOptions_toMap(t *testing.T) {
	type fields struct {
		Metadata metav1.ObjectMeta
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "map generate",
			fields: fields{
				Metadata: metav1.ObjectMeta{
					Namespace: "hoge",
				},
			},
			want: map[string]interface{}{
				"Metadata": metav1.ObjectMeta{
					Namespace: "hoge",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := &Options{
				Metadata: tt.fields.Metadata,
			}
			if got := opt.toMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Options.toMap() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
