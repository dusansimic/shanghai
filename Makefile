COMPILER=go
GO=CGO_ENABLED=0 $(COMPILER)
PROGRAM=shanghai

$(PROGRAM):
	$(GO) build $(PROGRAM).go

clean:
	rm -f $(PROGRAM)
