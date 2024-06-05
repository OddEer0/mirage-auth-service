package main

import "github.com/OddEer0/mirage-auth-service/scripts"

func main() {
	scripts.RenameAllGoCodeFragment("github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1", "github.com/OddEer0/mirage-src/protogen/mirage-auth-service")
}
