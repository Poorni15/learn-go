package error-handling

import (
"errors"
"fmt"
"strings"
)

 var errorNameIsEmpty = errors.New("name is empty")
 var errorAgeNotValid =errors.New("age is less than 0 or greater than 150")
 var errorEmailNotValid =errors.New( "email is not valid")

type ValidationError struct {
   message string
}

func (validationError *ValidationError) Error ()  string {
 return fmt.Sprintf("validation error  %s",  validationError.message)
}

type User struct {
   Name string
   Age int
   Email string
}
func (user User) isAdult() bool{
  return user.Age >18
}

func newUser(name string , age int ,email string) ( *User , error ){
   if name =="" {
    return nil,fmt.Errorf("Invalid use name %s %w",name,errorNameIsEmpty)
    }
   if age <0 || age >150 {
     return nil,fmt.Errorf("Invalid user age  %d %w",age,errorAgeNotValid)
   }
   if !strings.Contains(email, "@") {
    return nil,fmt.Errorf("invalid user email   %s %w",email,errorEmailNotValid)
   }
   user:=User{Name:name,Age:age,Email:email}
   return &user,nil
}
func main (){
   name:= "John Doe"
   age := 34
   email := "johndoexyz.com"
   user, err := newUser(name,age,email)
   if err != nil {
   		fmt.Println("Error:", err)

   	} else {
   		fmt.Println("User created successfully:", *user)
   		fmt.Println("Is adult?", user.isAdult())
   	}
  }