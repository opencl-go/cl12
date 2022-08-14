package cl12_test

import (
	"testing"
	"unsafe"

	cl "github.com/opencl-go/cl12"
)

func TestImageFormatSize(t *testing.T) {
	t.Parallel()
	if (cl.ImageFormatByteSize != unsafe.Sizeof(cl.ImageFormat{})) {
		t.Errorf("byte size mismatch")
	}
}

func TestImageDescSize(t *testing.T) {
	t.Parallel()
	if (cl.ImageDescByteSize != unsafe.Sizeof(cl.ImageDesc{})) {
		t.Errorf("byte size mismatch")
	}
}
