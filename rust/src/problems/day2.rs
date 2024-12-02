use std::fs;

pub fn solve() {
    let sum: i32 = fs::read_to_string("../input/day2.txt".to_string())
        .unwrap()
        .lines()
        .map(transform)
        .sum();
    println!("day2: {sum}")
}

fn transform(line: &str) -> i32 {
    let splits: Vec<&str> = line.split(": ").collect();
    // let gameNr: i32 = splits[0].split(" ").last().unwrap().parse().unwrap();
    splits[1]
        .split("; ")
        .map(|g| -> (u8, u8, u8) {
            let x: Vec<&str> = g.split(", ").collect();
            x.iter().map(|g| g.split(" ")).collect()
            // (0,0,0)
        })
        .for_each(|g| println!("{:?}", g));
    println!("\n");
    // println!("{}", splits[1]);

    return 0;
}

// fn isGamePassible() {
//     let (nr, ng, nb) = (12, 13, 14);
// }
