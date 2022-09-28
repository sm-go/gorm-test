# Testing the Go ORM (GROM)

 - This test is from it official site

gorm ဆိုတာက Laravel ရဲ့ Elequent နဲ့ သွားပြီး အလားသဏ္ဍာန်တူတယ်

သူ့ အလိုလို migration ဖြစ်သွားတယ်။ table ရဲ့ structure က 
type User struct {  // ဒီ structure အတိုင်းပဲ
	gorm.Model
	Name    string
	Email   string
}
db.AutoMigration(&User) <<< users ဆိုတဲ့ table နာမည်နဲ့ တည်ဆောက်သွားတယ်။