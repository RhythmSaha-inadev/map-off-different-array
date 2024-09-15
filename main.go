//Write a program to create a map by taking two different arrays data

package main

import "fmt"

func main() {

	env := []string{"dev", "qc", "uat", "prod"}
	ec2 := []string{"t3.micro", "t3.small", "t3.medium", "t3.large"}

	a := make(map[string]string)

	for i := 0; i < len(env); i++ {
		a[env[i]] = ec2[i]
	}

	fmt.Println(a)

}
