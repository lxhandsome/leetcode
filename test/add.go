package main

func Score(score int) string {
	var res string
	switch {
	case score >= 90:
		res = "优秀"
	case score >= 80:
		res = "良好"
	case score >= 60:
		res = "及格"
	default:
		res = "不及格"
	}
	return res
}

func main() {

}
