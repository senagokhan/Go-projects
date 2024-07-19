// --1--

package main

import "fmt"

func main() {

	problemAnswers1 := []int{1, 6, 3, 5, 10}
	score1 := 0

	problemAnswers2 := []int{2, 4, 7, 8, 9}
	score2 := 0

	for i := 0; i < 5; i++ {
		if problemAnswers1[i] > problemAnswers2[i] {
			score1 += 5
		} else {
			score2 += 5
		}
	}

	if score1 > score2 {
		fmt.Println("First array won, score :", score1)
	} else {
		fmt.Println("Second array won, score :", score2)
	}
}

/*
--2--

import (
	"fmt"
	"sort"
)

func main() {

	array1 := []int{4, 4, 6, 7, 3, 2, 4, 6, 80, 1, 25, 43, 5, 67, 9, 25, 5, 7, 8}
	var array2 []int

	sort.Ints(array1)

	for i := 0; i < len(array1); i++ {
		if i == 0 || array1[i] != array1[i-1] {
			array2 = append(array2, array1[i])
		}

	}
	fmt.Println(array2)
}
*/

/*
--3--

import (
    "fmt"
    "math/rand"
)

func main() {

    nonPalindrome := []int{}

    min, max := 999, 9999
    numbers := make([]int, 20)

    for i := 0; i < 20; i++ {
        numbers[i] = randomSayi(min, max)
    }

    for _, number := range numbers {
        if !isPalindrome(number) {
            nonPalindrome = append(nonPalindrome, number)
        }
    }

    fmt.Println("Random Numbers :", numbers)
    fmt.Println("Non-Polyndromic Numbers:", nonPalindrome)
}

func randomSayi(min, max int) int {
    return rand.Intn(max-min+1) + min
}

func isPalindrome(n int) bool {
    original := n
    reversed := 0

    for n > 0 {
        reversed = reversed*10 + n%10
        n /= 10
    }

    return original == reversed
}

*/

/*
--4--

import "fmt"

func main() {
    var input int
    fmt.Print("Faktöriyelini hesaplamak istediğiniz sayıyı giriniz: ")
    fmt.Scanln(&input)
    fmt.Println(faktoriyel(input))

}

func faktoriyel(a int) int {
    if a == 0 {
        return 1
    }
    return a * faktoriyel(a-1)
}

*/

/*
--5--

import (
    "fmt"
    "sort"
)

func main() {

    arr := []int{12, 5, 21, 77, 3, 42, 3, 11, 20, 27, 15, 99, 2, 58}

    sort.Ints(arr)

    fmt.Println("Dizi :", arr)
    fmt.Println("Dizinin Mode değeri :", modeBul(arr))
    fmt.Println("Dizinin Medyan değeri :", medyanBul(arr))

}

func modeBul(arr []int) int {

    var mode int
    var maxCount int

    for i := 0; i < len(arr); i++ {
        count := 0
        for j := 0; j < len(arr); j++ {
            if arr[j] == arr[i] {
                count++
            }
        }
        if count > maxCount {
            maxCount = count
            mode = arr[i]
        }
    }
    return mode
}

func medyanBul(arr []int) float64 {

    n := len(arr)

    if n%2 == 1 {
        return float64(arr[n/2])
    } else {
        return float64(arr[n/2-1]+arr[n/2]) / 2.0
    }
}

*/

/*
--6--

import (
    "fmt"
    "math/rand"
)

func main() {
    min := 1
    max := 3
    userScore := 0
    systemScore := 0
    var choice int

    fmt.Println("1-Taş ")
    fmt.Println("2-Kağıt ")
    fmt.Println("3-Makas ")

    for i := 0; i < 5; i++ {
        fmt.Println("1-2-3 değerlerinden birini seçiniz : ")
        fmt.Scanln(&choice)

        randomNumber := randomSayi(min, max)
        fmt.Println("Sistemin sayısı :", randomNumber)

        if randomNumber == 1 && choice == 3 {
            systemScore += 15
        } else if randomNumber == 2 && choice == 1 {
            systemScore += 15
        } else if randomNumber == 3 && choice == 2 {
            systemScore += 15
        } else if randomNumber == choice {

        } else {
            userScore += 15
        }}

    if systemScore > userScore {
        fmt.Println("Sistem kazandı, skor :", systemScore)
    } else if systemScore == userScore {
        fmt.Println("Skorlar eşit!")
    } else {
        fmt.Println("Kullanıcı kazandı, skor :", userScore)
    }
}

func randomSayi(min, max int) int {
    return rand.Intn(max-min+1) + min
}

*/
