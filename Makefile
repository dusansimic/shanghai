COMPILER=go
GO=$(COMPILER)
PROGRAM=shanghai

$(PROGRAM):
	$(GO) build $(PROGRAM).go

clean:
	rm -f $(PROGRAM)
