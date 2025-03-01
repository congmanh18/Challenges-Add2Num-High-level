# MyBigNumber - Ứng Dụng Console Cộng Hai Số Lớn

## Cấu trúc thư mục
```
task1/
│── bigsum/
│   ├── big_number.go        # Định nghĩa struct BigNumber và hàm Sum
│   ├── big_number_test.go   # Unit tests cho BigNumber
│
│── cmd/
│   ├── main.go               # Chương trình chính chạy ứng dụng
│
│── go.mod                    # Quản lý dependencies
│── go.sum                    # Danh sách checksum dependencies
│── README.md                 # Hướng dẫn sử dụng
```

## Cách cài đặt
Yêu cầu: Đã cài đặt [Go](https://go.dev/).

1. Clone repository về máy:
   ```sh
   git clone https://github.com/congmanh18/Challenges-Add2Num-High-level.git .
   cd task1
   ```

2. Cài đặt dependencies:
   ```sh
   go mod tidy
   ```

## Cách chạy chương trình
Chạy ứng dụng console bằng lệnh sau:
```sh
go run cmd/main.go
```
Sau đó, nhập hai số cần cộng và nhận kết quả.

### Ví dụ đầu vào/đầu ra:
```
Nhập số thứ nhất: 1234
Nhập số thứ hai: 897
Tổng của 1234 và 897 là: 2131
Bạn có muốn tiếp tục? (y/n): y
```

## Chạy Unit Tests
Để kiểm tra tính đúng đắn của thuật toán, chạy lệnh sau:
```sh
go test ./bigsum
```

## Hướng dẫn sử dụng
- Nhập hai số nguyên dương để tính tổng.
- Nhấn `y` để tiếp tục với phép tính mới.
- Nhấn `n` để thoát chương trình.

## Mô tả thuật toán
Thuật toán cộng hai số lớn mô phỏng cách cộng thủ công mà học sinh tiểu học thực hiện:

1. **Duyệt từ phải sang trái**: Bắt đầu từ chữ số cuối cùng của mỗi số.
2. **Chuyển từng ký tự sang số nguyên**: Chuyển ký tự thành số (`0-9`).
3. **Cộng từng chữ số với nhau**:
   - Cộng chữ số hiện tại của số thứ nhất với chữ số hiện tại của số thứ hai.
   - Nếu có số nhớ từ phép cộng trước đó, cộng thêm vào.
   - Nếu tổng lớn hơn hoặc bằng 10, giữ lại chữ số hàng đơn vị và ghi nhớ phần dư (`carry = tổng / 10`).
4. **Lặp lại cho đến khi hết các chữ số**.
5. **Xử lý số nhớ cuối cùng**: Nếu còn số nhớ (`carry > 0`), thêm vào kết quả.
6. **Đảo ngược kết quả**: Vì các số được tính từ phải sang trái, cần đảo ngược chuỗi kết quả để có đúng thứ tự.

Ví dụ với `sum("1234", "897")`:
```
Bước 1: 4 + 7 = 11 → Ghi 1, nhớ 1.
Bước 2: 3 + 9 = 12, cộng 1 nhớ = 13 → Ghi 3, nhớ 1.
Bước 3: 2 + 8 = 10, cộng 1 nhớ = 11 → Ghi 1, nhớ 1.
Bước 4: 1 + 0 = 1, cộng 1 nhớ = 2 → Ghi 2.
Kết quả cuối cùng: 2131.
```
Lưu ý:  Giả định giá trị tham số truyền vào hàm là đúng, chỉ chứa các kí số hợp lệ, không có kí tự nào khác.

---
