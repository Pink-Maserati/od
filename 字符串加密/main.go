package main

// 加密字符串

// 字符串中没有数字，全是小写字母，对该二进制文件进行压缩。

// 规则1：连续相同的字母，使用数字来表示重复出现的个数，bbbccc压缩成b3c3，但是如果只有两个连续的字母，压缩之后没有收益，不压缩，如bb就不再压缩:
// 规则2：对于重复出现的连续子串，压缩为大写子串加重复数字，例如abcabcabc压缩成ABC3;
// 规则3：重复字母和重复子串同时出现，优先进行重复子串的压缩，例如aaaabcabc,压缩成a3ABC2;
// 规则4：有多个重复子串时，优先压缩最长的重复子串，例如aaaabcabcaaaabcabc,压缩成AAAABCABC2;
// 输入：

// aaaabcaabcaabcaabc
// 输出：

// aaAABC4

func main() {

}
