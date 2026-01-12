package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	// 字符串
	strs := "hello world hello"
	fmt.Println(len(strs))
	strs1 := "您好 世界"
	fmt.Println(utf8.RuneCountInString(strs1))
	for i, ch := range strs {
		fmt.Println(i, string(ch))
	}
	for i, ch := range strs1 {
		fmt.Println(i, string(ch))
	}

	arr := []rune(strs)
	arr1 := []rune(strs1)
	fmt.Println(string(arr), string(arr1))

	a := "您"
	b := "好"
	fmt.Println(fmt.Sprintf("%v %v", a, b))

	part := strings.Split(strs, " ")
	fmt.Println(part[0])

	fmt.Println(strings.Join(part, "-"))

	fmt.Println(strings.HasPrefix(strs, "hello"))
	fmt.Println(strings.HasSuffix(strs1, "世界"))
	fmt.Println(strings.Contains(strs, "hello"))
	fmt.Println(strings.Replace(strs, "hello", "hi", 2))
	re := regexp.MustCompile("hello")
	fmt.Println(re.ReplaceAllString(strs, "hi"))

	str2 := `dhsah`
	fmt.Println(str2)

	fmt.Println(strings.TrimSpace(strs))
	fmt.Println(strings.Trim(strs, " "))

	// 数组
	nums := []int{1, 2, 3, 4, 5}
	nums = append(nums, 7, 6)
	fmt.Println(nums)

	fmt.Println(append(nums[:2], nums[3:]...))

	nums1 := make([]int, len(nums))
	copy(nums1, nums)
	fmt.Println(nums1)

	sub := nums1[1:2]
	fmt.Println(sub)

	for i, v := range nums1 {
		fmt.Println(i, v)
	}

	sort.Ints(nums)
	fmt.Println(nums)

	strArr := []string{"hello", "a", "b", "c"}
	fmt.Println(strArr)
	sort.Strings(strArr)
	fmt.Println(strArr)

	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	fmt.Println(strArr)

	// map
	m1 := make(map[string]int)
	m2 := map[string]int{
		"a": 1,
		"c": 2,
	}
	fmt.Println(m1, m2)

	m1["a"] = 10
	m2["a"] = 10
	fmt.Println(m1, m2)

	v, ok := m1["a"]
	if !ok {
		log.Fatal("error happen")
	}
	fmt.Println(v)
	delete(m2, "a")
	fmt.Println(m2)

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println(len(m2))

	m := map[string]int{
		"c": 3,
		"b": 1,
		"a": 2,
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, val := range keys {
		fmt.Println(val, m[val])
	}

	// value排序
	type KV struct {
		Key   string
		Value int
	}

	list := make([]KV, 0, len(m))
	for k, v := range m {
		list = append(list, KV{
			Key:   k,
			Value: v,
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Value < list[j].Value
	})
	for _, kv := range list {
		fmt.Println(kv.Key, kv.Value)
	}

	m3 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m4 := map[string]int{
		"a": 1,
		"b": 3,
	}
	fmt.Println(mapEqual(m3, m4))

	m5 := map[string][]int{
		"a": []int{1, 2, 3},
		"b": []int{1},
	}

	m6 := map[string][]int{
		"a": []int{1, 2, 3},
		"b": []int{1, 2},
	}
	fmt.Println(mapDeepEqual(m5, m6))

	fmt.Println(reflect.DeepEqual([]int{1, 23, 3}, []int{4, 5, 6}))

	fmt.Println((0 + 8) % 7)

	math.Max(1.0, 2)
}

func mapEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2k, ok := m2[k]; !ok || m2k != v {
			return false
		}
	}
	return true
}

func mapDeepEqual(m1, m2 map[string][]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		m2k, ok := m1[k]
		if !ok || !reflect.DeepEqual(m2k, v) {
			return false
		}
	}
	return true
}
