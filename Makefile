BUILD   = go build
INSTALL = install -s
MD      = install -d
RM      = rm -rf
TESTER  = go test -timeout 30s
ZIP     = zip -r

ifeq ($(UNAME), Windows_NT)
PLATFORM= Windows
TARGET= sudogo.exe
else
PLATFORM= $(shell go env GOOS | sed 's/^./\U&/')
TARGET= sudogo.x86_64
endif

VERSION= $(shell ./version.sh)
BINDIR= $(GOPATH)/bin
SOURCE= $(wildcard *.go sudoku/*.go ui/*.go)
TESTS= $(wildcard tests/*.go)
ZIPFILE= Sudogo-$(VERSION)-$(PLATFORM).zip


#-------------------------------------------------------------------------------
.PHONY: clean install itch mrproper play test uninstall

all: $(TARGET)


itch: $(ZIPFILE)


play: $(TARGET)
	./$<


$(TARGET): $(SOURCE)
	$(BUILD) -o $@ .
	strip $@


$(BINDIR):
	$(MD) $@


$(PLATFORM):
	$(MD) $@


$(ZIPFILE): $(PLATFORM) $(PLATFORM)/$(TARGET)
	$(ZIP) $@ $<


$(PLATFORM)/$(TARGET): $(TARGET) $(PLATFORM)
	$(INSTALL) $^


clean:
	$(RM) $(TARGET) $(PLATFORM) $(ZIPFILE)


install: $(TARGET)
	$(INSTALL) $(TARGET) $(BINDIR)


uninstall:
	$(RM) $(BINDIR)/$(TARGET)


test: $(TESTS) $(SOURCE)
	$(TESTER) $(TESTS)
