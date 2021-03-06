// Package reap contains the history reaping subsystem for horizon.  This system
// is designed to remove data from the history database such that it does not
// grow indefinitely.  The system can be configured with a number of ledgers to
// maintain at a minimum.
package reap

import (
	"time"

	"github.com/AnneNamuli/go-stellar/services/horizon/internal/db2/history"
	"github.com/AnneNamuli/go-stellar/services/horizon/internal/ledger"
	"github.com/AnneNamuli/go-stellar/support/db"
)

// System represents the history reaping subsystem of horizon.
type System struct {
	HistoryQ       *history.Q
	RetentionCount uint
	ledgerState    *ledger.State

	nextRun time.Time
}

// New initializes the reaper, causing it to begin polling the stellar-core
// database for now ledgers and ingesting data into the horizon database.
func New(retention uint, dbSession *db.Session, ledgerState *ledger.State) *System {
	r := &System{
		HistoryQ:       &history.Q{dbSession},
		RetentionCount: retention,
		ledgerState:    ledgerState,
	}

	r.nextRun = time.Now().Add(1 * time.Hour)
	return r
}
