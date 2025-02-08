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
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/checkout_page.proto  --service frontend --module $(MODULE)/app/frontend -I ../../idl/

.PHONY: gen-user
gen-user:
	@cd rpc_gen  && cwgo client --type RPC --service user --module $(MODULE)/rpc_gen --I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server -type RPC --service user --module $(MODULE)/app/user --pass "-use github.com/py/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto
	
.PHONY: gen-product
gen-product:
	@cd rpc_gen  && cwgo client --type RPC --service product --module $(MODULE)/rpc_gen --I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server -type RPC --service product --module $(MODULE)/app/product --pass "-use github.com/py/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen  && cwgo client --type RPC --service cart --module $(MODULE)/rpc_gen --I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server -type RPC --service cart --module $(MODULE)/app/cart --pass "-use github.com/py/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen  && cwgo client --type RPC --service payment --module $(MODULE)/rpc_gen --I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server -type RPC --service payment --module $(MODULE)/app/payment --pass "-use github.com/py/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/payment.proto
	
.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen  && cwgo client --type RPC --service checkout --module $(MODULE)/rpc_gen --I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server -type RPC --service checkout --module $(MODULE)/app/checkout --pass "-use github.com/py/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/checkout.proto

.PHONY: gen-order
gen-order:
	@cd rpc_gen  && cwgo client --type RPC --service order --module $(MODULE)/rpc_gen --I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server -type RPC --service order --module $(MODULE)/app/order --pass "-use github.com/py/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.proto