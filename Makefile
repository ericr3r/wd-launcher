BINS := build/bin/wd-launcher

PLUGINDIR := ~/.local/share/pop-launcher/plugins/wd_launcher 

.PHONY: all
all: $(BINS)

.PHONY: build
build:
	go build ./...

.PHONY: clean
clean:
	rm -rf build

$(BINS): build/bin/%: cmd/%/main.go
	@echo building $@
	go build -o $@ $^

$(PLUGINDIR):
	mkdir -p $@ 

.PHONY: install
install: $(PLUGINDIR) $(BINS)
	install -Dm700 $(BINS) $(PLUGINDIR)
	install -Dm0644 configs/plugin.ron $(PLUGINDIR)
	install -Dm0700 $(BINS) /home/eric/.local/bin/

.PHONY: uninstall
uninstall:
	-rm -rf  ~/.local/share/pop-launcher/plugins/wd_launcher
