# Define the Go compiler command
GO=go

# Define the target directory for binaries
BIN_DIR=~/bin

# Define the plist destination. Use ~/Library/LaunchAgents for per-user services.
PLIST=~/Library/LaunchAgents/com.nickpoorman.colord.plist

# Define the source files
SOURCES=colord.go colord_display.go

# Define the binary names (extracted from source file names)
BINARIES=$(SOURCES:.go=)

# Define the full path to the colord binary for the plist
COLORD_PATH=$(BIN_DIR)/colord

.PHONY: all build move daemonize clean

all: build move daemonize

# Build each program
build:
	$(GO) build colord.go
	$(GO) build colord_display.go

# Make the binaries executable and move them to the target directory
move:
	chmod +x $(BINARIES)
	mv $(BINARIES) $(BIN_DIR)

# Daemonize colord
daemonize:
	@if [ ! -f $(PLIST) ]; then \
		echo Creating plist file to daemonize colord; \
		echo "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" > $(PLIST); \
		echo "<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">" >> $(PLIST); \
		echo "<plist version=\"1.0\">" >> $(PLIST); \
		echo "<dict>" >> $(PLIST); \
		echo "  <key>Label</key>" >> $(PLIST); \
		echo "  <string>com.nickpoorman.colord</string>" >> $(PLIST); \
		echo "  <key>ProgramArguments</key>" >> $(PLIST); \
		echo "  <array>" >> $(PLIST); \
		echo "      <string>$(COLORD_PATH)</string>" >> $(PLIST); \
		echo "  </array>" >> $(PLIST); \
		echo "  <key>RunAtLoad</key>" >> $(PLIST); \
		echo "  <true/>" >> $(PLIST); \
		echo "  <key>KeepAlive</key>" >> $(PLIST); \
		echo "  <true/>" >> $(PLIST); \
		echo "  <key>StandardErrorPath</key>" >> $(PLIST); \
		echo "  <string>/tmp/com.nickpoorman.colord.err</string>" >> $(PLIST); \
		echo "  <key>StandardOutPath</key>" >> $(PLIST); \
		echo "  <string>/tmp/com.nickpoorman.colord.out</string>" >> $(PLIST); \
		echo "</dict>" >> $(PLIST); \
		echo "</plist>" >> $(PLIST); \
		launchctl load $(PLIST); \
		echo "Daemon setup complete"; \
	else \
		echo "Plist already exists"; \
	fi

# Clean up the binaries from the target directory
clean:
	rm -f $(addprefix $(BIN_DIR)/, $(BINARIES))
	launchctl unload $(PLIST)
	rm -f $(PLIST)
