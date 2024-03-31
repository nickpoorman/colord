# Define the target directory for binaries and scripts
BIN_DIR=~/bin

# Define the source files
SOURCES=colord_monitor.go colord_display.go

# Define the binary names (extracted from source file names)
BINARIES=$(SOURCES:.go=)

.PHONY: all build move clean

all: build move script

# Build each program
build:
	go build colord_monitor.go
	go build colord_display.go

# Make the binaries executable and move them to the target directory
move:
	chmod +x $(BINARIES)
	mv $(BINARIES) $(BIN_DIR)

# Create the bash script
script:
	echo '#!/bin/bash' > $(BIN_DIR)/colord
	echo 'nohup colord_monitor >/dev/null 2>&1 &' >> $(BIN_DIR)/colord
	chmod +x $(BIN_DIR)/colord

# Clean up the binaries and script from the target directory
clean:
	rm -f $(addprefix $(BIN_DIR)/, $(BINARIES)) $(BIN_DIR)/colord
