package main

func pageDivision(p int)(I1 []int, B1 []BlogsInfo){
	db := connect()
	var B []BlogsInfo
	var b BlogsInfo
	var i int
	var I []int
	result := db.Find(&b)
	q := int(result.RowsAffected)
	result = db.Limit(10).Offset((p - 1) * 10).Find(&b)
	rows, _ := result.Rows()
	if p==0{p=1}
	if p*10-q > 0 {
		for j := 0; j < 10-(p*10-q); j++ {
			rows.Next()
			db.ScanRows(rows, &b)
			B = append(B, b)
		}
	} else {
		for j := 0; j < 10; j++ {
			rows.Next()
			db.ScanRows(rows, &b)
			B = append(B, b)
		}
	}
	if q > 60 {
		if p <= 3 {
			p = 4
		}
		if p > q/10-3 && q%10==0{
			p = q/10 - 3
		}
		if p > q/10-3 && q%10!=0{
			p = q/10 - 2
		}
		for i = p - 3; i < p+4; i++ {
			I = append(I, i)
		}
	} else {
		if q%10 == 0 {
			for i = 1; i <= q/10; i++ {
				I = append(I, i)
			}
		} else {
			for i = 1; i <= q/10+1; i++ {
				I = append(I, i)
			}
		}
	}
	return I, B
}
