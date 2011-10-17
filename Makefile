include $(GOROOT)/src/Make.inc

TARG=wombat/engine/expr
GOFILES=\
	consts.go\
	expressions.go\
	trees.go\
	value.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w $(GOFILES)
