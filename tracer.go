package trace 

// Tracer is the interface that describes an object capable of 
//tracing events through code 
type Tracer interface{
	// Interfaces are simply definations that cannot be implemented
	// Accepts zero or more arguments of any type
	Trace(...interface{})
	
}