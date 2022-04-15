INSTALL = install -s
MD      = install -d
RM      = rm -rf
TESTER  = go test -timeout 30s
ZIP     = zip -r
CC     ?= clang
CXX    ?= clang++

ifeq ($(UNAME), Windows_NT)
PLATFORM= Windows
else
PLATFORM= $(shell go env GOOS | sed 's/^./\U&/')
endif


ifeq ($(PLATFORM), Windows)
TARGET= sudogo.exe
ifeq ($(shell go env GOARCH), amd64)
CC = x86_64-w64-mingw32-gcc
CXX= x86_64-w64-mingw32-gcc++
else
CC = i686-w64-mingw32-gcc
CXX= i686-w64-mingw32-gcc++
endif
CGO_ENABLED= 1
else
TARGET= sudogo.x86_64
endif

BUILD= CC=$(CC) CXX=$(CXX) go build

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
