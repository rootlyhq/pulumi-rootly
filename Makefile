PROJECT_NAME := rootly Package

SHELL            := /bin/bash
PACK             := rootly
PROJECT          := github.com/rootlyhq/pulumi-rootly
NODE_MODULE_NAME := @rootlyhq/${PACK}
TF_NAME          := ${PACK}
PROVIDER_PATH    := provider
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell pulumictl get version)

TESTPARALLELISM := 4

WORKING_DIR     := $(shell pwd)

OS := $(shell uname)
EMPTY_TO_AVOID_SED := 

prepare::
	@if test -z "${NAME}"; then echo "NAME not set"; exit 1; fi
	@if test -z "${REPOSITORY}"; then echo "REPOSITORY not set"; exit 1; fi
	@if test -z "${ORG}"; then echo "ORG not set"; exit 1; fi
	@if test ! -d "provider/cmd/pulumi-tfgen-x${EMPTY_TO_AVOID_SED}yz"; then "Project already prepared"; exit 1; fi

	mv "provider/cmd/pulumi-tfgen-x${EMPTY_TO_AVOID_SED}yz" provider/cmd/pulumi-tfgen-${NAME}
	mv "provider/cmd/pulumi-resource-x${EMPTY_TO_AVOID_SED}yz" provider/cmd/pulumi-resource-${NAME}

	if [[ "${OS}" != "Darwin" ]]; then \
		find ./ ! -path './.git/*' -type f -exec sed -i 's,github.com/pulumi/pulumi-[x]yz,${REPOSITORY},g' {} \; &> /dev/null; \
		find ./ ! -path './.git/*' -type f -exec sed -i 's/[x]yz/${NAME}/g' {} \; &> /dev/null; \
		find ./ ! -path './.git/*' -type f -exec sed -i 's/[a]bc/${ORG}/g' {} \; &> /dev/null; \
	fi

	# In MacOS the -i parameter needs an empty string to execute in place.
	if [[ "${OS}" == "Darwin" ]]; then \
		find ./ ! -path './.git/*' -type f -exec sed -i '' 's,github.com/pulumi/pulumi-[x]yz,${REPOSITORY},g' {} \; &> /dev/null; \
		find ./ ! -path './.git/*' -type f -exec sed -i '' 's/[x]yz/${NAME}/g' {} \; &> /dev/null; \
		find ./ ! -path './.git/*' -type f -exec sed -i '' 's/[a]bc/${ORG}/g' {} \; &> /dev/null; \
	fi

.PHONY: development provider build_sdks build_nodejs build_dotnet build_go build_python cleanup

development:: install_plugins provider lint_provider build_sdks install_sdks cleanup # Build the provider & SDKs for a development environment

# Required for the codegen action that runs in pulumi/pulumi and pulumi/pulumi-terraform-bridge
build:: install_plugins provider build_sdks install_sdks
only_build:: build

tfgen:: install_plugins
	cd provider && go build -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN}
	$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}
	(cd provider && VERSION=$(VERSION) go generate cmd/${PROVIDER}/main.go)

update_provider::
	(cd provider && go get -u github.com/rootlyhq/terraform-provider-rootly@latest)
	(cd provider && go mod tidy -compat=1.19)

provider:: tfgen install_plugins # build the provider binary
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER})

build_sdks:: install_plugins provider build_nodejs build_python build_go build_dotnet # build all the sdks

build_nodejs:: VERSION := $(shell pulumictl get version)
build_nodejs:: install_plugins tfgen # build the node sdk
	$(WORKING_DIR)/bin/$(TFGEN) nodejs --overlays provider/overlays/nodejs --out sdk/nodejs/
	cd sdk/nodejs/ && \
        yarn install && \
        yarn run tsc && \
        cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json && \
		sed -i.bak -e "s/pulumi\/rootly/rootly\/pulumi/g" ./bin/package.json

build_python:: PYPI_VERSION := $(shell pulumictl get version --language python)
build_python:: install_plugins tfgen # build the python sdk
	$(WORKING_DIR)/bin/$(TFGEN) python --overlays provider/overlays/python --out sdk/python/
	cd sdk/python/ && \
        cp ../../README.md . && \
        python3 setup.py clean --all 2>/dev/null && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
        sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
        rm ./bin/setup.py.bak && \
        cd ./bin && python3 setup.py build sdist

