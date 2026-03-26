package hook

import "fmt"

type Handler func(collection string, data map[string]interface{}) error

type HookManager struct {
  onBeforeSave []Handler
	onAfterSave  []Handler
}

func NewManager() *HookManager {
  return &HookManager{}
}

func (m *HookManager) TriggerBeforeSave(collection string, data map[string]interface{}) error {
  for _, handler := range m.onBeforeSave {
		if err := handler(collection, data); err != nil {
      return fmt.Errorf("[%s] hook validation failded: %w", collection, err)
		}
	}
	return nil
}

func (m *HookManager) RegisterBeforeSave(h Handler) {
  m.onBeforeSave = append(m.onBeforeSave, h)
}
