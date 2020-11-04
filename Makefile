TOPTARGETS := build
SUBDIRS := wsgateway

$(TOPTARGETS): $(SUBDIRS)

$(SUBDIRS):
	@$(MAKE) -C $@ $(MAKECMDGOALS)

clean:
	make -C vm clean
	rm -f bin/*

.PHONY: $(TOPTARGETS) $(SUBDIRS)

