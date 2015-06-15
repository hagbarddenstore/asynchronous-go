package scrapers

type stringArray []string

// Contains determines if s contains item.
func (s stringArray) Contains(item string) bool {
	for _, arrayItem := range s {
		if item == arrayItem {
			return true
		}
	}

	return false
}

// Append an item to s.
func (s stringArray) Append(item string) {
	s = append(s, item)
}