build_dotnet:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
build_dotnet:: install_plugins tfgen # build the dotnet sdk
	pulumictl get version --language dotnet
	$(WORKING_DIR)/bin/$(TFGEN) dotnet --overlays provider/overlays/dotnet --out sdk/dotnet/
	cd sdk/dotnet/ && \
		echo "${DOTNET_VERSION}" >version.txt && \
        dotnet build /p:Version=${DOTNET_VERSION}

build_go:: install_plugins tfgen # build the go sdk
	$(WORKING_DIR)/bin/$(TFGEN) go --overlays provider/overlays/go --out sdk/go/

lint_provider:: provider # lint the provider code
	cd provider && golangci-lint run -c ../.golangci.yml

cleanup:: # cleans up the temporary directory
	rm -r $(WORKING_DIR)/bin
	rm -f provider/cmd/${PROVIDER}/schema.go

help::
	@grep '^[^.#]\+:\s\+.*#' Makefile | \
		sed "s/\(.\+\):\s*\(.*\) #\s*\(.*\)/`printf "\033[93m"`\1`printf "\033[0m"`	\3 [\2]/" | \
		expand -t20
	@echo ""
	@echo "Version Management Commands:"
	@echo "  make version-show     - Show current and next versions"
	@echo "  make version-patch    - Bump patch version (3.0.0 → 3.0.1)"
	@echo "  make version-minor    - Bump minor version (3.0.0 → 3.1.0)"
	@echo "  make version-major    - Bump major version (3.0.0 → 4.0.0)"
	@echo "  make release-patch    - Bump patch and push tag (triggers CI release)"
	@echo "  make release-minor    - Bump minor and push tag (triggers CI release)"
	@echo "  make release-major    - Bump major and push tag (triggers CI release)"

clean::
	rm -rf sdk/{dotnet,nodejs,go,python}

install_plugins::
	[ -x $(shell which pulumi) ] || curl -fsSL https://get.pulumi.com | sh
	pulumi plugin install resource random 4.3.1

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::

install_go_sdk::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk

test::
	cd examples && go test -v -tags=all -parallel ${TESTPARALLELISM} -timeout 2h

# Version management targets
.PHONY: version-patch version-minor version-major version-show version-next version-help

version-show: # Show current and next versions
	@echo "Current version: $$(git describe --tags --abbrev=0 2>/dev/null || echo 'No tags found')"
	@echo "Next patch: $$(scripts/bump-version.sh show patch 2>/dev/null || echo 'No script found')"
	@echo "Next minor: $$(scripts/bump-version.sh show minor 2>/dev/null || echo 'No script found')"
	@echo "Next major: $$(scripts/bump-version.sh show major 2>/dev/null || echo 'No script found')"

version-patch: # Bump patch version (1.2.3 → 1.2.4)
	@scripts/bump-version.sh patch

version-minor: # Bump minor version (1.2.3 → 1.3.0)
	@scripts/bump-version.sh minor

version-major: # Bump major version (1.2.3 → 2.0.0)
	@scripts/bump-version.sh major

version-next: # Show next patch version
	@scripts/bump-version.sh show patch

version-help: # Show detailed version help
	@scripts/bump-version.sh help

# Release targets - these create git tags which trigger CI releases
.PHONY: release-patch release-minor release-major

release-patch: version-patch # Bump patch and push tag (triggers CI release)
	@echo "✅ Tag v$$(git describe --tags --abbrev=0) pushed"
	@echo "🚀 CI will automatically build and publish the release"

release-minor: version-minor # Bump minor and push tag (triggers CI release)
	@echo "✅ Tag v$$(git describe --tags --abbrev=0) pushed"
	@echo "🚀 CI will automatically build and publish the release"

release-major: version-major # Bump major and push tag (triggers CI release)
	@echo "✅ Tag v$$(git describe --tags --abbrev=0) pushed"
	@echo "🚀 CI will automatically build and publish the release"

