package hashmap_test

import (
	
	"hashmap"
	"testing"

	"github.com/google/uuid"
)

func TestHashMapGet(t *testing.T) {
	t.Run("happy flow - simple", func(t *testing.T) {
		key := "abc"
		value := "value"
		hmap := hashmap.HashMap{}
		hmap.Insert(key, value)
		_, node, _ := hmap.Get(key)
		if (node != hashmap.Node{Key: key, Value: value}) {
			t.Fail()
		}
	})

	t.Run("happy flow - when key is not located in its original hashfunc index", func(t *testing.T) {
		key := "abc"
		key1 := "cba"
		value1 := "value1"
		value := "value"
		hmap := hashmap.HashMap{}
		hmap.Insert(key1, value1)
		hmap.Insert(key, value)
		_, node, _ := hmap.Get(key)
		if (node != hashmap.Node{Key: key, Value: value}) {
			t.Fail()
		}
	})

	t.Run("return error when key is not exist", func(t *testing.T) {
		hmap := hashmap.HashMap{}
		_, _, err := hmap.Get("any value")
		if err == nil {
			t.Fail()
		}
	})
}

func TestHashMapInsert(t *testing.T) {

	type KeyValuePair struct {
		Key   string
		Value string
	}

	testCases := []struct {
		Key             string
		Value           string
		ExpectedIndex   int
		PairsToInsert   []KeyValuePair
		TestDescription string
	}{
		{Key: "abc", Value: "value", ExpectedIndex: hashmap.HashFunc("abc"), TestDescription: " simple happy flow"},
		{Key: "abc", Value: "value", ExpectedIndex: hashmap.HashFunc("abc") + 1%hashmap.FIXED_SIZE, TestDescription: "happy flow when node is taken",
			PairsToInsert: []KeyValuePair{{Key: "acb", Value: "value1"}}},
		{Key: "abc", Value: "value", ExpectedIndex: hashmap.HashFunc("abc") + 2, TestDescription: "happy flow when node is taken and the next node too",
			PairsToInsert: []KeyValuePair{{Key: "acb", Value: "value1"}, {Key: "bca", Value: "value1"}}},
	}

	for _, tc := range testCases {
		t.Run(tc.TestDescription, func(t *testing.T) {
			hmap := hashmap.HashMap{Hmap: [hashmap.FIXED_SIZE]hashmap.Node{}}
			for _, pairToInsert := range tc.PairsToInsert {
				hmap.Insert(pairToInsert.Key, pairToInsert.Value)
			}
			hmap.Insert(tc.Key, tc.Value)
			if (hmap.Hmap[tc.ExpectedIndex] != hashmap.Node{Key: tc.Key, Value: tc.Value}) {
				t.Fail()
			}
		})
	}

	t.Run("insert to deleted node", func(t *testing.T) {
		key := "abc"
		hmap := hashmap.HashMap{}
		hmap.Hmap[hashmap.HashFunc(key)] = hashmap.Node{Key: hashmap.DELETED_NODE_SIGN, Value: ""}
		hmap.Insert(key, "value")
		if (hmap.Hmap[hashmap.HashFunc(key)] != hashmap.Node{Key: key, Value: "value"}) {
			t.Fail()
		}
	})

	t.Run("insert when node is taken and next node is deleted", func(t *testing.T) {
		key := "abc"
		key1 := "cba"
		hmap := hashmap.HashMap{}
		hmap.Insert(key, "value")
		hmap.Hmap[hashmap.HashFunc(key)+1] = hashmap.Node{Key: hashmap.DELETED_NODE_SIGN, Value: ""}
		hmap.Insert(key1, "value")
		if (hmap.Hmap[hashmap.HashFunc(key1)+1] != hashmap.Node{Key: key1, Value: "value"}) {
			t.Fail()
		}
	})

	t.Run("return error when hashmap is full", func(t *testing.T) {
		hmap := hashmap.HashMap{}
		for i := 0; i < hashmap.FIXED_SIZE; i++ {
			key := uuid.New().String()
			hmap.Hmap[i] = hashmap.Node{Key: key, Value: "value"}
		}
		err := hmap.Insert("abc", "value")
		if err == nil {
			t.Fail()
		}
	})

	t.Run("return error when key is already exist", func(t *testing.T) {
		hmap := hashmap.HashMap{}
		key := "abc"
		hmap.Insert("cba", "value")
		hmap.Insert(key, "value")

		err := hmap.Insert(key, "any value")
		if err == nil {
			t.Fail()
		}
	})
}

func TestHashMapDelete(t *testing.T) {
	t.Run("happy flow", func(t *testing.T) {
		key := "abc"
		hmap := hashmap.HashMap{}
		hmap.Insert(key, "value")
		hmap.Delete(key)
		if hmap.Hmap[hashmap.HashFunc(key)].Key != hashmap.DELETED_NODE_SIGN {
			t.Fail()
		}

	})

	t.Run("return error when key does not exist", func(t *testing.T) {
		hmap := hashmap.HashMap{}
		err := hmap.Delete("any value")
		if err == nil {
			t.Fail()
		}
	})
}
