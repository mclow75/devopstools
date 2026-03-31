/*
* Напишите программу, которая считывает два целых числа. Далее инициализирует массив нулей размерностью десять и заменяет элемент в массиве по указанному индексу (первое число) на значение второго числа. После этого программа выводит получившийся массив. Гарантируется, что первое число не превышает длину массива!
Sample Input:
0
-10
Sample Output:
[-10, 0, 0, 0, 0, 0, 0, 0, 0, 0]
*/
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut arr = [0; 10];

    let mut str1 = String::new();
    let mut str2 = String::new();

    io::stdin().read_line(&mut str1).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str2).expect(ERR_INPUT_READ);
    let idx: usize = str1.trim().parse().expect(ERR_INTEGER_READ);
    arr[idx] = str2.trim().parse().expect(ERR_INTEGER_READ);
    println!("{:?}", arr);
}
