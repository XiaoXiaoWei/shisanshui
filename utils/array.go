package utils

//ArrayQuery 查询d是否在s存在
func ArrayQuery(s, d []int) bool {
	if len(s) < len(d) {
		return false
	}
	s = CopyIntArray(s)
	d = CopyIntArray(d)
	for i := 0; i < len(s); i++ {
		if s[i] == d[0] {
			if len(d) > 1 {
				s = append(s[:i], s[i+1:]...)
				d = d[1:]
				return ArrayQuery(s, d)
			} else {
				return true
			}
		} else if i+1 == len(s) {
			return false
		}
	}
	return true
}

//CopyIntArray 复制数组
func CopyIntArray(s []int) (r []int) {
	r = make([]int, len(s))
	copy(r, s)
	return
}

//DelArrayMember 删除数组成员
func DelArrayMember(s, d []int, n int) []int {
	if len(d) == n {
		return s
	}
	if len(d) != 0 {
		for i := 0; i < len(s); i++ {
			if s[i] == d[n] {
				s = append(s[:i], s[i+1:]...)
				n++
				return DelArrayMember(s, d, n)
			}
		}
	}
	return s
}

//DelArrayMembers 删除数组成员数
func DelArrayMembers(s, d []int) (ret []int) {
	if len(d) == 0 {
		return s
	}
	c := d[0]
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			ret = append(s[:i], s[i+1:]...)
			d = append(d[1:])
			return DelArrayMembers(ret, d)
		}
	}
	d = append(d[1:])
	return DelArrayMembers(s, d)
}

//SetArrayAllMember 设置数组全部成员
func SetArrayAllMember(s []int, n int) {
	for i := 0; i < len(s); i++ {
		s[i] = n
	}
}

//ArrayMuddled 乱序
func ArrayMuddled(s []int) {
	if len(s) < 2 {
		return
	}
	for i := 0; i < len(s); i++ {
		j := Rands.GetInt(len(s) - 1)
		s[i%len(s)], s[j] = s[j], s[i%len(s)]
	}
}
