package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	slow := 0
	duplicates := make(map[byte]struct{})
	maxLen := 0

	for fast := 0; fast < len(s); fast++ {
		_, ok := duplicates[s[fast]]
		for ok {
			delete(duplicates, s[slow])
			slow++
			_, ok = duplicates[s[fast]]
		}
		duplicates[s[fast]] = struct{}{}
		maxLen = max((fast-slow)+1, maxLen)
	}

	return maxLen
}

func max(left int, right int) int {
	if left > right {
		return left
	}

	return right
}

// [p,w,w,k,e,w]
//.   ^ ^

// [a,b,c,a,b,c,b,b]
//      ^     ^
// [a,b,c,a,b,c,b,b]
//    ^   ^
// [a,b,c,a,b,c,b,b]
//      ^ ^
// [a,b,c,a,b,c,b,b]
//      ^ ^

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

// problem with train of thought:

// you successfully figured out that you
// have to increment the fastPtr to that duplicate 'a' then stop
// [a,b,c,a,b,c,b,b]
//  ^     ^

// You then successfully figured out that you had
// to shrink your window.
// [a,b,c,a,b,c,b,b]
//    ^   ^

// The only problem was, you kept shrinking it depending on
// if the value at the slow pointer was already present in the map,
// stopping the shrinking once a value at the slowPointer wasnt seen (which
// here wouldve been the fast pointers 'a')
// When we shrink our window, we do so to catch other sub arrays, but if you
// are exiting this process only when you get to no more slow pointer seens, you
// miss these sub arrays (e.g, we'd miss the sub array 'b,c,a' if that was the case)

// to catch these smaller sub arrays, you must stop incrementing the slow pointer if the
// fast pointers value was no longer seen.
