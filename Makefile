GOPKG ?=	moul.io/captcha
DOCKER_IMAGE ?=	moul/captcha
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp

	echo 'foo@bar:~$$ captcha -engine=math' > .tmp/usage.txt
	(echo 42 | captcha -seed=42 -engine=math -retries=1 2>/dev/null || true) >> .tmp/usage.txt
	echo >> .tmp/usage.txt

	echo 'foo@bar:~$$ captcha -engine=math' >> .tmp/usage.txt
	(echo 42 | captcha -seed=4242 -engine=math -retries=1 2>/dev/null || true) >> .tmp/usage.txt
	echo >> .tmp/usage.txt

	echo 'foo@bar:~$$ captcha -engine=banner' >> .tmp/usage.txt
	(echo 42 | captcha -seed=42 -engine=banner -retries=1 2>/dev/null || true) >> .tmp/usage.txt
	echo >> .tmp/usage.txt

	echo 'foo@bar:~$$ captcha -engine=banner' >> .tmp/usage.txt
	(echo 42 | captcha -seed=4242 -engine=banner -retries=1 2>/dev/null || true) >> .tmp/usage.txt
	echo >> .tmp/usage.txt

	embedmd -w README.md
	sed -i 's/[[:blank:]]*$$//' README.md
	rm -rf .tmp
.PHONY: generate

lint:
	cd tool/lint; make
.PHONY: lint
