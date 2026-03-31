/*
* Напишите программу, которая считывает вещественное число, определяет, является оно положительным или отрицательным, и выводит (до 1 знака) результат, как показано в примере. Поскольку 0 не является ни отрицательным, ни положительным, то в тестах его нет.

Тестовые данные ✅
Sample Input:
10
Sample Output:
Число 10.0 является положительным
*/
fn main() {
    let n = input() as f64;
    if n > 0.0 {
        println!("Число {n:.1} является положительным");
    } else {
        println!("Число {n:.1} является отрицательным");
    }
}

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn input() -> f64 {
    let mut buf = String::new();
    std::io::stdin().read_line(&mut buf).expect(ERR_INPUT_READ);
    return buf.trim().parse::<f64>().expect(ERR_INTEGER_READ);
}
