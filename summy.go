package summy

// Sum รับ slice ของ int และคืนค่าผลรวมทั้งหมด
func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
