package auth

import (
	"app/typefile"
	"app/utility"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// 指定のEmailアドレスが登録されているか確認するための関数
func isUserRegisteredByEmail(email string) bool {
	var uid string
	query := `SELECT id FROM users WHERE email = ?;`
	err := utility.Db.QueryRow(query, email).Scan(&uid)
	if err != nil {
		return false
	}
	return true
}

// ユーザー EmailとPasswordのバリデーション
func ValidateUser(user *typefile.User) error {
	var reg string

	if user.Email == "" || user.Password == "" {
		return errors.New("Required Field Error")
	}

	//  Email バリデーション　正規表現
	reg = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailCheck := checkReg(reg, user.Email)

	if !emailCheck {
		return errors.New("Email Format Error")
	}

	// Password バリデーション　正規表現
	validLen := len(user.Password) >= 8
	hasUpper := checkReg(`[A-Z]`, user.Password)
	hasLower := checkReg(`[a-z]`, user.Password)
	hasPunc := checkReg(`[\p{P}\p{S}]`, user.Password)

	passwordCheck := validLen && hasLower && hasUpper && hasPunc

	if !passwordCheck {
		return errors.New("PassWord Format Error")
	}

	return nil
}

// 正規表現マッチング確認
func checkReg(reg string, str string) bool {
	r := regexp.MustCompile(reg)
	match := r.Match([]byte(str))
	return match
}

// Emailを指定して引数をとして渡した構造体userにデータをマッピングする関数
func getUserByEmail(user *typefile.User, email string) error {
	query := `SELECT * FROM users WHERE email = ?;`
	err := utility.Db.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return err
	}
	return nil
}

// IDを指定して引数をとして渡した構造体userにデータをマッピングする関数
func getUserById(user *typefile.User, id string) error {
	query := `SELECT * FROM users WHERE id = ?;`
	err := utility.Db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return err
	}
	return nil
}

// ユーザー登録処理　必須項目：name email password
func Register(c echo.Context) error {
	user := typefile.User{}
	user.Id = uuid.New().String()

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if isUserRegisteredByEmail(user.Email) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "User already exists",
		})
	}

	err := ValidateUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	user.Password = string(hash)

	fmt.Println(user)
	query := `INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?);`

	_, err = utility.Db.Exec(query, user.Id, user.Name, user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	// パスワードのマスキング
	user.Password = "********"

	return c.JSON(http.StatusOK, user)

}

// ログイン処理　emailとpassword指定でJWTtokenを返す
func Login(c echo.Context) error {

	user := typefile.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	regUser := typefile.User{}
	if !isUserRegisteredByEmail(user.Email) {
		// userが見つからない場合
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "User does not exist",
		})
	}

	// userが見つかった場合
	getUserByEmail(&regUser, user.Email)

	if err := bcrypt.CompareHashAndPassword([]byte(regUser.Password), []byte(user.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Password do not match!",
			})
		}
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Bad Request",
		})
	}

	accToken := CreateJwtToken(regUser.Id)

	return c.JSON(http.StatusOK, echo.Map{
		"accToken": accToken,
	})
}

// ユーザー情報更新処理
// TODO:EmailとPasswordのバリデーションを追加する
// TODO: Passwordが変更されている時はハッシュ化して保存する
func Update(c echo.Context) error {
	claims := GetClaims(c)
	id := claims.ID
	newUser := typefile.User{}
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	userMap := newUser.CreateUserMap()
	query := `UPDATE users SET`
	var nList []string
	var vList []any
	for fieldName, value := range userMap {
		if value != "" {
			nList = append(nList, fieldName)
			vList = append(vList, value)
		}
	}
	vList = append(vList, id)

	t := strings.Join(nList, "=?, ") + "=?"
	query += " " + t + " WHERE id = ?;"
	_, err := utility.Db.Exec(query, vList...)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	updateUser := typefile.User{}
	getUserById(&updateUser, id)

	return c.JSON(http.StatusOK, updateUser)

}

func Delete(c echo.Context) error {
	query := `DELETE FROM users WHERE id = ?;`
	claims := GetClaims(c)
	id := claims.ID
	res, err := utility.Db.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	deleteNum, err := res.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if deleteNum == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "delete failed",
		})
	} else if deleteNum == 1 {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "delete success",
		})
	}

	return c.JSON(http.StatusBadRequest, echo.Map{
		"error": "delete failed",
	})
}

func GetAuthUser(c echo.Context) error {
	claims := GetClaims(c)
	id := claims.ID
	user := typefile.User{}
	user.Password = "********"
	query := `SELECT id, name, email FROM users WHERE id = ?;`
	err := utility.Db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, "user not founded")
	}

	return c.JSON(http.StatusOK, user)
}
