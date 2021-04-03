module github.com/planetdecred/pdapi/exchanges/rateserver

go 1.14

replace github.com/planetdecred/pdapi/exchanges/v3 => ../

require (
	github.com/decred/dcrd/certgen v1.1.1
	github.com/decred/dcrd/dcrutil/v3 v3.0.0
	github.com/planetdecred/pdapi/exchanges/v3 v3.0.0
	github.com/decred/slog v1.1.0
	github.com/jessevdk/go-flags v1.4.1-0.20200711081900-c17162fe8fd7
	github.com/jrick/logrotate v1.0.0
	google.golang.org/grpc v1.32.0
)
