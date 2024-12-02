use std::fs;

pub fn solve() {
    let sum: i32 = fs::read_to_string("../input/day1.txt".to_string())
        .unwrap()
        .lines()
        .map(transform)
        .sum();
    println!("day1: {sum}")
}

fn transform(line: &str) -> i32 {
    let mut c1 = '0';
    let mut c2 = '0';
    for i in 0..line.len() {
        let c = line.chars().nth(i).unwrap();
        if c.is_digit(10) {
            c1 = c;
            break;
        }
        c1 = convert(&line[i..]);
        if c1 != '0' {
            break;
        }
    }
    for i in (0..line.len()).rev() {
        let c = line.chars().nth(i).unwrap();
        if c.is_digit(10) {
            c2 = c;
            break;
        }
        c2 = convert(&line[i..]);
        if c2 != '0' {
            break;
        }
    }
    format!("{}{}", c1, c2).parse().unwrap()
}

fn convert(line: &str) -> char {
    match line {
        ln if ln.starts_with("one") => '1',
        ln if ln.starts_with("two") => '2',
        ln if ln.starts_with("three") => '3',
        ln if ln.starts_with("four") => '4',
        ln if ln.starts_with("five") => '5',
        ln if ln.starts_with("six") => '6',
        ln if ln.starts_with("seven") => '7',
        ln if ln.starts_with("eight") => '8',
        ln if ln.starts_with("nine") => '9',
        _ => '0',
    }
}
