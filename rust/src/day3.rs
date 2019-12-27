use std::cmp::Ordering;
use std::collections::{HashMap, HashSet};

#[derive(Debug, Eq, Ord, PartialEq, PartialOrd, Hash)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn from(x: i32, y: i32) -> Point {
        Point { x, y }
    }

    fn distance(&self) -> i32 {
        self.x.abs() + self.y.abs()
    }
}

pub fn manhattan() {
    let file = include_str!("input_day3");

    let mut wires: Vec<HashSet<Point>> = Vec::new();
    let mut steps_at_pos: Vec<HashMap<Point, i32>> = Vec::new();
    for line in file.lines() {
        let mut x = 0;
        let mut y = 0;
        let mut stepcount = 0;
        wires.push(HashSet::new());
        steps_at_pos.push(HashMap::new());
        for (d, steps) in line
            .split(',')
            .map(|x| x.split_at(1))
            .map(|(x, y)| (x, y.parse::<i32>().unwrap()))
        {
            for _ in 0..steps {
                stepcount += 1;
                match d {
                    "R" => {
                        x += 1;
                    }
                    "L" => {
                        x -= 1;
                    }
                    "U" => {
                        y += 1;
                    }
                    "D" => {
                        y -= 1;
                    }
                    _ => panic!("invalid direction"),
                }
                wires.last_mut().unwrap().insert(Point::from(x, y));
                if !steps_at_pos
                    .last()
                    .unwrap()
                    .contains_key(&Point::from(x, y))
                {
                    steps_at_pos
                        .last_mut()
                        .unwrap()
                        .insert(Point::from(x, y), stepcount);
                }
            }
        }
    }

    let mut intersections: Vec<&Point> = wires[0].intersection(&wires[1]).collect();
    // Part 1:
    intersections.sort_unstable_by(|x, y| compare(x, y));
    println!("Manhattan distance {}", intersections[0].distance());

    // Part 2:
    let mut step_pairs: Vec<_> = intersections
        .iter()
        .map(|x| (x, steps_at_pos[0][x]))
        .zip(intersections.iter().map(|x| (x, steps_at_pos[1][x])))
        .map(|((lp1, s1), (_, s2))| (lp1, s1 + s2))
        .collect();

    step_pairs.sort_unstable_by(|(_, s1), (_, s2)| s1.cmp(s2));
    println!("Fewest combined steps: {}", step_pairs[0].1);
}

fn compare(x: &Point, y: &Point) -> Ordering {
    return x.distance().cmp(&y.distance());
}