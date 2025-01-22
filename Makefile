.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/py/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-proto:
	@cd demo/demo_thrift && cwgo server --type RPC --module github.com/py/biz-demo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: demo-link-fix
demo-link-fix:
	cd demo/demo_proto && golandci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto  --service frontend --module github.com/py/biz-demo/gomall/app/frontend -I ../../idl/
.PHONY: gen-rpc-client
gen-rpc-client:
	@cd rpc_gen &&  cwgo client --type RPC --service user --module github.com/py/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/user.proto
.PHONY: gen-rpc-server
gen-rpc-server:
	@cd app/user && cwgo server -type RPC --service user --module github.com/py/biz-demo/gomall/app/user --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto