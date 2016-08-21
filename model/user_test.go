package model

import (
	"testing"
	"myStep/form"
)

// 英大文字、英小文字、数字、記号
const TEST_PASSWORD = "Tar0!?"

/****************************
 * 1（成功）
 * 2（失敗）パスワード不一致
 * 3（失敗）ユーザー名不一致
 ****************************/
func TestAuth(t *testing.T) {
	initTestUserData()
	testAuth1(t)
	testAuth2(t)
	testAuth3(t)
	deleteTestUserData()
}

func testAuth1(t *testing.T) {
	form := form.S01Form {
		UserName: "Taro",
		Password: TEST_PASSWORD,
	}
	_, err := Auth(&form)

	if err != nil {t.Errorf("testAuth1 Error")}
}

func testAuth2(t *testing.T) {
	form := validation.S01Form {
		UserName: "Taro",
		Password: TEST_PASSWORD + "a",
	}
	_, err := Auth(&form)

	if err == nil {t.Errorf("testAuth2 Error")}
}

func testAuth3(t *testing.T) {
	form := validation.S01Form {
		UserName: "Saburo",
		Password: TEST_PASSWORD,
	}
	_, err := Auth(&form)

	if err == nil {t.Errorf("testAuth3 Error")}
}

func initTestUserData() {
	user := User{
		UserName:  "Taro",
		Password:  PasswordHash(TEST_PASSWORD),
	}
	db.Debug().Create(&user)
	user = User{
		UserName:  "Jiro",
		Password:  PasswordHash(TEST_PASSWORD),
	}
	db.Debug().Create(&user)
}

func deleteTestUserData() {
	user := User{
		UserName:  "Taro",
	}
	db.Debug().Unscoped().Model(&User{}).Where("user_name = ?", user.UserName).Delete(&User{})
	user = User{
		UserName:  "Jiro",
	}
	db.Debug().Unscoped().Model(&User{}).Where("user_name = ?", user.UserName).Delete(&User{})
}