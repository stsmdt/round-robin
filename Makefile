GO ?= go

.PHONY: test
test:
	echo "mode: count" > coverage.out
	$(GO) test -v -covermode=count -coverprofile=profile.out > tmp.out; \
	cat tmp.out; \
	if grep -q "^--- FAIL" tmp.out; then \
		rm tmp.out; \
		exit 1; \
	elif grep -q "build failed" tmp.out; then \
		rm tmp.out; \
		exit 1; \
	elif grep -q "setup failed" tmp.out; then \
		rm tmp.out; \
		exit 1; \
	fi; \
	if [ -f profile.out ]; then \
		cat profile.out | grep -v "mode:" >> coverage.out; \
		rm profile.out; \
	fi; \
