GO=go
PROGRAMM=wtop
PREFIX=/usr/

$(PROGRAMM):
	$(GO) build -o $(PROGRAMM)
install: $(PROGRAMM)
	cp -rfv ./wtop  $(PREFIX)/bin
clean:
	rm -fv ./wtop
