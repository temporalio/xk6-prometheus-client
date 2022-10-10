k6: go.mod go.sum *.go
	xk6 build --with github.com/temporalio/xk6-prometheus-client=.
