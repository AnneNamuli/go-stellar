package ticker

import (
	"github.com/AnneNamuli/go-stellar/services/ticker/internal/gql"
	"github.com/AnneNamuli/go-stellar/services/ticker/internal/tickerdb"
	hlog "github.com/AnneNamuli/go-stellar/support/log"
)

func StartGraphQLServer(s *tickerdb.TickerSession, l *hlog.Entry, port string) {
	graphql := gql.New(s, l)

	graphql.Serve(port)
}
