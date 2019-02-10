package tracer

import(
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of 
//tracing events through code 
type Tracer interface{
	// Interfaces are simply definations that cannot be implemented
	// Accepts zero or more arguments of any type
	Trace(...interface{})
	
}

type tracer struct{
	out io.Writer
}

func (tr *tracer) Trace(a...interface{}){
	fmt.Fprint(tr.out,a...)
	fmt.Fprintln(tr.out)
}


//New function allows us to create a  new instance  the Tracer type by passing in the io.Writer interface
func New(w io.Writer) Tracer{
	return &tracer{out:w}
}