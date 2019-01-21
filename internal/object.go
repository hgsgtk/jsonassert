package internal

func (a *asserter) checkObject(level string, act, exp map[string]interface{}) {
	if len(act) != len(exp) {
		a.printer.Errorf("different number of keys at level '%s' in actual JSON (%d) and expected JSON (%d)", level, len(act), len(exp))
	}
	if unique := difference(act, exp); len(unique) != 0 {
		a.printer.Errorf("at level '%s', key(s) %+v present in actual JSON but not in expected JSON", level, unique)
	}
	if unique := difference(exp, act); len(unique) != 0 {
		a.printer.Errorf("at level '%s', key(s) %+v present in expected JSON but not in actual JSON", level, unique)
	}
	for key := range act {
		if contains(exp, key) {
			a.Assert(level+"."+key, serialize(act[key]), serialize(exp[key]))
		}
	}
}

func difference(act, exp map[string]interface{}) []string {
	unique := []string{}
	for key := range act {
		if !contains(exp, key) {
			unique = append(unique, key)
		}
	}
	return unique
}

func contains(container map[string]interface{}, candidate string) bool {
	for _, key := range container {
		if key == candidate {
			return true
		}
	}
	return false
}