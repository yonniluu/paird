package main
import (
// "github.com/yonniluu/paird/vault"
"github.com/yonniluu/paird/yelp"
"time"


)
// main page to test functions
func main(){

// vault.Decrypt("vault:v1:oXS9ezMRv6UFulP9n2q4DSZ2WPc9JyuSTTzEbByXaGg=")
yelp.GetSuggestions(yelp.Location{City:"",Long:-123.118591, Lat:49.276654},time.Time{},[]string{"vegetarian"})
}