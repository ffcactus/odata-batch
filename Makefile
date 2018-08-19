all: client service

client:
	$(MAKE) -C client

service:
	$(MAKE) -C service

clean:
	$(MAKE) -C client clean
	$(MAKE) -C service clean

.PHONY: client service	