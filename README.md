# bot1
Trada Academy bot1

- vào trong $GOPATH.
- tạo folder takedatabase nằm trong $GOTPATH/src/
- import take.go vào folder takedatabase
- xóa folder anotherpackage
- set biến os = botToken

- go build tradabot.go
- go run tradabot.go --> to run program
- lưu ý: 
  + take.go đã set database mặc định : oop, schema: student, table: column : idstudent, name, more_info, lesson,
  --> myshost,host: 127.0.0.1:@tcp3306
  + chỉ có thể chạy sau khi tạo database(easy to config) và setup package (takedatabase) theo đúng nguyên mẫu.

  + chức năng /list, /export sẽ sớm được hoàn thiện.
