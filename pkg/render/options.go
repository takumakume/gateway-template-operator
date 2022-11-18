package render

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Options struct {
	Metadata metav1.ObjectMeta
}

func (opt *Options) toMap() map[string]interface{} {
	return map[string]interface{}{
		"Metadata": opt.Metadata,
	}
}
