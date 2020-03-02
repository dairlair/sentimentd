package training

type TextServiceCache struct {
	service TokenServiceInterface
	cache   map[int64]map[string]int64
}

func NewTextServiceCache(service TokenServiceInterface) TextServiceCache {
	return TextServiceCache{
		service: service,
		cache: make(map[int64]map[string]int64),
	}
}

func (c TextServiceCache) FindOrCreate(brainID int64, text string) int64 {
	if _, ok := c.cache[brainID]; !ok {
		c.cache[brainID] = make(map[string]int64)
	}

	id, ok := c.cache[brainID][text]
	if !ok {
		id = c.service.FindOrCreate(brainID, text)
		c.cache[brainID][text] = id
	}

	return id
}