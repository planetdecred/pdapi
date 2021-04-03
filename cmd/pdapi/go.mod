module github.com/planetdecred/pdapi/cmd/pdapi

go 1.14

replace (
	github.com/planetdecred/pdapi/db/dcrpg/v6 => ../../db/dcrpg/
	github.com/planetdecred/pdapi/exchanges/v3 => ../../exchanges/
	github.com/planetdecred/pdapi/gov/v4 => ../../gov/
	github.com/planetdecred/pdapi/v6 => ../../
)

require (
	github.com/caarlos0/env/v6 v6.4.0
	github.com/decred/dcrd/blockchain/standalone/v2 v2.0.0
	github.com/decred/dcrd/chaincfg/chainhash v1.0.2
	github.com/decred/dcrd/chaincfg/v3 v3.0.0
	github.com/decred/dcrd/dcrutil/v3 v3.0.0
	github.com/decred/dcrd/rpc/jsonrpc/types/v2 v2.3.0
	github.com/decred/dcrd/rpcclient/v6 v6.0.2
	github.com/decred/dcrd/txscript/v3 v3.0.0
	github.com/decred/dcrd/wire v1.4.0
	github.com/decred/slog v1.1.0
	github.com/didip/tollbooth/v5 v5.2.0
	github.com/dmigwi/go-piparser/proposals v0.0.0-20191219171828-ae8cbf4067e1
	github.com/go-chi/chi v1.5.1
	github.com/go-chi/docgen v1.1.0
	github.com/google/gops v0.3.13
	github.com/googollee/go-socket.io v1.4.4
	github.com/jessevdk/go-flags v1.4.1-0.20200711081900-c17162fe8fd7
	github.com/jrick/logrotate v1.0.0
	github.com/planetdecred/pdapi/db/dcrpg/v6 v6.0.0
	github.com/planetdecred/pdapi/exchanges/v3 v3.0.0
	github.com/planetdecred/pdapi/gov/v4 v4.0.0
	github.com/planetdecred/pdapi/v6 v6.0.0
	github.com/rs/cors v1.7.0
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
)
