package processors

import (
	"github.com/stretchr/testify/mock"

	"github.com/AnneNamuli/go-stellar/ingest"
)

var _ ChangeProcessor = (*MockChangeProcessor)(nil)

type MockChangeProcessor struct {
	mock.Mock
}

func (m *MockChangeProcessor) ProcessChange(change ingest.Change) error {
	args := m.Called(change)
	return args.Error(0)
}
