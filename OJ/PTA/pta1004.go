package pta

import (
"fmt"
)

type Student struct {
	Name  string
	StuID string
	Score int
}

func GradeSort() {
	var (
		n, i, max, min, maxID, minID int
	)
	fmt.Scanln(&n)
	stu := make([]Student, n)
	for i = 0; i < n; {
		if i == 0 {
			fmt.Scanln(&stu[i].Name, &stu[i].StuID, &stu[i].Score)
			if len(stu[i].Name) <= 10 && len(stu[i].StuID) <= 10 && stu[i].Score >= 0 && stu[i].Score <= 100 {
				max = stu[i].Score
				min = stu[i].Score
				maxID = i
				minID = i
				i++
			} else {
				return
			}
		} else {
			fmt.Scanln(&stu[i].Name, &stu[i].StuID, &stu[i].Score)
			if len(stu[i].Name) <= 10 && len(stu[i].StuID) <= 10 && stu[i].Score >= 0 && stu[i].Score <= 100 {
				if min > stu[i].Score {
					min = stu[i].Score
					minID = i
				}
				if max < stu[i].Score {
					max = stu[i].Score
					maxID = i
				}
				i++
			}
		}
	}
	fmt.Println(stu[maxID].Name + " " + stu[maxID].StuID)
	fmt.Println(stu[minID].Name + " " + stu[minID].StuID)
}

