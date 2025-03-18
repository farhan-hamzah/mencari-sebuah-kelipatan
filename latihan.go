// Fungsi untuk mencetak kelipatan x atau y dari x sampai n
package main
import "fmt"

func kelipatan(n, x, y int) {
	found := false 

	for i := x; i <= n; i++ {
		if i%x == 0 || i%y == 0 {
			fmt.Print(i, " ")
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada kelipatan yang ditemukan")
	}
}

func main() {
	var n1, x1, y1 int

	fmt.Print("Masukkan n, x, y: ")
	fmt.Scan(&n1, &x1, &y1)

	kelipatan(n1, x1, y1)
}

