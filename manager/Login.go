import (
    "fmt"
    "time"

    "github.com/dgrijalva/jwt-go"
)

func (dao LoginImpl) SaveNewUser(login models.Login) (int64, error) {

	query := "insert into registration(email,password) values(?,?)"
	db := connection()
	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
	return 0, err
	}

func ( []byte) (string, error) {
    // Create the token
    token := jwt.New(jwt.SigningMethodHS256)
    // Set some claims
    token.Claims["foo"] = "bar"
    token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(mySigningKey)
    return tokenString, err
}

func ExampleParse(myToken string, myKey string) {
    token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(myKey), nil
    })

    if err == nil && token.Valid {
        fmt.Println("Your token is valid.  I like your style.")
    } else {
        fmt.Println("This token is terrible!  I cannot accept this.")
    }
}