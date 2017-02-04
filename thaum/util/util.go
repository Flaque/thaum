package util

func Keys(m map[string]string) []string{
  keys := make([]string, 0, len(m))
  for k := range m {
    keys = append(keys, k)
  }
  return keys
}

func AddStringsToSet(strs []string, set map[string]string) map[string]string {
	for _, s := range strs {
		set[s] = ""
	}
	return set
}
