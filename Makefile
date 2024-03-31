# Define the Go compiler command
GO=go

# Define the target directory for binaries
BIN_DIR=~/bin

# Define the source files
SOURCES=colord.go colord_display.go

# Define the binary names (extracted from source file names)
BINARIES=$(SOURCES:.go=)

# Define the full path to the colord binary for the plist
COLORD_PATH=$(BIN_DIR)/colord

.PHONY: all build move clean

all: build move

# Build each program
build:
	$(GO) build colord.go
	$(GO) build colord_display.go

# Make the binaries executable and move them to the target directory
move:
	chmod +x $(BINARIES)
	mv $(BINARIES) $(BIN_DIR)

# Clean up the binaries from the target directory
clean:
	rm -f $(addprefix $(BIN_DIR)/, $(BINARIES))
