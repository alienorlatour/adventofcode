package dec04

func newPassport() passport {
	return make(map[string]string, 0)
}

type passport map[string]string

func (p passport) IsValid() bool {
	for k, validate := range mandatoryKeys {
		value, ok := p[k]
		if !ok {
			// missing entry
			return false
		}
		if !validate(value) {
			// invalid entry
			return false
		}
	}
	return true
}

// Reset empties the passport for reuse.
func (p passport) Reset() {
	for k := range p {
		delete(p, k)
	}
}

// Add a key/value entry
func (p passport) Add(k, v string) {
	p[k] = v
}
