package ephemeral

import (
	"testing"

	eh "github.com/vercly/eventhorizon"
	"github.com/vercly/eventhorizon/mocks"
)

func TestInnerHandler(t *testing.T) {
	m := NewMiddleware()
	h := m(mocks.NewEventHandler("test"))
	_, ok := h.(eh.EventHandlerChain)
	if !ok {
		t.Error("handler is not an EventHandlerChain")
	}
}
