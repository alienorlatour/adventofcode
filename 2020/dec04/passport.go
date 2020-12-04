package dec04

func newPassport() passport {
	return make(map[string]string, 0)
}

type passport map[string]string

func (p passport) IsValid() bool {
	mandatoryKeys := []string{
		"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
	}
	for _, k := range mandatoryKeys {
		_, ok := p[k]
		if !ok {
			return false
		}
	}
	return true
}

func (p passport) Reset() {
	for k := range p {
		delete(p, k)
	}
}

func (p passport) Add(k, v string) {
	p[k] = v
}
