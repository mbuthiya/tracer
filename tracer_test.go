package tracer

import(
	"bytes"
	"testing"

)

func TestNew(t *testing.T){
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil{
		t.Error("Return from New should not be nil")
	}else{
		tracer.Trace("Hello Trace package")
		if buf.String() != "Hello Trace package\n"{
			t.Errorf("Tracer should write '%s'.",buf.String())
		}
	}
}