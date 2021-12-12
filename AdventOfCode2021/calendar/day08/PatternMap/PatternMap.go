package PatternMap

type Pattern struct {
	mp map[int]int
}

func New() Pattern {
	return Pattern{
		map[int]int{},
	}
}

func (p *Pattern) Add(key string, value int) {
	hash := hashString(key)

	p.mp[hash] = value
}

func (p *Pattern) GetValue(key string) int {
	hash := hashString(key)

	return p.mp[hash]
}

func hashString(s string) int {
	sum := 0

	for _, l := range s {
		switch l {
		case 'a':
			sum += 223
			break
		case 'b':
			sum += 31
			break
		case 'c':
			sum += 59
			break
		case 'd':
			sum += 7
			break
		case 'e':
			sum += 19
			break
		case 'f':
			sum += 111
			break
		case 'g':
			sum += 13
			break
		}
	}

	return sum
}
