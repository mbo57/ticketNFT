package auth

import (
	"app/typefile"
	"app/utility"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func isUserRegisteredByEmail(email string) bool {
	var uid string
	query := `SELECT id FROM users WHERE email = ?;`
	err := utility.Db.QueryRow(query, email).Scan(&uid)
	if err != nil {
		return false
	}
	return true
}

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
	// validLen := len(user.Password) >= 8
	// hasUpper := checkReg(`[A-Z]`, user.Password)
	// hasLower := checkReg(`[a-z]`, user.Password)
	// hasPunc := checkReg(`[\p{P}\p{S}]`, user.Password)

	// passwordCheck := validLen && hasLower && hasUpper && hasPunc

	// if !passwordCheck {
	// 	return errors.New("PassWord Format Error")
	// }

	return nil
}

func checkReg(reg string, str string) bool {
	r := regexp.MustCompile(reg)
	match := r.Match([]byte(str))
	return match
}

func getUserByEmail(user *typefile.User, email string) error {
	query := `SELECT * FROM users WHERE email = ?;`
	err := utility.Db.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return err
	}
	return nil
}

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

func GetAuthUser(c echo.Context) error {
	claims := GetClaims(c)
	id := claims.ID
	user := typefile.User{}
	user.Password = "********"
	query := `select id, name, email from users where id = ?;`
	err := utility.Db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, "user not founded")
	}

	return c.JSON(http.StatusOK, user)
}
