package hashtable

//哈希表

const expandFactor = 0.75

type HashMap struct {
	m   []*KeyPairs
	cap int
	len int
}

type KeyPairs struct {
	key   string
	value interface{}
	next  *KeyPairs
}

func NewHashMap(cap int) *HashMap {
	if cap < 16 {
		cap = 16
	}
	HashMap := new(HashMap)
	HashMap.m = make([]*KeyPairs,cap)
	return HashMap
}

func (h *HashMap) Index(key string) int {
	return BKDRHash(key, h.cap)
}

func BKDRHash(str string,cap int) int {
	seed := int(131)
	hash := int(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + int(str[i])
	}
	return hash%cap
}

func (h *HashMap) Put(key string, value interface{}) {
	index := h.Index(key)
	element := h.m[index]
	if element == nil {
		h.m[index] = &KeyPairs{key, value, nil}
	} else {
		for element.next != nil {
			if element.key == key {
				element.value = value
				return
			}
			element = element.next
		}
		element.next = &KeyPairs{key, value, nil}
	}
	h.len++
	if (float64(h.len) / float64(h.cap)) > expandFactor {
		newH := NewHashMap(2 * h.cap)
		for _, pairs := range h.m {
			for pairs != nil {
				newH.Put(pairs.key, pairs.value)
			}
		}
		h.cap = newH.cap
		h.m = newH.m
	}
}

func (h *HashMap) Get(key string) interface{} {
	index := h.Index(key)
	pairs := h.m[index]
	if pairs.key == key {
		return pairs.value
	}
	for pairs.next != nil {
		if pairs.key == key {
			return pairs.value
		}
		pairs = pairs.next
	}
	return nil
}
