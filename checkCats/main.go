// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

func checkCats(CatsTuti []int, CatsNining []int) {
    var CatsTutiCopy []int
    for i := 0; i < len(CatsTuti); i++ {
        if i == 0 || i > len(CatsTuti)-3 {
            CatsTutiCopy = append(CatsTutiCopy, 0)
        } else {
            CatsTutiCopy = append(CatsTutiCopy, CatsTuti[i])
        }
    }

    var correctionData []int
    correctionData = append(correctionData, CatsTutiCopy...)
    correctionData = append(correctionData, CatsNining...)

    for i, age := range correctionData {
        if age >= 3 {
            fmt.Printf("Kucing Nomor %d adalah Kucing Dewasa dan berusia %d tahun\n", i+1, age)
        } else {
            fmt.Printf("Kucing Nomor %d adalah Kucing Kecil (Kitten) dan masih anak-anak\n", i+1)
        }
    }
}

func main() {
    // Data uji pertama
    dataTuti1 := []int{3, 5, 2, 12, 7}
    dataNining1 := []int{4, 1, 15, 8, 3}
    fmt.Println("Data Uji 1:")
    checkCats(dataTuti1, dataNining1)

    // Data uji kedua
    dataTuti2 := []int{9, 16, 6, 8, 3}
    dataNining2 := []int{5, 10, 6, 1, 4}
    fmt.Println("\nData Uji 2:")
    checkCats(dataTuti2, dataNining2)
}
