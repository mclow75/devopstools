/*
* Напишите программу, которая считывает 5 вещественных значений и 1 целое число (usize), записывает первые пять значений в массив и выводит (до 2 знаков) элемент коллекции по индексу последнего считанного числа, как показано в примере. Гарантируется, что последнее считанное число не превышает длины массива!
Sample Input:
1
2
3
4
5
0
Sample Output:
1.00
*/
use std::io;
const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut arr = [0_f64; 5];
    let mut str1 = String::new();
    let mut str2 = String::new();
    let mut str3 = String::new();
    let mut str4 = String::new();
    let mut str5 = String::new();
    let mut str_num = String::new();
    io::stdin().read_line(&mut str1).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str2).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str3).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str4).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str5).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str_num).expect(ERR_INPUT_READ);

    arr[0] = str1.trim().parse().expect(ERR_INTEGER_READ);
    arr[1] = str2.trim().parse().expect(ERR_INTEGER_READ);
    arr[2] = str3.trim().parse().expect(ERR_INTEGER_READ);
    arr[3] = str4.trim().parse().expect(ERR_INTEGER_READ);
    arr[4] = str5.trim().parse().expect(ERR_INTEGER_READ);
    let num: usize = str_num.trim().parse().expect(ERR_INTEGER_READ);
    println!("{:.2}", arr[num]);
}
