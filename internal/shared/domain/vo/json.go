package vo

type JSON map[string]any

func NewJSON(value map[string]any) *JSON {
	return (*JSON)(&value)
}

func (this *JSON) Value() map[string]any {
	if this == nil {
		return nil
	}
	return *this
}
