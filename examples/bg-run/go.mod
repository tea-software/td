module github.com/gotd/td/examples/bg-run

go 1.16

require (
	github.com/gotd/contrib v0.12.0
	github.com/gotd/td v0.54.1
	go.uber.org/zap v1.21.0
)

replace github.com/gotd/td => ./../..
