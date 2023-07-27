# Project Teaching Management
Author: Tran Quang Kiet
<br>
Date start: 25/07/2023

## Description
**Quản lí giáo vụ:**
1. Các đối tượng trong hệ thống: Quản trị viên, sinh viên, giảng viên
2. Tính năng: 
+ Đăng kí tài khoản người dùng (Có 3 loại tài khoản: tài khoản sinh viên, tài khoản giảng viên, tài khoản quản trị viên)
+ Đăng nhập để sử dụng các tính năng trong hệ thống
+ Quản lí thông tin giảng viên: 
       . Xem thông tin giảng viên
       . Sửa thông tin giảng viên
       . Xem thông tin lớp, môn học mà giảng viên dạy
       . Nhập điểm thi cho sinh viên
       . Đăng kí lớp dạy
+ Quản lí thông tin sinh viên: 
       . Xem thông tin sinh viên
       . Sửa thông tin sinh viên
       . Xem thông tin lớp học, môn học mà sinh viên đăng kí
       . Đăng kí lớp học
       . Tra cứu điểm số
       . Xem lịch học
+ Quản lí thông tin quản trị viên:
       . Thêm lớp học
       . Xem danh sách giảng viên
       . Xem danh sách học viên
       . Xóa giảng viên, sinh viên
       . Sửa điểm
3. Cơ sở dữ liệu:
+ TAIKHOAN (ID, USERNAME, PASSWORD, ROLE)
+ HOCVIEN (MAHV, HO, TEN, NGSINH, GIOITINH, NOISINH, MALOP,DIACHI)
  Tân từ: mỗi học viên phân biệt với nhau bằng mã học viên, lưu trữ họ tên, ngày sinh, giới tính, nơi sinh, thuộc
  lớp nào.
+ LOP (MALOP, TENLOP, TRGLOP, SISO, MAGVCN,MAKHOA)
  Tân từ: mỗi lớp gồm có mã lớp, tên lớp, học viên làm lớp trưởng của lớp, sỉ số lớp và giáo viên chủ nhiệm.
+ KHOA (MAKHOA, TENKHOA, NGTLAP, TRGKHOA)
  Tân từ: mỗi khoa cần lưu trữ mã khoa, tên khoa, ngày thành lập khoa và trưởng khoa (cũng là một giáo viên
  thuộc khoa).
+ MONHOC (MAMH, TENMH, TCLT, TCTH, MAKHOA)
  Tân từ: mỗi môn học cần lưu trữ tên môn học, số tín chỉ lý thuyết, số tín chỉ thực hành và khoa nào phụ trách.
+ DIEUKIEN (MAMH, MAMH_TRUOC)
  Tân từ: có những môn học học viên phải có kiến thức từ một số môn học trước.
+ GIAOVIEN (MAGV, HOTEN, HOCVI,HOCHAM,GIOITINH, NGSINH, NGVL,HESO, MUCLUONG, MAKHOA)
  Tân từ: mã giáo viên để phân biệt giữa các giáo viên, cần lưu trữ họ tên, học vị, học hàm, giới tính, ngày sinh,
  ngày vào làm, hệ số, mức lương và thuộc một khoa.
+ GIANGDAY (MALOP, MAMH, MAGV, HOCKY, NAM, TUNGAY, DENNGAY)
  Tân từ: mỗi học kỳ của năm học sẽ phân công giảng dạy: lớp nào học môn gì do giáo viên nào phụ trách.
+ KETQUATHI (MAHV, MAMH, LANTHI, NGTHI, DIEM, KQUA)
  Tân từ: lưu trữ kết quả thi của học viên: học viên nào thi môn học gì, lần thi thứ mấy, ngày thi là ngày nào,
  điểm thi bao nhiêu và kết quả là đạt hay không đạt.