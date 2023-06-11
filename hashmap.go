package hashmap

import "errors"

const (
	FIXED_SIZE        = 100
	DELETED_NODE_SIGN = "!"
)

type Node struct {
	Key   string
	Value string
}

type HashMap struct {
	Hmap [FIXED_SIZE]Node
}

func HashFunc(str string) int {
	hash := 0
	for _, character := range str {
		hash += int(character)
	}
	return hash % FIXED_SIZE
}

func (hashmap HashMap) Get(key string) (int, Node, error) {
	hash := HashFunc(key)

	for i := hash; i < FIXED_SIZE+hash; i++ {
		index := i % FIXED_SIZE
		if hashmap.Hmap[index].Key == key {
			return index, hashmap.Hmap[index], nil
		}
		if hashmap.Hmap[index].Key == "" {
			return -1, Node{}, errors.New("key does not exist")
		}
	}

	return -1, Node{}, errors.New("key does not exist")
}

func (hashmap *HashMap) Insert(key string, value string) error {
	hash := HashFunc(key)
	_, _, err := hashmap.Get(key)
	if err == nil {
		return errors.New("key already exist")
	}
	for i := hash; i < FIXED_SIZE+hash; i++ {
		index := i % FIXED_SIZE
		if hashmap.Hmap[index].Key == "" || hashmap.Hmap[index].Key == DELETED_NODE_SIGN {
			hashmap.Hmap[index] = Node{Key: key, Value: value}
			return nil
		}
	}
	return errors.New("hash table is full :(")
}

func (hashmap *HashMap) Delete(key string) error {
	index, _, err := hashmap.Get(key)
	if err != nil {
		return errors.New("key does not exist")
	}

	hashmap.Hmap[index].Key = DELETED_NODE_SIGN
	hashmap.Hmap[index].Value = ""
	return nil
}
