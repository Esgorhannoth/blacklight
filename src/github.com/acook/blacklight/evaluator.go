package main

func eval(ops []operation) {

	meta := NewMetaStack()
	current := NewSystemStack()
	meta.Push(current)

	doEval(meta, ops)
}

func doEval(meta *MetaStack, ops []operation) {
	var current *SystemStack
	defer rescue(meta)

	for _, op := range ops {
		s := *meta.Peek()
		current = s.(*SystemStack)

		switch op.(type) {
		case *metaOp:
			meta = op.Eval(meta).(*MetaStack)
		case operation:
			current = op.Eval(current).(*SystemStack)
		default:
			panic("urecognized operation:" + op.String())
		}

	}
}

func rescue(meta stack) {
	if err := recover(); err != nil {
		warn("evaluation error")
		warn("$meta:", meta.String())
		s := *meta.Peek()
		warn("@current:", s.String())
		panic(err)
	}
}