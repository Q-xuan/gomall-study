MODULE = github.com/py/biz-demo/gomall

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/py/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --module github.com/py/biz-demo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: demo-link-fix
demo-link-fix:
	cd demo/demo_proto && golandci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/category_page.proto  --service frontend --module $(MODULE)/app/frontend -I ../../idl/

.PHONY: gen-user
gen-user:
	@cd rpc_gen  && cwgo client --type RPC --service user --module $(MODULE)/rpc_gen --I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server -type RPC --service user --module $(MODULE)/app/user --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto
	
.PHONY: gen-product
gen-product:
	@cd rpc_gen  && cwgo client --type RPC --service product --module $(MODULE)/rpc_gen --I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server -type RPC --service product --module $(MODULE)/app/product --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto