#Chức năng mô tả
Tạm thời: mỗi người chỉ đăng kí 1 khoá cũng được. Tuy nhiên, về tương lai muốn đăng kí nhiều khoá cũng ok.
Chỉ cho đăng kí khi chat riêng (private chat) với bot. Không cho đăng kí khi chat với bot ở trong group hoặc supergroup.

Danh sách lệnh:

———
/start

Bot: Cảm ơn bạn đã liên lạc với Trada Tech. Đây là bot trả lời tự động hỗ trợ các lệnh sau.
/register - đăng kí khoá học, dịch vụ, hoặc đặt câu hỏi cho Trade Tech
/me - xem danh sách khoá học đã đăng kí
/cancel - thông báo muốn huỷ đăng kí
/help - hiện hướng dẫn.

———
/register

Bot: Vui lòng chọn.
DApp - Khoá học Ethereum DApp Development
Khác - Các yêu cầu khác

=> hiện 2 button cho user lựa chọn, khi chọn xong thì hide đi. Trên button ghi chữ “DApp” và “Khác”.

User: DApp hoặc "Khác".

(nếu mà trả lời ko đúng 2 options trên thì bảo "Vui lòng dùng 2 nút có sẵn để trả lời")

Nếu mà “Khác” thì hỏi thêm (nếu là DApp thì ko cần hỏi thêm gì):
Bot: Vui lòng cung cấp chi tiết
User: Blah blah.

Cảm ơn bạn. Nhân viên của Trada Tech sẽ liên lạc lại với bạn để hỏi thêm chi tiết.

Note 1: ko cần confirm, đăng kí nhầm kệ nó cũng được, spam cũng ko sao, vì giờ còn ít người đăng kí lắm.
Note 2: Gửi message cho admin.

Nội dung:
Có đăng kí mới!
Từ: @xxx
Khoá học: YYY
Thông tin thêm: 

——
/me
Bot: Bạn chưa đăng kí hay có câu hỏi nào.
Bot: Khoá học đã đăng kí: Khác
Thông tin thêm: blah blah

——
/cancel
Bot: bạn chưa đăng kí khoá học nào nên không cần huỷ.
Bot: yêu cầu huỷ của bạn đã được thông báo cho nhân viên Trada Tech. Chúng tôi sẽ liên lạc lại để xác nhận.

Note: ko xoá record, chỉ set Cancel: requested.
Gửi thông báo cho admin.

—-
/help
Hiện hướng dẫn giống phần start, nhưng ko cần lời chào.

————
Chức năng cho admin (cái này là web hay lệnh bot đều được, hiện tại cũng chưa cần, sau này có thời gian thì thêm sau cũng được).
/list -> hiện danh sách đăng kí (5 ông gần đây)
/export -> gửi lại admin full danh sách đăng kí (json, text, hoặc csv thế nào cũng được)


Trong trường hợp người dùng chọn khóa học DApp thì kết thúc đăng ký và gửi msg cho admin — 
Hay còn yêu cầu người dùng nhập thông tin khác sau đó mới kết thúc đăng ký a.


/info

Thi măng cụt, [22.11.18 10:51]
1. Đào tạo Ethereum Dapp Developer
2. Đào tạo theo yêu cầu riêng của công ty bạn
3. Tư vấn về phát triển và ứng dụng blockchain