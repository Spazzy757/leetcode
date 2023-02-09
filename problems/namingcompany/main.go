/*
* https://leetcode.com/problems/naming-a-company/
 */

package namingcompany

func distinctNames(ideas []string) int64 {
	ans := 0

	initialGroup := make(map[string]map[string]bool)
	foo := "abcdefghijklmnopqrstuvwxyz"
	for _, v := range ideas {
		_, ok := initialGroup[v[:1]]
		if ok {
			initialGroup[v[:1]][v[1:]] = true
		} else {
			initialGroup[v[:1]] = map[string]bool{v[1:]: true}
		}
	}

	for i := 0; i < 26; i++ {
		for j := i + 1; j < 26; j++ {
			numOfMutual := 0

			for v, _ := range initialGroup[string(foo[i])] {
				if initialGroup[string(foo[j])][v] {
					numOfMutual++
				}
			}
			ans += (len(initialGroup[string(foo[i])]) - numOfMutual) * (len(initialGroup[string(foo[j])]) - numOfMutual)
		}
	}
	return int64(2 * ans)
}
