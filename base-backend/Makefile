# Makefile - generate grpc to generate/grpc
PROTOC           := protoc
PROTO_DIR        := modules
OUT_DIR          := generate/grpc

# modules that require protoc-gen-validate
GRPC_MODULES              := app_user
# modules that do not require validation plugin
GRPC_MODULES_NO_VALIDATION := health

# go plugin options
GO_OUT       := --go_out=$(OUT_DIR) --go_opt=paths=source_relative
GO_GRPC_OUT  := --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative
VALIDATE_OUT_FLAG := --validate_out="lang=go:$(OUT_DIR)"

# Try to locate PGV proto dir in GOPATH module cache (vendor not used here)
PGV_PROTO_DIR := $(shell ls -d $(go env GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@* 2>/dev/null | head -n1)
TMP_PATTERNS := '*.pb.go' '*_grpc.pb.go' '*.pb.validate.go'

OUT_DIR := generated/grpc

protoc:
	@echo "Ensure output dir: $(OUT_DIR)"
	@mkdir -p $(OUT_DIR)
	@echo

	@echo "=== Running protoc for validated modules (validate files written alongside pb files) ==="
	$(foreach mod,$(GRPC_MODULES), \
	  echo "-> $(mod)"; \
	  protoc \
	    --proto_path=$(PGV_PROTO_DIR) \
	    --go_opt=paths=source_relative \
	    --go_out=. \
	    --go-grpc_opt=paths=source_relative \
	    --go-grpc_out=. \
	    --validate_out="lang=go:." \
	    --validate_opt="paths=source_relative" \
	    modules/$(mod)/grpc/$(mod).proto; \
	)

	@echo
	@echo "=== Running protoc for non-validated modules ==="
	$(foreach mod,$(GRPC_MODULES_NO_VALIDATION), \
	  echo "-> $(mod)"; \
	  protoc \
	    --proto_path=$(PGV_PROTO_DIR) \
	    --go_opt=paths=source_relative \
	    --go_out=. \
	    --go-grpc_opt=paths=source_relative \
	    --go-grpc_out=. \
	    modules/$(mod)/grpc/$(mod).proto; \
	)

	@echo
	@echo "=== Moving generated pb, grpc and validate files into $(OUT_DIR) ==="
	@set -e; \
	# find generated files scattered in repo and move them to OUT_DIR
	find . -type f \( -name '*.pb.go' -o -name '*_grpc.pb.go' -o -name '*.validate.go' \) ! -path "./$(OUT_DIR)/*" -print0 | \
	while IFS= read -r -d $$'\0' file; do \
	  bn=$$(basename "$$file"); \
	  echo "  moving $$file -> $(OUT_DIR)/$$bn"; \
	  mv -f "$$file" "$(OUT_DIR)/$$bn" || { echo "ERROR moving $$file"; exit 1; }; \
	done; \
	# cleanup empty directories under modules (optional)
	find modules -type d -empty -delete || true

	@echo "=== Done. All generated files are in $(OUT_DIR) ==="
