package backtraking

// 电话号码的字母组合
// 力扣：https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
//
// 参考：

// 时间复杂度：O(n * 4^n)
// 空间复杂度：O(n)

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res   = []string{}
	final = 0
)

// 解法一 DFS
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
	return
}

// 解法二 非递归
func letterCombinations_(digits string) []string {
	if digits == "" {
		return []string{}
	}
	index := digits[0] - '0'
	letter := letterMap[index]
	tmp := []string{}
	for i := 0; i < len(letter); i++ {
		if len(res) == 0 {
			res = append(res, "")
		}
		for j := 0; j < len(res); j++ {
			tmp = append(tmp, res[j]+string(letter[i]))
		}
	}
	res = tmp
	final++
	letterCombinations(digits[1:])
	final--
	if final == 0 {
		tmp = res
		res = []string{}
	}
	return tmp
}

// 解法三 回溯（参考回溯模板，类似DFS）
var result []string
var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinationsBT(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	letterFunc("", digits)
	return result
}

func letterFunc(path string, digits string) {
	// string是不可变的。这意味着一旦一个字符串被创建，它的内容就不能被修改。任何对字符串的操作，如拼接、切片或替换，都会生成一个新的字符串。
	// 所以不用copy
	if digits == "" { // string是
		result = append(result, path)
		return
	}
	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		path += dict[k][i]
		letterFunc(path, digits)
		path = path[0 : len(path)-1]
	}
}

func letterCombinationsBT2(digits string) []string {
	_result := []string{}
	if digits == "" {
		return _result
	}
	dfs("", digits, &_result)
	return _result
}

func dfs(path string, digits string, result *[]string) {
	if digits == "" {
		*result = append(*result, path)
		return
	}

	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		path += dict[k][i]
		dfs(path, digits, result)
		path = path[:len(path)-1]
	}
}
