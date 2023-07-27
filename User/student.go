package User

type Student struct {
	mahv, ho, ten, ngsinh, gioitinh, noisinh, malop string
}

func (s *Student) SetStudent(student ...string) {
	s.mahv = student[0]
	s.ho = student[1]
	s.ten = student[2]
	s.ngsinh = student[3]
	s.gioitinh = student[4]
	s.noisinh = student[5]
	s.malop = student[6]
}

func (s Student) GetMahv() string {
	return s.mahv
}

func (s Student) GetHo() string {
	return s.ho
}

func (s Student) GetTen() string {
	return s.ten
}

func (s Student) GetNgsinh() string {
	return s.ngsinh
}

func (s Student) GetGioitinh() string {
	return s.gioitinh
}

func (s Student) GetNoisinh() string {
	return s.noisinh
}

func (s Student) GetMalop() string {
	return s.malop
}
