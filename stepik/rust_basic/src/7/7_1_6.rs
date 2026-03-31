/*
* Напишите программу, которая считывает два вещественных числа и выводит (до 3 знаков) максимальное из них, как показано в примере.

Тестовые данные ✅
Sample Input:
5
10
Sample Output:
Из 5.000 и 10.000 больше 10.000
*/
fn main() {
    let a = input() as f64;
    let b = input() as f64;
    println!("Из {a:.3} и {b:.3} больше {:.3}", if a > b { a } else { b });
}

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn input() -> f64 {
    let mut buf = String::new();
    std::io::stdin().read_line(&mut buf).expect(ERR_INPUT_READ);
    return buf.trim().parse::<f64>().expect(ERR_INTEGER_READ);
}
