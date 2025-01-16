.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/py/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-proto:
	@cd demo/demo_thrift && cwgo server --type RPC --module github.com/py/biz-demo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift
