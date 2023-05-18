COMPILER=go
GO=$(COMPILER)

shg:
	$(GO) build -o shg cmd.go

clean:
	rm -f shg
