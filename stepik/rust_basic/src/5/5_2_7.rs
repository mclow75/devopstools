/*
* Напишите программу, которая считывает 5 целых чисел, записывает их в массив и выводит элементы коллекции по индексу считанных чисел, как показано в примере. Гарантируется, что числа не отрицательны и не превышают длины массива!

Тестовые данные ✅
Sample Input 1:
0
1
2
1
0
Sample Output 1:
0, 1, 2, 1, 0
Sample Input 2:
4
0
3
2
1
Sample Output 2:
1, 4, 2, 3, 0
*/
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut arr = [0_u8; 5];
    let mut str = String::new();

    for i in 0..5 {
        io::stdin().read_line(&mut str).expect(ERR_INPUT_READ);
        arr[i] = str.trim().parse().expect(ERR_INTEGER_READ);
        str.clear();
    }

    println!("{}, {}, {}, {}, {}", arr[0], arr[1], arr[2], arr[3], arr[4]);
}
