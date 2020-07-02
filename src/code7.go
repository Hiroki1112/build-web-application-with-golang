package main
import "fmt"

//ここではperson型を宣言
type person struct {
	name string
	age int
}

// 二人の年齢を比較します。年齢が大きい方の人を返し、また年齢差も返します。
// structも値渡しです。
func Older(p1, p2 person) (person, int) {
	if p1.age>p2.age {  // p1とp2の二人の年齢を比較します。
		return p1, p1.age-p2.age
	}
	return p2, p2.age-p1.age
}
//返り値はperson型、int型

func main() {
	var tom person

	// 初期値を代入します。
	tom.name, tom.age = "Tom", 18

	// ２つのフィールドを明確に初期化します。
	bob := person{age:25, name:"Bob"}

	// structの定義の順番に従って初期化します。
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)
}