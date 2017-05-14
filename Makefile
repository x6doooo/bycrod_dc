define build_env
	eGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags=-s -o ./bin/bycrod_dc_$(1) ./main.go
endef

all: clean prod

clean: clean_test clean_prod

test:
	$(call build_env,test)

prod:
	$(call build_env,prod)
