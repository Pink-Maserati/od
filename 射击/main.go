package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
  给定一个射击比赛成绩单
  包含多个选手若干次射击的成绩分数
  请对每个选手按其最高三个分数之和进行降序排名
  输出降序排名后的选手id序列
  条件如下
    1. 一个选手可以有多个射击成绩的分数，且次序不固定
    2. 如果一个选手成绩少于3个，则认为选手的所有成绩无效，排名忽略该选手
    3. 如果选手的成绩之和相等，则相等的选手按照其id降序排列

   输入描述:
     输入第一行
         一个整数N
         表示该场比赛总共进行了N次射击
         产生N个成绩分数
         2<=N<=100

     输入第二行
       一个长度为N整数序列
       表示参与每次射击的选手id
       0<=id<=99

     输入第三行
        一个长度为N整数序列
        表示参与每次射击选手对应的成绩
        0<=成绩<=100

   输出描述:
      符合题设条件的降序排名后的选手ID序列

   示例一
      输入:
        13
        3,3,7,4,4,4,4,7,7,3,5,5,5
        53,80,68,24,39,76,66,16,100,55,53,80,55
      输出:
        5,3,7,4
      说明:
        该场射击比赛进行了13次
        参赛的选手为{3,4,5,7}
        3号选手成绩53,80,55 最高三个成绩的和为188
        4号选手成绩24,39,76,66  最高三个成绩的和为181
        5号选手成绩53,80,55  最高三个成绩的和为188
        7号选手成绩68,16,100  最高三个成绩的和为184
        比较各个选手最高3个成绩的和
        有3号=5号>7号>4号
        由于3号和5号成绩相等  且id 5>3
        所以输出5,3,7,4
*/

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var idStr string
	fmt.Scanf("%s", &idStr)
	ids := strings.Split(idStr, ",")
	var scoreStr string
	fmt.Scanf("%s", &scoreStr)
	scores := strings.Split(scoreStr, ",")
	idScoresMap := make(map[string][]string)
	for i := 0; i < n; i++ {
		id := ids[i]
		score := scores[i]
		if _, ok := idScoresMap[id]; ok {
			temp := idScoresMap[id]
			temp = append(temp, score)
			idScoresMap[id] = temp
		} else {
			temp := []string{}
			temp = append(temp, score)
			idScoresMap[id] = temp
		}

	}

	idScoress := idScores{}

	for k, v := range idScoresMap {
		sum := 0

		tmpv := []int{}
		for _, v1 := range v {
			tmp, _ := strconv.Atoi(v1)
			tmpv = append(tmpv, tmp)
		}
		sort.Ints(tmpv)
		count := 0
		for i := len(tmpv) - 1; i >= 0; i-- {

			if count < 3 {
				sum += tmpv[i]
				count++
			} else {
				break
			}

		}
		ktmp, _ := strconv.Atoi(k)
		idScoreTemp := idScore{ktmp, sum}
		idScoress = append(idScoress, idScoreTemp)
	}

	fmt.Println(idScoress)
	sort.Sort(idScoress)
	for _, v := range idScoress {
		fmt.Println(v.id)

	}

}

type idScore struct {
	id  int
	sum int
}

type idScores []idScore

func (idSort idScores) Len() int {
	return len(idSort)
}

func (idSort idScores) Less(i, j int) bool {
	if idSort[i].id != idSort[j].id && idSort[i].sum > idSort[j].sum {
		return true
	} else if idSort[i].sum == idSort[j].sum && idSort[i].id > idSort[j].id {
		return true
	} else {
		return false
	}

}

func (idSort idScores) Swap(i, j int) {
	idSort[i], idSort[j] = idSort[j], idSort[i]

}
